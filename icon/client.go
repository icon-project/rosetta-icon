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
	"net/http"
	"strings"

	RosettaTypes "github.com/coinbase/rosetta-sdk-go/types"
	"github.com/icon-project/goloop/common"
)

// Client is used to fetch blocks from ICON Node and
// to parser ICON block data into Rosetta types.
type Client struct {
	v3 *ClientV3
	rc *JsonRpcClient
}

func NewClient(endpoint string) *Client {
	// increase the maximum idle connections to solve
	// "connect: cannot assign requested address" problem
	http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = 100
	client := new(http.Client)
	url := []string{
		endpoint,
		EndpointPrefix,
		EndpointRosetta,
	}
	return &Client{
		v3: NewClientV3(endpoint),
		rc: NewJsonRpcClient(client, strings.Join(url, "/")),
	}
}

func (ic *Client) Status() (*RosettaTypes.BlockIdentifier, int64, []*RosettaTypes.Peer, error) {
	block, err := ic.v3.getLastBlock()
	if err != nil {
		return nil, -1, nil, err
	}

	blockIdentifier := &RosettaTypes.BlockIdentifier{
		Index: block.Height,
		Hash:  block.BlockHash.String(),
	}

	peers, err := ic.GetPeer()
	if err != nil {
		return nil, -1, nil, err
	}

	return blockIdentifier, block.Timestamp / 1000, peers, nil
}

func (ic *Client) GetBlock(params *RosettaTypes.PartialBlockIdentifier) (*RosettaTypes.Block, error) {
	reqParams := &RosettaTraceParam{}
	if params.Hash != nil {
		reqParams.Block = *params.Hash
	} else if params.Index != nil {
		if *params.Index == 0 {
			return ic.v3.getBlock(params)
		}
		reqParams.Height = common.HexInt64{Value: *params.Index}.String()
	}
	trace, err := ic.getRosettaTrace(reqParams)
	if err != nil {
		return nil, err
	}
	transactions, err := ic.populateTransactions(trace.BalanceChanges)
	if err != nil {
		return nil, err
	}
	return &RosettaTypes.Block{
		BlockIdentifier: &RosettaTypes.BlockIdentifier{
			Index: trace.Index(),
			Hash:  trace.BlockHash,
		},
		ParentBlockIdentifier: &RosettaTypes.BlockIdentifier{
			Index: trace.Index() - 1,
			Hash:  trace.PrevBlockHash,
		},
		Timestamp:    trace.TimestampInMillis(),
		Transactions: transactions,
	}, nil
}

func (ic *Client) getRosettaTrace(param *RosettaTraceParam) (*RosettaTraceResponse, error) {
	req, err := GetRpcRequest("rosetta_getTrace", param, -1)
	if err != nil {
		return nil, err
	}
	trace := &RosettaTraceResponse{}
	_, err = ic.rc.Request(req, trace)
	if err != nil {
		return nil, err
	}
	return trace, nil
}

func (ic *Client) populateTransactions(balChanges []*BalanceChange) ([]*RosettaTypes.Transaction, error) {
	transactions := make([]*RosettaTypes.Transaction, len(balChanges))
	for i, bc := range balChanges {
		tx, err := ic.populateTransaction(bc)
		if err != nil {
			return nil, fmt.Errorf("%w: cannot parse %s", err, bc.TxHash)
		}
		transactions[i] = tx
	}
	return transactions, nil
}

