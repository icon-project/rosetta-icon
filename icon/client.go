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
	"fmt"

	RosettaTypes "github.com/coinbase/rosetta-sdk-go/types"
	"github.com/icon-project/goloop/common"
	"github.com/icon-project/rosetta-icon/icon/client_v1"
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

func (ic *Client) Status() (*RosettaTypes.BlockIdentifier, int64, []*RosettaTypes.Peer, error) {
	block, err := ic.iconV1.GetLastBlock(nil)
	if err != nil {
		return nil, -1, nil, err
	}

	peers, err := ic.GetPeer()
	if err != nil {
		return nil, -1, nil, err
	}

	return block.BlockIdentifier, block.Timestamp, peers, nil
}

func (ic *Client) GetBlock(params *RosettaTypes.PartialBlockIdentifier) (*RosettaTypes.Block, error) {
	var reqParams *client_v1.BlockRPCRequest
	var err error
	var block *RosettaTypes.Block
	if params.Index == nil && params.Hash == nil {
		reqParams = &client_v1.BlockRPCRequest{}
		block, err = ic.iconV1.GetLastBlock(reqParams)
	} else if params.Index != nil {
		reqParams = &client_v1.BlockRPCRequest{
			Height: common.HexInt64{Value: *params.Index}.String(),
		}
		block, err = ic.iconV1.GetBlockByHeight(reqParams)
	} else if params.Hash != nil {
		reqParams = &client_v1.BlockRPCRequest{
			Hash: *params.Hash,
		}
		block, err = ic.iconV1.GetBlockByHash(reqParams)
	} else {
		return nil, fmt.Errorf("invalid Params")
	}

	if err != nil {
		return nil, fmt.Errorf("%w: could not get block", err)
	}

	trsArray, err := ic.iconV1.GetReceipts(block)
	if err != nil {
		return nil, fmt.Errorf("%w: could not get blockReceipts", err)
	}
	ic.iconV1.MakeBlockWithReceipts(block, trsArray)
	return block, nil
}

func (ic *Client) GetTransaction(params *RosettaTypes.TransactionIdentifier) (*RosettaTypes.Transaction, error) {

	var reqParams *client_v1.TransactionRPCRequest
	reqParams = &client_v1.TransactionRPCRequest{
		Hash: params.Hash,
	}

	tx, err := ic.iconV1.GetTransaction(reqParams)
	if err != nil {
		return nil, fmt.Errorf("%w: could not get transaction", err)
	}

	txR, err := ic.iconV1.GetTransactionResult(reqParams)
	if err != nil {
		return nil, fmt.Errorf("%w: could not get transaction result", err)
	}
	ic.iconV1.MakeTransactionWithReceipt(tx, txR)
	return tx, nil
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

func (ic *Client) SendTransaction(tx client_v1.Transaction) error {
	js, err := tx.ToJSON()
	if err != nil {
		return err
	}
	if err := ic.iconV1.SendTransaction(js); err != nil {
		return err
	}
	return nil
}

func (ic *Client) GetDefaultStepCost() (*common.HexInt, error) {
	res, err := ic.iconV1.GetStepDefaultStepCost()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (ic *Client) GetBalance(
	account *RosettaTypes.AccountIdentifier,
	block *RosettaTypes.PartialBlockIdentifier,
) (*RosettaTypes.AccountBalanceResponse, error) {
	balReq := &client_v1.BalanceRPCRequest{
		Address: account.Address,
	}
	if block != nil && block.Index != nil {
		// result resides in the next block
		balReq.Height = common.HexInt64{Value: *block.Index + 1}.String()
	}
	balance, err := ic.iconV1.GetBalance(balReq)
	if err != nil {
		return nil, err
	}

	var blockId *RosettaTypes.BlockIdentifier
	if block != nil && block.Index != nil {
		blockReq := &client_v1.BlockRPCRequest{
			Height: common.HexInt64{Value: *block.Index}.String(),
		}
		blockByHeight, err := ic.iconV1.GetBlockByHeight(blockReq)
		if err != nil {
			return nil, fmt.Errorf("%w: could not get block", err)
		}
		blockId = blockByHeight.BlockIdentifier
	} else {
		lastBlock, err := ic.iconV1.GetLastBlock(nil)
		if err != nil {
			return nil, fmt.Errorf("%w: could not get last block", err)
		}
		blockId = lastBlock.ParentBlockIdentifier
	}

	return &RosettaTypes.AccountBalanceResponse{
		BlockIdentifier: &RosettaTypes.BlockIdentifier{
			Index: blockId.Index,
			Hash:  blockId.Hash,
		},
		Balances: []*RosettaTypes.Amount{
			{
				Value:    balance.Text(10),
				Currency: client_v1.ICXCurrency,
			},
		},
	}, nil
}
