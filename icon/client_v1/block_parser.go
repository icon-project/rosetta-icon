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
	"errors"
	"github.com/coinbase/rosetta-sdk-go/types"
	"math/big"
)

func ParseBlock(raw map[string]interface{}) (*types.Block, error) {
	version := raw["version"]
	switch version {
	case "0.1a":
		return Block_0_1a(raw)
	case "0.3", "0.4", "0.5":
		return Block_0_3(raw)
	}
	return nil, errors.New("Unsupported Block Version")
}

func Block_0_1a(raw map[string]interface{}) (*types.Block, error) {
	meta := map[string]interface{}{
		"version":               raw["version"],
		"peer_id":               raw["peer_id"],
		"signature":             raw["signature"],
		"next_leader":           raw["next_leader"],
		"merkle_tree_root_hash": raw["merkle_tree_root_hash"],
	}

	index := int64(raw["height"].(float64))
	timestamp := int64(raw["time_stamp"].(float64)) // 1000

	if index == GenesisBlockIndex {
		txs, _ := ParseGenesisTransaction(raw["confirmed_transaction_list"].([]interface{}))
		return &types.Block{
			BlockIdentifier: &types.BlockIdentifier{
				Index: index,
				Hash:  raw["block_hash"].(string),
			},
			Timestamp:    timestamp,
			Transactions: txs,
			Metadata:     meta,
		}, nil
	} else {
		txs, _ := ParseTransactions(raw["confirmed_transaction_list"].([]interface{}))
		return &types.Block{
			BlockIdentifier: &types.BlockIdentifier{
				Index: index,
				Hash:  raw["block_hash"].(string),
			},
			ParentBlockIdentifier: &types.BlockIdentifier{
				Index: index - 1,
				Hash:  raw["prev_block_hash"].(string),
			},
			Timestamp:    timestamp,
			Transactions: txs,
			Metadata:     meta,
		}, nil
	}
}

func Block_0_3(raw map[string]interface{}) (*types.Block, error) {
	meta := map[string]interface{}{
		"version":          raw["version"],
		"transactionsHash": raw["transactionsHash"],
		"stateHash":        raw["stateHash"],
		"receiptsHash":     raw["receiptsHash"],
		"repsHash":         raw["repsHash"],
		"nextRepsHash":     raw["nextRepsHash"],
		"leaderVotesHash":  raw["leaderVotesHash"],
		"prevVotesHash":    raw["prevVotesHash"],
		"logsBloom":        raw["logsBloom"],
		"leaderVotes":      raw["leaderVotes"],
		"prevVotes":        raw["prevVotes"],
		"leader":           raw["leader"],
		"signature":        raw["signature"],
		"nextLeader":       raw["nextLeader"],
	}

	txs, _ := ParseTransactions(raw["confirmed_transaction_list"].([]interface{}))

	indexBig := new(big.Int)
	indexBig.SetString(raw["height"].(string)[2:], 16)
	index := indexBig.Int64()

	timestampBig := new(big.Int)
	timestampBig.SetString(raw["timestamp"].(string)[2:], 16) // 1000
	timestamp := timestampBig.Int64()

	if index == GenesisBlockIndex {
		return &types.Block{
			BlockIdentifier: &types.BlockIdentifier{
				Index: index,
				Hash:  raw["hash"].(string),
			},
			Timestamp:    timestamp,
			Transactions: txs,
			Metadata:     meta,
		}, nil
	} else {
		return &types.Block{
			BlockIdentifier: &types.BlockIdentifier{
				Index: index,
				Hash:  raw["hash"].(string),
			},
			ParentBlockIdentifier: &types.BlockIdentifier{
				Index: index - 1,
				Hash:  raw["prevHash"].(string),
			},
			Timestamp:    timestamp,
			Transactions: txs,
			Metadata:     meta,
		}, nil
	}
}

func ParseTransactions(raw []interface{}) ([]*types.Transaction, error) {
	var transactions []*types.Transaction

	for index, transaction := range raw {
		version := transaction.(map[string]interface{})["version"]
		switch version {
		case 3:
			return nil, nil
		default:
			tx, _ := ParseTransactionV2(int64(index), transaction.(map[string]interface{}))
			transactions = append(transactions, tx)
		}
		println(index, transaction)
	}
	return transactions, nil
}

func ParseGenesisTransaction(raw []interface{}) ([]*types.Transaction, error) {
	var transactions []*types.Transaction
	for index, transaction := range raw {
		operations, _ := ParseGenesisOperationsV2(transaction.(map[string]interface{}))
		tx := &types.Transaction{
			TransactionIdentifier: &types.TransactionIdentifier{
				Hash: "",
			},
			Operations: operations,
			Metadata:   nil,
		}
		transactions = append(transactions, tx)
		println(index, transaction)
	}
	return transactions, nil
}

