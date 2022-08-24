// Copyright 2020 ICON Foundation, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package icon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"
	"time"

	"github.com/icon-project/goloop/server/jsonrpc"
)

const (
	headerContentType   = "Content-Type"
	headerAccept        = "Accept"
	typeApplicationJSON = "application/json"
)

type JsonRpcClient struct {
	hc           *http.Client
	Endpoint     string
	CustomHeader map[string]string
	Pre          func(req *http.Request) error
}

type Response struct {
	Version string          `json:"jsonrpc"`
	Result  json.RawMessage `json:"result"`
	Error   *jsonrpc.Error  `json:"error,omitempty"`
	ID      interface{}     `json:"id"`
}

type HttpError struct {
	response string
	message  string
}

func (e *HttpError) Error() string {
	return e.message
}

func (e *HttpError) Response() string {
	return e.response
}

func NewHttpError(r *http.Response) error {
	var response string
	if rb, err := ioutil.ReadAll(r.Body); err != nil {
		response = fmt.Sprintf("Fail to read body err=%+v", err)
	} else {
		response = string(rb)
	}
	return &HttpError{
		message:  "HTTP " + r.Status,
		response: response,
	}
}

func NewJsonRpcClient(hc *http.Client, endpoint string) *JsonRpcClient {
	return &JsonRpcClient{hc: hc, Endpoint: endpoint, CustomHeader: make(map[string]string)}
}

func (c *JsonRpcClient) do(req *http.Request) (resp *http.Response, err error) {
	if c.Pre != nil {
		if err = c.Pre(req); err != nil {
			return nil, err
		}
	}
	resp, err = c.hc.Do(req)
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("http-status(%s) is not StatusOK", resp.Status)
		return
	}
	return
}

func (c *JsonRpcClient) Request(jrReq *jsonrpc.Request, respPtr interface{}) (*Response, error) {
	req, err := getHttpRequest(c.Endpoint, jrReq)
	if err != nil {
		return nil, err
	}
	resp, err := c.do(req)
	if err != nil {
		_, err = handleErrorResponse(resp, err)
		return nil, err
	}
	response, err := decodeResponse(resp, respPtr)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *JsonRpcClient) RequestBatch(jrReq []*jsonrpc.Request, respPtr []interface{}) ([]*Response, error) {
	req, err := getHttpRequest(c.Endpoint, jrReq)
	if err != nil {
		return nil, err
	}
	resp, err := c.do(req)
	if err != nil {
		_, err = handleErrorResponse(resp, err)
		return nil, err
	}
	res, err := decodeBatchResponse(resp, respPtr)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func handleErrorResponse(resp *http.Response, jrErr error) (jrResp *Response, err error) {
	var dErr error
	if resp != nil {
		if ct, _, mErr := mime.ParseMediaType(resp.Header.Get(headerContentType)); mErr != nil {
			err = mErr
			return
		} else if ct == typeApplicationJSON {
			if jrResp, dErr = decodeResponseBody(resp); dErr != nil {
				err = dErr
				return
			}
			err = jrResp.Error
			return
		} else {
			err = NewHttpError(resp)
			return
		}
	}
	err = jrErr
	return
}

func decodeResponse(resp *http.Response, respPtr interface{}) (jrResp *Response, err error) {
	var dErr error

	if jrResp, dErr = decodeResponseBody(resp); dErr != nil {
		err = fmt.Errorf("fail to decode response body err:%+v, jsonrpcResp:%+v",
			dErr, resp)
		return
	}
	if jrResp.Error != nil {
		err = jrResp.Error
		return
	}
	if respPtr != nil {
		err = json.Unmarshal(jrResp.Result, respPtr)
		if err != nil {
			return
		}
	}
	return
}

func decodeBatchResponse(resp *http.Response, respPtr []interface{}) ([]*Response, error) {
	var jrResp []*Response
	var err error
	if jrResp, err = decodeBatchResponseBody(resp); err != nil {
		return nil, fmt.Errorf("fail to decode response body err:%+v, jsonrpcResp:%+v",
			err, resp)
	}
	if respPtr != nil {
		for i, res := range jrResp {
			if res.Error != nil {
				return nil, res.Error
			}
			err = json.Unmarshal(res.Result, &respPtr[i])
			if err != nil {
				return nil, err
			}
		}
	}
	return jrResp, nil
}

func decodeResponseBody(resp *http.Response) (jrResp *Response, err error) {
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&jrResp)
	return
}

func decodeBatchResponseBody(resp *http.Response) (jrResp []*Response, err error) {
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&jrResp)
	return
}

func GetRpcRequest(method string, reqPtr interface{}, id int64) (*jsonrpc.Request, error) {
	if id == -1 {
		id = time.Now().UnixNano() / int64(time.Millisecond)
	}
	jrReq := &jsonrpc.Request{
		ID:      id,
		Version: jsonrpc.Version,
		Method:  &method,
	}
	if reqPtr != nil {
		b, mErr := json.Marshal(reqPtr)
		if mErr != nil {
			err := mErr
			return nil, err
		}
		jrReq.Params = b
	}
	return jrReq, nil
}

func getHttpRequest(url string, jrReq interface{}) (*http.Request, error) {
	reqB, err := json.Marshal(jrReq)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewReader(reqB))
	if err != nil {
		return nil, err
	}
	req.Header.Set(headerContentType, typeApplicationJSON)
	req.Header.Set(headerAccept, typeApplicationJSON)
	return req, nil
}
