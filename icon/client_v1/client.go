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
	"encoding/json"
	"math/big"
	"net/http"
	"time"

	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/icon-project/goloop/common"
	"github.com/icon-project/goloop/server/jsonrpc"
)

type ClientV3 struct {
	*JsonRpcClient
}

func NewClientV3(endpoint string) *ClientV3 {
	client := new(http.Client)
	apiClient := NewJsonRpcClient(client, endpoint)

	return &ClientV3{
		JsonRpcClient: apiClient,
	}
}

func (c *ClientV3) GetLastBlock(param *BlockRPCRequest) (*types.Block, error) {
	blockRaw := map[string]interface{}{}
	id := time.Now().UnixNano() / int64(time.Millisecond)

	jrReq, err := GetRpcRequest("icx_getLastBlock", param, id)
	if err != nil {
		return nil, err
	}
	_, err = c.Request(jrReq, &blockRaw)
	if err != nil {
		return nil, err
	}

	block, err := ParseBlock(blockRaw)
	if err != nil {
		return nil, err
	}
	return block, nil
}

func (c *ClientV3) GetBlockByHeight(param *BlockRPCRequest) (*types.Block, error) {
	blockRaw := map[string]interface{}{}
	id := time.Now().UnixNano() / int64(time.Millisecond)

	jrReq, err := GetRpcRequest("icx_getBlockByHeight", param, id)
	if err != nil {
		return nil, err
	}
	_, err = c.Request(jrReq, &blockRaw)
	if err != nil {
		return nil, err
	}

	block, err := ParseBlock(blockRaw)
	if err != nil {
		return nil, err
	}

	return block, nil
}

func (c *ClientV3) GetBlockByHash(param *BlockRPCRequest) (*types.Block, error) {
	blockRaw := map[string]interface{}{}

	id := time.Now().UnixNano() / int64(time.Millisecond)
	jrReq, err := GetRpcRequest("icx_getBlockByHash", param, id)
	if err != nil {
		return nil, err
	}
	_, err = c.Request(jrReq, &blockRaw)
	if err != nil {
		return nil, err
	}

	block, err := ParseBlock(blockRaw)
	if err != nil {
		return nil, err
	}
	return block, nil
}

func (c *ClientV3) GetReceipts(block *types.Block) ([]*TransactionResult, error) {
	trsRaw := make([]interface{}, len(block.Transactions))
	reqs := make([]*jsonrpc.Request, 0)
	for i, tx := range block.Transactions {
		index := i + 1
		mod := index % 10
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
			_, err = c.RequestBatch(reqs, trsRaw)
			reqs = make([]*jsonrpc.Request, 0)
		} else if index == len(block.Transactions) {
			_, err = c.RequestBatch(reqs, trsRaw)
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

func (c *ClientV3) MakeBlockWithReceipts(block *types.Block, trsArray []*TransactionResult) (*types.Block, error) {
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
					userStep = GetUserStep(tx.Operations[2].Account.Address, sd)
				}
				tx.Operations[2].Amount.Value = "-" + userStep.Mul(userStep, &sp.Int).Text(10)
			}
			HandleBugTransaction(tx, fa)
		}
		if trsArray[index].EventLogs != nil {
			ops := GetOperations(fa, trsArray[index].EventLogs, int64(len(tx.Operations))-1)
			tx.Operations = append(tx.Operations, ops...)
		}
		for i, op := range tx.Operations {
			op.Status = trsArray[index].StatusFlag
			if i >= FeeOpFromIndex {
				op.Status = types.String(SuccessStatus)
			}
		}
	}
	return block, nil
}

func (c *ClientV3) GetTransaction(param *TransactionRPCRequest) (*types.Transaction, error) {
	txRaw := map[string]interface{}{}
	id := time.Now().UnixNano() / int64(time.Millisecond)
	jrReq, err := GetRpcRequest("icx_getTransactionByHash", param, id)
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

func (c *ClientV3) GetTransactionResult(param *TransactionRPCRequest) (*TransactionResult, error) {
	trRaw := map[string]interface{}{}
	id := time.Now().UnixNano() / int64(time.Millisecond)
	jrReq, err := GetRpcRequest("icx_getTransactionResult", param, id)
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

func (c *ClientV3) MakeTransactionWithReceipt(tx *types.Transaction, txResult *TransactionResult) (*types.Transaction, error) {
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
				userStep = GetUserStep(tx.Operations[2].Account.Address, sd)
			}
			tx.Operations[2].Amount.Value = "-" + userStep.Mul(userStep, &sp.Int).Text(10)
			fa = tx.Operations[0].Account.Address
		}
	}
	if txResult.EventLogs != nil {
		ops := GetOperations(fa, txResult.EventLogs, int64(len(tx.Operations))-1)
		tx.Operations = append(tx.Operations, ops...)
	}
	for i, op := range tx.Operations {
		op.Status = txResult.StatusFlag
		if i >= FeeOpFromIndex {
			op.Status = types.String(SuccessStatus)
		}
	}
	return tx, nil
}

func (c *ClientV3) GetBalance(param *BalanceRPCRequest) (*common.HexInt, error) {
	id := time.Now().UnixNano() / int64(time.Millisecond)
	req, err := GetRpcRequest("icx_getBalance", param, id)
	if err != nil {
		return nil, err
	}
	balance := &common.HexInt{}
	if _, err := c.Request(req, &balance); err != nil {
		return nil, err
	}
	return balance, nil
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

	id := time.Now().UnixNano() / int64(time.Millisecond)
	jrReq, err := GetRpcRequest("icx_call", params, id)
	if err != nil {
		return nil, err
	}
	_, err = c.Request(jrReq, &resp)
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

	id := time.Now().UnixNano() / int64(time.Millisecond)
	jrReq, err := GetRpcRequest("icx_call", params, id)
	if err != nil {
		return nil, err
	}

	_, err = c.Request(jrReq, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *ClientV3) SendTransaction(req interface{}) error {
	resp := ""
	id := time.Now().UnixNano() / int64(time.Millisecond)
	jrReq, err := GetRpcRequest("icx_sendTransaction", req, id)
	if err != nil {
		return err
	}
	_, err = c.Request(jrReq, &resp)
	if err != nil {
		return err
	}
	return nil
}

func (c *ClientV3) GetStepDefaultStepCost() (*common.HexInt, error) {
	resp := map[string]*common.HexInt{}
	id := time.Now().UnixNano() / int64(time.Millisecond)
	params := map[string]interface{}{
		"to":       "cx0000000000000000000000000000000000000000",
		"dataType": "call",
		"data": map[string]interface{}{
			"method": "getStepCosts",
		},
	}
	jrReq, err := GetRpcRequest("icx_call", params, id)
	if err != nil {
		return nil, err
	}
	_, err = c.Request(jrReq, &resp)
	if err != nil {
		return nil, err
	}
	return resp["default"], nil
}

func GetUserStep(from string, stepDetails map[string]*common.HexInt) *big.Int {
	userUsed := new(big.Int)
	for f, v := range stepDetails {
		if from == f {
			userUsed.Set(&v.Int)
		}
	}
	return userUsed
}