func (ic *Client) populateTransaction(bc *BalanceChange) (*RosettaTypes.Transaction, error) {
	var ops []*RosettaTypes.Operation
	for _, op := range bc.Ops {
		lastIndex := int64(len(ops))
		fromOp := &RosettaTypes.Operation{
			OperationIdentifier: &RosettaTypes.OperationIdentifier{
				Index: lastIndex,
			},
			Type:   op.OpType,
			Status: RosettaTypes.String(SuccessStatus),
			Account: &RosettaTypes.AccountIdentifier{
				Address: op.From,
			},
			Amount: &RosettaTypes.Amount{
				Value:    "-" + op.IntValue(),
				Currency: ICXCurrency,
			},
		}
		ops = append(ops, fromOp)

		toOp := &RosettaTypes.Operation{
			OperationIdentifier: &RosettaTypes.OperationIdentifier{
				Index: lastIndex + 1,
			},
			RelatedOperations: []*RosettaTypes.OperationIdentifier{
				{
					Index: lastIndex,
				},
			},
			Type:   op.OpType,
			Status: RosettaTypes.String(SuccessStatus),
			Account: &RosettaTypes.AccountIdentifier{
				Address: op.To,
			},
			Amount: &RosettaTypes.Amount{
				Value:    op.IntValue(),
				Currency: ICXCurrency,
			},
		}
		ops = append(ops, toOp)

		// some assertion check
		if op.OpType == FeeOpType && op.To != TreasuryAddress {
			return nil, fmt.Errorf("invalid fee operation")
		}
	}

	return &RosettaTypes.Transaction{
		TransactionIdentifier: &RosettaTypes.TransactionIdentifier{
			Hash: bc.TxHash,
		},
		Operations: ops,
	}, nil
}

func (ic *Client) GetTransaction(params *RosettaTypes.TransactionIdentifier) (*RosettaTypes.Transaction, error) {
	var reqParams *TransactionRPCRequest
	reqParams = &TransactionRPCRequest{
		Hash: params.Hash,
	}

	tx, err := ic.v3.getTransaction(reqParams)
	if err != nil {
		return nil, fmt.Errorf("%w: could not get transaction", err)
	}

	txR, err := ic.v3.getTransactionResult(reqParams)
	if err != nil {
		return nil, fmt.Errorf("%w: could not get transaction result", err)
	}
	ic.v3.makeTransactionWithReceipt(tx, txR)
	return tx, nil
}

func (ic *Client) GetPeer() ([]*RosettaTypes.Peer, error) {
	resp, err := ic.v3.getMainPReps()
	if err != nil {
		return nil, fmt.Errorf("%w: could not get peer", err)
	}

	var peers []*RosettaTypes.Peer
	preps := (*resp)["preps"]

	for _, element := range preps.([]interface{}) {
		address := element.(map[string]interface{})["address"]
		resp, _ := ic.v3.getPRep(address.(string))
		peers = append(peers, &RosettaTypes.Peer{
			PeerID:   address.(string),
			Metadata: *resp,
		})
	}

	return peers, nil
}

func (ic *Client) SendTransaction(tx Transaction) error {
	js, err := tx.ToJSON()
	if err != nil {
		return err
	}
	if err := ic.v3.sendTransaction(js); err != nil {
		return err
	}
	return nil
}

func (ic *Client) GetDefaultStepCost() (*common.HexInt, error) {
	res, err := ic.v3.getStepDefaultStepCost()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (ic *Client) GetBalance(
	account *RosettaTypes.AccountIdentifier,
	block *RosettaTypes.PartialBlockIdentifier,
) (*RosettaTypes.AccountBalanceResponse, error) {
	balReq := &BalanceRPCRequest{
		Address: account.Address,
	}
	if block != nil && block.Index != nil {
		// result resides in the next block
		balReq.Height = common.HexInt64{Value: *block.Index + 1}.String()
	}
	balance, err := ic.v3.getBalance(balReq)
	if err != nil {
		return nil, err
	}

	var blockResp *Block
	if block != nil && block.Index != nil {
		blockReq := &BlockRPCRequest{
			Height: common.HexInt64{Value: *block.Index}.String(),
		}
		blockResp, err = ic.v3.getBlockByHeight(blockReq)
		if err != nil {
			return nil, fmt.Errorf("%w: could not get block", err)
		}
	} else {
		blockResp, err = ic.v3.getLastBlock()
		if err != nil {
			return nil, fmt.Errorf("%w: could not get last block", err)
		}
	}

	return &RosettaTypes.AccountBalanceResponse{
		BlockIdentifier: &RosettaTypes.BlockIdentifier{
			Index: blockResp.Height,
			Hash:  blockResp.BlockHash.String(),
		},
		Balances: []*RosettaTypes.Amount{
			{
				Value:    balance.Text(10),
				Currency: ICXCurrency,
			},
		},
	}, nil
}
