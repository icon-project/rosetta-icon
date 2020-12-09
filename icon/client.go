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
	"errors"
	"fmt"
	RosettaTypes "github.com/coinbase/rosetta-sdk-go/types"
	"github.com/icon-project/goloop/common"
	"github.com/leeheonseung/rosetta-icon/icon/client_v1"
)

// Client is used to fetch blocks from ICON Node and
// to parser ICON block data into Rosetta types.
//
// We opted not to use existing ICON RPC libraries
// because they don't allow providing context
// in each request.

type Client struct {
	currency *RosettaTypes.Currency
	iconV1   *client_v1.ClientV3
}

func NewClient(
	endpoint string,
	currency *RosettaTypes.Currency,
) *Client {
	return &Client{
		currency,
		client_v1.NewClientV3(endpoint),
	}
}

func (ic *Client) GetBlock(params *RosettaTypes.PartialBlockIdentifier) (*RosettaTypes.Block, error) {

	//이렇게 하는 방법밖에 없는가?
	var reqParams *client_v1.BlockRPCRequest
	if params.Index != nil && params.Hash != nil {
		return nil, errors.New("Invalid Both value")
	} else if params.Hash != nil {
		reqParams = &client_v1.BlockRPCRequest{
			Hash: *params.Hash,
		}
	} else if params.Index != nil {
		reqParams = &client_v1.BlockRPCRequest{
			Height: common.HexInt64{Value: *params.Index}.String(),
		}
	} else {
		reqParams = &client_v1.BlockRPCRequest{}
	}

	block, err := ic.iconV1.GetBlock(reqParams)
	if err != nil {
		return nil, fmt.Errorf("%w: could not get block", err)
	}

	//Get transactionResults
	//RPCServer, LoopChain branch참고: append_icx_getBlockReceipts

	//trsRaw, err := ic.iconV1.GetBlockReceipts(reqParams)
	//if err != nil {
	//	return nil, fmt.Errorf("%w: could not get block", err)
	//}
	//fmt.Print(trsRaw)
	return block, nil
}

func (ic *Client) GetPeer() ([]*RosettaTypes.Peer, error) {
	resp, err := ic.iconV1.GetMainPReps()
	if err != nil {
		return nil, fmt.Errorf("%w: could not get peer", err)
	}

	var peers []*RosettaTypes.Peer
	preps := (*resp)["preps"]

	for _, element := range preps.([]interface{}) {
		address := element.(map[string]interface{})["address"]
		resp, _ := ic.iconV1.GetPRep(address.(string))
		peers = append(peers, &RosettaTypes.Peer{
			PeerID:   address.(string),
			Metadata: *resp,
		})
	}

	return peers, nil
}
