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
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strings"

	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/icon-project/goloop/common"
	"github.com/icon-project/goloop/server/jsonrpc"
)

type ClientV3 struct {
	*JsonRpcClient
}

func NewClientV3(endpoint string) *ClientV3 {
	client := new(http.Client)
	url := []string{
		endpoint,
		EndpointPrefix,
		EndpointVersion,
	}
	return &ClientV3{
		JsonRpcClient: NewJsonRpcClient(client, strings.Join(url, "/")),
	}
}

func (c *ClientV3) getBlock(params *types.PartialBlockIdentifier) (*types.Block, error) {
	var reqParams *BlockRPCRequest
	var err error
	var block *Block
	if params.Index == nil && params.Hash == nil {
		block, err = c.getLastBlock()
	} else if params.Index != nil {
		reqParams = &BlockRPCRequest{
			Height: common.HexInt64{Value: *params.Index}.String(),
		}
		block, err = c.getBlockByHeight(reqParams)
	} else if params.Hash != nil {
		reqParams = &BlockRPCRequest{
			Hash: *params.Hash,
		}
		block, err = c.getBlockByHash(reqParams)
	} else {
		return nil, fmt.Errorf("invalid Params")
	}
	if err != nil {
		return nil, fmt.Errorf("%w: could not get block", err)
	}

	rtBlock, err := ParseBlock(block)
	if err != nil {
		return nil, err
	}

	trsArray, err := c.getReceipts(rtBlock)
	if err != nil {
		return nil, fmt.Errorf("%w: could not get blockReceipts", err)
	}
	c.makeBlockWithReceipts(rtBlock, trsArray)
	return rtBlock, nil
}

func (c *ClientV3) getLastBlock() (*Block, error) {
	block := &Block{}
	jrReq, err := GetRpcRequest("icx_getLastBlock", nil, -1)
	if err != nil {
		return nil, err
	}
	_, err = c.Request(jrReq, block)
	if err != nil {
		return nil, err
	}
	return block, nil
}

func (c *ClientV3) getBlockByHeight(param *BlockRPCRequest) (*Block, error) {
	block := &Block{}
	jrReq, err := GetRpcRequest("icx_getBlockByHeight", param, -1)
	if err != nil {
		return nil, err
	}
	_, err = c.Request(jrReq, block)
	if err != nil {
		return nil, err
	}
	return block, nil
}

func (c *ClientV3) getBlockByHash(param *BlockRPCRequest) (*Block, error) {
	block := &Block{}
	jrReq, err := GetRpcRequest("icx_getBlockByHash", param, -1)
	if err != nil {
		return nil, err
	}
	_, err = c.Request(jrReq, block)
	if err != nil {
		return nil, err
	}
	return block, nil
}

func (c *ClientV3) getReceipts(block *types.Block) ([]*TransactionResult, error) {
	batchSize := 10
	trsRaw := make([]interface{}, len(block.Transactions))
	reqs := make([]*jsonrpc.Request, 0)
	resp := make([]interface{}, batchSize)
	destIdx := 0
	for i, tx := range block.Transactions {
		mod := (i + 1) % batchSize
		txHash := tx.TransactionIdentifier.Hash
		if txHash == GenesisTxHash {
			continue
		}
		reqParams := &TransactionRPCRequest{
			Hash: txHash,
		}
		jrReq, err := GetRpcRequest("icx_getTransactionResult", reqParams, int64(i))
		if err != nil {
			return nil, err
		}
		reqs = append(reqs, jrReq)
		if mod == 0 {
			_, err = c.RequestBatch(reqs, resp)
			copy(trsRaw[destIdx:], resp)
			destIdx += batchSize
			reqs = make([]*jsonrpc.Request, 0)
		} else if (i + 1) == len(block.Transactions) {
			_, err = c.RequestBatch(reqs, resp)
			copy(trsRaw[destIdx:], resp[:mod])
		}
		if err != nil {
			return nil, err
		}
	}
	trsArray, err := ParseTransactionResults(&trsRaw)
	if err != nil {
		return nil, err
	}
	return trsArray, nil
}

