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
	"context"
	"encoding/json"
	"fmt"
	"github.com/coinbase/rosetta-sdk-go/types"
	sdkUtils "github.com/coinbase/rosetta-sdk-go/utils"
	"github.com/icon-project/goloop/common"
	"github.com/icon-project/goloop/server/jsonrpc"
	"math/big"
	"net/http"
	"time"
)

const (
	retryLimit = 5
	retryDelay = 2
)

type ClientV3 struct {
	*JsonRpcClient
	genesisBlockIdentifier *types.BlockIdentifier
}

func NewClientV3(endpoint string, gbi *types.BlockIdentifier) *ClientV3 {
	client := new(http.Client)
	apiClient := NewJsonRpcClient(client, endpoint)

	return &ClientV3{
		JsonRpcClient: apiClient,
		genesisBlockIdentifier: gbi,
	}
}

func (c *ClientV3) getBlock(method string, param *BlockRPCRequest) (*types.Block, error) {
	blockRaw := map[string]interface{}{}
	id := time.Now().UnixNano() / int64(time.Millisecond)

	jrReq, err := GetRpcRequest(method, param, id)
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

func (c *ClientV3) GetBlock(ctx context.Context, param *BlockRPCRequest) (*types.Block, error) {
	var err error
	var block *types.Block

	for index := 0; index <= retryLimit; index++ {
		if block != nil {
			break
		}
		if param.Height == "" && param.Hash == "" {
			block, err = c.getBlock("icx_getLastBlock", param)
		} else if param.Hash != "" {
			block, err = c.getBlock("icx_getBlockByHash", param)
		} else {
			block, err = c.getBlock("icx_getBlockByHeight", param)
		}
		if err := sdkUtils.ContextSleep(ctx, retryDelay); err != nil {
			return nil, fmt.Errorf("%s: unable to get Block %+v", err, block.BlockIdentifier.Index)
		}
	}
	param = &BlockRPCRequest{Hash: "0x" + block.BlockIdentifier.Hash}
	for index := 0; index <= retryLimit; index++ {
		trsArray, err := c.GetReceipts(block)
		if err == nil {
			c.MakeBlockWithReceipts(block, trsArray)
			return block, nil
		}
		if err := sdkUtils.ContextSleep(ctx, retryDelay); err != nil {
			return nil, fmt.Errorf("%s: unable to get BlockReciept %+v", err, block.BlockIdentifier.Index)
		}
	}
	return nil, fmt.Errorf("%s: unable to get parsed block BH: %+v", err, param.Height)
}

func (c *ClientV3) GetReceipts(block *types.Block) ([]*TransactionResult, error) {
	var trsRaw []interface{}
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
		jrReq, err := GetRpcRequest("icx_getTransactionResult", reqParams, int64(index-1))
		if err != nil {
			return nil, err
		}
		reqs = append(reqs, jrReq)
		if mod == 0 {
			trsRaw = make([]interface{}, 10)
			_, err = c.RequestBatch(reqs, trsRaw)
			reqs = make([]*jsonrpc.Request, 0)
		} else if index == len(block.Transactions) {
			trsRaw = make([]interface{}, mod)
			_, err = c.RequestBatch(reqs, trsRaw)
		}
	}
	trsArray, err := ParseTransactionResults(&trsRaw)
	if err != nil {
		return nil, err
	}
	return trsArray, nil
}

func (c *ClientV3) MakeBlockWithReceipts(block *types.Block, trsArray []*TransactionResult) *types.Block {
	zeroBigInt := new(big.Int)
	fa := SystemScoreAddress
	success := SuccessStatus
	for index, tx := range block.Transactions {
		tx = block.Transactions[index]
		if tx.TransactionIdentifier.Hash == GenesisTxHash {
			continue
		}
		if len(tx.Operations) >= 4 { //general tx(transfer, call, deploy...)
			su := trsArray[index].StepUsed
			sp := trsArray[index].StepPrice
			sd := trsArray[index].StepDetails
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
		if trsArray[index].EventLogs != nil {
			ops := GetOperations(fa, trsArray[index].EventLogs, int64(len(tx.Operations))-1)
			tx.Operations = append(tx.Operations, ops...)
		}
		for i, op := range tx.Operations {
			op.Status = &trsArray[index].StatusFlag
			if i >= FeeOpFromIndex {
				op.Status = &success
			}
		}
	}
	return block
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
	success := SuccessStatus
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
		op.Status = &txResult.StatusFlag
		if i >= FeeOpFromIndex {
			op.Status = &success
		}
	}
	return tx, nil
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

func (c *ClientV3) NetworkStatus(ctx context.Context) (*types.NetworkStatusResponse, error) {
	block, err := c.GetBlock(ctx, &BlockRPCRequest{})
	if err != nil {
		return nil, fmt.Errorf("%w: unable to get current block", err)
	}

	peers, err := c.GetPeer()
	if err != nil {
		return nil, err
	}

	return &types.NetworkStatusResponse{
		CurrentBlockIdentifier: block.BlockIdentifier,
		CurrentBlockTimestamp:  block.Timestamp,
		GenesisBlockIdentifier: c.genesisBlockIdentifier,
		Peers:                  peers,
	}, nil
}

func (c *ClientV3) GetPeer() ([]*types.Peer, error) {
	resp, err := c.GetMainPReps()
	if err != nil {
		return nil, fmt.Errorf("%w: could not get peer", err)
	}

	var peers []*types.Peer
	preps := (*resp)["preps"]

	for _, element := range preps.([]interface{}) {
		address := element.(map[string]interface{})["address"]
		resp, _ := c.GetPRep(address.(string))
		peers = append(peers, &types.Peer{
			PeerID:   address.(string),
			Metadata: *resp,
		})
	}

	return peers, nil
}