func ParseTransactionV2(index int64, raw map[string]interface{}) (*types.Transaction, error) {
	meta := map[string]interface{}{
		"nonce":     raw["nonce"],
		"signature": raw["signature"],
		"method":    raw["method"],
	}
	operations, _ := ParseOperationsV2(index, raw)
	return &types.Transaction{
		TransactionIdentifier: &types.TransactionIdentifier{
			Hash: raw["tx_hash"].(string),
		},
		Operations: operations,
		Metadata:   meta,
	}, nil
}

func ParseGenesisOperationsV2(raw map[string]interface{}) ([]*types.Operation, error) {
	var ops []*types.Operation

	accounts := raw["accounts"].([]interface{})
	for _, element := range accounts {
		account := element.(map[string]interface{})

		accountOp := &types.Operation{
			OperationIdentifier: &types.OperationIdentifier{
				Index: int64(len(ops)),
			},
			Type:   CallOpType,
			Status: SuccessStatus,
			Account: &types.AccountIdentifier{
				Address: account["address"].(string),
			},
			Amount: &types.Amount{
				Value:    account["balance"].(string),
				Currency: ICXCurrency,
			},
			Metadata: map[string]interface{}{
				"name": account["name"].(string),
			},
		}
		ops = append(ops, accountOp)
	}

	messageOp := &types.Operation{
		OperationIdentifier: &types.OperationIdentifier{
			Index: int64(len(ops)),
		},
		Type:   CallOpType,
		Status: SuccessStatus,
		Metadata: map[string]interface{}{
			"message": raw["message"].(string),
		},
	}
	ops = append(ops, messageOp)
	return ops, nil
}

func ParseOperationsV2(startIndex int64, raw map[string]interface{}) ([]*types.Operation, error) {
	var ops []*types.Operation

	value := new(big.Int)
	value.SetString(raw["value"].(string)[2:], 16)

	fromOp := &types.Operation{
		OperationIdentifier: &types.OperationIdentifier{
			Index: int64(len(ops)) + startIndex,
		},
		Type:   CallOpType,
		Status: "",
		Account: &types.AccountIdentifier{
			Address: raw["from"].(string),
		},
		Amount: &types.Amount{
			Value:    new(big.Int).Neg(value).String(),
			Currency: ICXCurrency,
		},
	}
	ops = append(ops, fromOp)
	lastOpIndex := ops[len(ops)-1].OperationIdentifier.Index

	toOp := &types.Operation{
		OperationIdentifier: &types.OperationIdentifier{
			Index: lastOpIndex + 1,
		},
		RelatedOperations: []*types.OperationIdentifier{
			{
				Index: lastOpIndex,
			},
		},
		Type:   CallOpType,
		Status: "",
		Account: &types.AccountIdentifier{
			Address: raw["to"].(string),
		},
		Amount: &types.Amount{
			Value:    value.String(),
			Currency: ICXCurrency,
		},
	}
	ops = append(ops, toOp)
	lastOpIndex = ops[len(ops)-1].OperationIdentifier.Index

	fee := new(big.Int)
	fee.SetString(raw["fee"].(string)[2:], 16)

	feeFromOp := &types.Operation{
		OperationIdentifier: &types.OperationIdentifier{
			Index: lastOpIndex + 1,
		},
		RelatedOperations: []*types.OperationIdentifier{
			{
				Index: lastOpIndex,
			},
		},
		Type:   FeeOpType,
		Status: "",
		Account: &types.AccountIdentifier{
			Address: raw["from"].(string),
		},
		Amount: &types.Amount{
			Value:    new(big.Int).Neg(fee).String(),
			Currency: ICXCurrency,
		},
	}
	ops = append(ops, feeFromOp)
	lastOpIndex = ops[len(ops)-1].OperationIdentifier.Index

	feeToOp := &types.Operation{
		OperationIdentifier: &types.OperationIdentifier{
			Index: lastOpIndex + 1,
		},
		RelatedOperations: []*types.OperationIdentifier{
			{
				Index: lastOpIndex,
			},
		},
		Type:   FeeOpType,
		Status: "",
		Account: &types.AccountIdentifier{
			Address: TreasuryAddress,
		},
		Amount: &types.Amount{
			Value:    fee.String(),
			Currency: ICXCurrency,
		},
	}
	ops = append(ops, feeToOp)
	return ops, nil
}
