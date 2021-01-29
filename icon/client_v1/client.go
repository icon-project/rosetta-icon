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
	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/icon-project/goloop/common"
	"github.com/icon-project/goloop/server/jsonrpc"
	"math/big"
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
				ps[i+1] = "debug"
				ps = append(ps, "v3")
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

func (c *ClientV3) GetBlockReceipts(param *BlockRPCRequest) ([]*TransactionResult, error) {
	trsRaw := &[]interface{}{}

	_, err := c.Do("icx_getBlockReceipts", param, trsRaw)
	if err != nil {
		return nil, err
	}

	trsArray, err := ParseTransactionResults(trsRaw)
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
		if len(tx.Operations) >= 4 { //general tx(transfer, call, deploy...)
			su := trsArray[index].StepUsed
			sp := trsArray[index].StepPrice
			sd := trsArray[index].StepDetails
			if su.Cmp(zeroBigInt) != 0 {
				f := new(big.Int).Mul(&su.Int, &sp.Int)
				fee := f.Text(10)
				tx.Operations[3].Amount.Value = fee
				userStep := GetUserStep(tx.Operations[2].Account.Address, sd)
				tx.Operations[2].Amount.Value = "-" + userStep.Mul(userStep, &sp.Int).Text(10)
			}
			fa = tx.Operations[0].Account.Address
		}
		if trsArray[index].EventLogs != nil {
			ops := GetOperations(fa, trsArray[index].EventLogs, int64(len(tx.Operations))-1)
			tx.Operations = append(tx.Operations, ops...)
		}
		for _, op := range tx.Operations {
			if op.Type == TransferOpType {
				op.Status = trsArray[index].StatusFlag
			}
		}
	}
	return block, nil
}

func (c *ClientV3) GetTransaction(param *TransactionRPCRequest) (*types.Transaction, error) {
	txRaw := map[string]interface{}{}

	_, err := c.Do("icx_getTransactionByHash", param, &txRaw)
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

	_, err := c.Do("icx_getTransactionResult", param, &trRaw)
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
			userStep := GetUserStep(tx.Operations[2].Account.Address, sd)
			tx.Operations[2].Amount.Value = "-" + userStep.Mul(userStep, &sp.Int).Text(10)
		}
		fa = tx.Operations[0].Account.Address
	}
	if txResult.EventLogs != nil {
		ops := GetOperations(fa, txResult.EventLogs, int64(len(tx.Operations))-1)
		tx.Operations = append(tx.Operations, ops...)
	}
	for _, op := range tx.Operations {
		if op.Type == TransferOpType {
			op.Status = txResult.StatusFlag
		}
	}
	return tx, nil
}

func (c *ClientV3) GetBalance(param *BalanceRPCRequest) (*types.AccountBalanceResponse, error) {
	var debugAccount *DebugAccount
	var blk BalanceWithBlockId

	if _, blkErr := c.Do("icx_getLastBlock", nil, &blk); blkErr != nil {
		return nil, blkErr
	}

	if _, err := c.DoURL(c.DebugEndPoint, "debug_getAccount", param, &debugAccount); err != nil {
		return nil, err
	}

	return &types.AccountBalanceResponse{
		BlockIdentifier: &types.BlockIdentifier{
			Index: blk.Number(),
			Hash:  blk.Hash(),
		},
		Balances: []*types.Amount{
			{
				Value:    debugAccount.Balance(),
				Currency: ICXCurrency,
			},
		},
	}, nil
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

func (c *ClientV3) SendTransaction(req interface{}) error {
	resp := ""
	_, err := c.Do("icx_sendTransaction", req, &resp)
	if err != nil {
		return err
	}
	return nil
}

func (c *ClientV3) EstimateStep(req interface{}) (*Response, error) {
	resp := ""
	res, err := c.DoURL(c.DebugEndPoint, "debug_estimateStep", req, &resp)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetUserStep(from string, stepDetails map[string]*common.HexInt) *big.Int {
	userUsed := new(big.Int)
	if len(stepDetails) == 0 {
		return userUsed
	}
	for f, v := range stepDetails {
		if from == f {
			userUsed.Set(&v.Int)
		}
	}
	return userUsed
}
