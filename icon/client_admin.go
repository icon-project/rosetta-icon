// Copyright 2022 ICON Foundation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package icon

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type ClientAdmin struct {
	endpoint string
	cid      string
}

func NewClientAdmin(endpoint string) *ClientAdmin {
	url := []string{
		endpoint,
		EndpointAdmin,
	}
	return &ClientAdmin{
		endpoint: strings.Join(url, "/"),
	}
}

func (c *ClientAdmin) getChainList() ([]map[string]interface{}, error) {
	res, err := http.Get(fmt.Sprintf("%s/chain", c.endpoint))
	if err != nil {
		return nil, fmt.Errorf("%w: could not get chain list", err)
	}
	var chains []map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&chains)
	if err != nil {
		return nil, fmt.Errorf("%w: could not decode response body", err)
	}
	return chains, nil
}

func (c *ClientAdmin) getChainInfo(cid string) (map[string]interface{}, error) {
	res, err := http.Get(fmt.Sprintf("%s/chain/%s", c.endpoint, c.cid))
	if err != nil {
		return nil, fmt.Errorf("%w: could not get chain info for %s", err, c.cid)
	}
	var chainInfo map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&chainInfo)
	if err != nil {
		return nil, fmt.Errorf("%w: could not decode response body", err)
	}
	return chainInfo, nil
}

func (c *ClientAdmin) getPeers() ([]interface{}, error) {
	if c.cid == "" {
		chains, err := c.getChainList()
		if err != nil {
			return nil, err
		}
		if len(chains) == 0 {
			return nil, fmt.Errorf("no active chains")
		}
		c.cid = chains[0]["cid"].(string)
	}
	chainInfo, err := c.getChainInfo(c.cid)
	if err != nil {
		return nil, err
	}
	network := chainInfo["module"].(map[string]interface{})["network"]
	p2p := network.(map[string]interface{})["p2p"]
	parent := p2p.(map[string]interface{})["parent"]
	uncles := p2p.(map[string]interface{})["uncles"]

	var peers []interface{}
	peers = append(peers, parent)
	peers = append(peers, uncles.([]interface{})...)
	return peers, nil
}
