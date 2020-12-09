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

package client_v1

import (
	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/icon-project/goloop/server/jsonrpc"
	v3 "github.com/icon-project/goloop/server/v3"
	"net/http"
	"net/url"
	"strings"
)

type ClientV3 struct {
	*JsonRpcClient
	DebugEndPoint string
}

func guessDebugEndpoint(endpoint string) string {
	uo, err := url.Parse(endpoint)
	if err != nil {
		return ""
	}
	ps := strings.Split(uo.Path, "/")
	for i, v := range ps {
		if v == "api" {
			if len(ps) > i+1 && ps[i+1] == "v3" {
				ps[i+1] = "v3d"
				uo.Path = strings.Join(ps, "/")
				return uo.String()
			}
			break
		}
	}
	return ""
}

func NewClientV3(endpoint string) *ClientV3 {
	client := new(http.Client)
	apiClient := NewJsonRpcClient(client, endpoint)

	return &ClientV3{
		JsonRpcClient: apiClient,
		DebugEndPoint: guessDebugEndpoint(endpoint),
	}
}

func (c *ClientV3) GetBlock(param *BlockRPCRequest) (*types.Block, error) {
	blockRaw := map[string]interface{}{}

	_, err := c.Do("icx_getBlock", param, &blockRaw)
	if err != nil {
		return nil, err
	}

	// TODO Block을 Version에 맞춰서 변환
	block, err := ParseBlock(blockRaw)
	if err != nil {
		return nil, err
	}
	return block, nil
}

func (c *ClientV3) GetBlockReceipts(param *BlockRPCRequest) (interface{}, error) {
	trsRaw := map[string]interface{}{}

	_, err := c.Do("icx_getBlockReceipts", param, trsRaw)
	if err != nil {
		return nil, err
	}
	return trsRaw, nil
}

func (c *ClientV3) GetTransaction(param *TransactionRPCRequest) (interface{}, error) {
	txRaw := map[string]interface{}{}

	_, err := c.Do("icx_getTransactionByHash", param, txRaw)
	if err != nil {
		return nil, err
	}
	return txRaw, nil
}

func (c *ClientV3) GetTransactionResult(param *TransactionRPCRequest) (interface{}, error) {
	trRaw := map[string]interface{}{}

	_, err := c.Do("icx_getTransactionResult", param, trRaw)
	if err != nil {
		return nil, err
	}
	return trRaw, nil
}

func (c *ClientV3) GetBalance(param *v3.AddressParam) (*jsonrpc.HexInt, error) {
	var result jsonrpc.HexInt
	_, err := c.Do("icx_getBalance", param, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *ClientV3) GetTotalSupply() (*jsonrpc.HexInt, error) {
	var result jsonrpc.HexInt
	_, err := c.Do("icx_getTotalSupply", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *ClientV3) GetMainPReps() (*map[string]interface{}, error) {
	resp := map[string]interface{}{}

	params := map[string]interface{}{
		"to":       "cx0000000000000000000000000000000000000000",
		"dataType": "call",
		"data": map[string]interface{}{
			"method": "getMainPReps",
		},
	}

	_, err := c.Do("icx_call", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *ClientV3) GetPRep(prep string) (*map[string]interface{}, error) {
	resp := map[string]interface{}{}

	params := map[string]interface{}{
		"to":       "cx0000000000000000000000000000000000000000",
		"dataType": "call",
		"data": map[string]interface{}{
			"method": "getPRep",
			"params": map[string]interface{}{
				"address": prep,
			},
		},
	}

	_, err := c.Do("icx_call", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