func (c *ClientV3) makeBlockWithReceipts(block *types.Block, trsArray []*TransactionResult) (*types.Block, error) {
	zeroBigInt := new(big.Int)
	fa := SystemScoreAddress
	for index, tx := range block.Transactions {
		tx = block.Transactions[index]
		if tx.TransactionIdentifier.Hash == GenesisTxHash {
			continue
		}
		if len(tx.Operations) >= 4 { //general tx(transfer, call, deploy...)
			su := trsArray[index].StepUsed
			sp := trsArray[index].StepPrice
			sd := trsArray[index].StepDetails
			fa = tx.Operations[0].Account.Address
			if su.Cmp(zeroBigInt) != 0 {
				f := new(big.Int).Mul(&su.Int, &sp.Int)
				fee := f.Text(10)
				tx.Operations[3].Amount.Value = fee
				userStep := &su.Int
				if len(sd) != 0 {
					userStep = getUserStep(tx.Operations[2].Account.Address, sd)
				}
				tx.Operations[2].Amount.Value = "-" + userStep.Mul(userStep, &sp.Int).Text(10)
			}
			HandleBugTransaction(tx, fa)
		}
		if trsArray[index].EventLogs != nil {
			ops := GetOperations(fa, trsArray[index].EventLogs, int64(len(tx.Operations))-1)
			tx.Operations = append(tx.Operations, ops...)
		}
		for _, op := range tx.Operations {
			if op.Type != FeeOpType {
				op.Status = trsArray[index].StatusFlag
			}
		}
	}
	return block, nil
}

func (c *ClientV3) getTransaction(param *TransactionRPCRequest) (*types.Transaction, error) {
	txRaw := map[string]interface{}{}
	jrReq, err := GetRpcRequest("icx_getTransactionByHash", param, -1)
	if err != nil {
		return nil, err
	}
	_, err = c.Request(jrReq, &txRaw)
	if err != nil {
		return nil, err
	}

	var txRaws []json.RawMessage
	b, _ := json.Marshal(&txRaw)
	raw := json.RawMessage(b)
	txRaws = append(txRaws, raw)
	txs, _ := ParseTransactions(txRaws)
	return txs[0], nil
}

func (c *ClientV3) getTransactionResult(param *TransactionRPCRequest) (*TransactionResult, error) {
	trRaw := map[string]interface{}{}
	jrReq, err := GetRpcRequest("icx_getTransactionResult", param, -1)
	if err != nil {
		return nil, err
	}
	_, err = c.Request(jrReq, &trRaw)
	if err != nil {
		return nil, err
	}

	txRs, _ := ParseTransactionResult(trRaw)
	return txRs, nil
}

func (c *ClientV3) makeTransactionWithReceipt(tx *types.Transaction, txResult *TransactionResult) (*types.Transaction, error) {
	zeroBigInt := new(big.Int)
	fa := SystemScoreAddress
	if len(tx.Operations) >= 4 { //general tx(transfer, call, deploy...)
		su := txResult.StepUsed
		sp := txResult.StepPrice
		sd := txResult.StepDetails
		if su.Cmp(zeroBigInt) != 0 {
			f := new(big.Int).Mul(&su.Int, &sp.Int)
			fee := f.Text(10)
			tx.Operations[3].Amount.Value = fee
			userStep := &su.Int
			if len(sd) != 0 {
				userStep = getUserStep(tx.Operations[2].Account.Address, sd)
			}
			tx.Operations[2].Amount.Value = "-" + userStep.Mul(userStep, &sp.Int).Text(10)
			fa = tx.Operations[0].Account.Address
		}
	}
	if txResult.EventLogs != nil {
		ops := GetOperations(fa, txResult.EventLogs, int64(len(tx.Operations))-1)
		tx.Operations = append(tx.Operations, ops...)
	}
	for _, op := range tx.Operations {
		if op.Type != FeeOpType {
			op.Status = txResult.StatusFlag
		}
	}
	return tx, nil
}

func (c *ClientV3) getBalance(param *BalanceRPCRequest) (*common.HexInt, error) {
	req, err := GetRpcRequest("icx_getBalance", param, -1)
	if err != nil {
		return nil, err
	}
	balance := &common.HexInt{}
	if _, err := c.Request(req, &balance); err != nil {
		return nil, err
	}
	return balance, nil
}

func (c *ClientV3) sendTransaction(req interface{}) error {
	resp := ""
	jrReq, err := GetRpcRequest("icx_sendTransaction", req, -1)
	if err != nil {
		return err
	}
	_, err = c.Request(jrReq, &resp)
	if err != nil {
		return err
	}
	return nil
}

func (c *ClientV3) getStepDefaultStepCost() (*common.HexInt, error) {
	resp := map[string]*common.HexInt{}
	params := map[string]interface{}{
		"to":       "cx0000000000000000000000000000000000000000",
		"dataType": "call",
		"data": map[string]interface{}{
			"method": "getStepCosts",
		},
	}
	jrReq, err := GetRpcRequest("icx_call", params, -1)
	if err != nil {
		return nil, err
	}
	_, err = c.Request(jrReq, &resp)
	if err != nil {
		return nil, err
	}
	return resp["default"], nil
}

func getUserStep(from string, stepDetails map[string]*common.HexInt) *big.Int {
	userUsed := new(big.Int)
	for f, v := range stepDetails {
		if from == f {
			userUsed.Set(&v.Int)
		}
	}
	return userUsed
}
