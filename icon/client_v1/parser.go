package client_v1

import (
	"encoding/json"
	"errors"
	"github.com/coinbase/rosetta-sdk-go/types"
)

func ParseBlock(raw map[string]interface{}) (*types.Block, error) {
	version := raw["version"]

	switch version {
	case "0.1a":
		blk := &Block01a{}
		if err := UnmarshalJSONMap(raw, blk); err != nil {
			return nil, err
		}
		return ParseBlock01a(blk)
	case "0.3", "0.4", "0.5":
		blk := &Block03{}
		if err := UnmarshalJSONMap(raw, blk); err != nil {
			return nil, err
		}
		return ParseBlock03(blk)
	}
	return nil, errors.New("Unsupported Block Version")
}

func ParseTransactions(txArray []json.RawMessage) ([]*types.Transaction, error) {
	var transactions []*types.Transaction

	for _, raw := range txArray {
		var tx *types.Transaction

		bs, _ := raw.MarshalJSON()
		transaction := Transaction{}
		if err := json.Unmarshal(bs, &transaction); err != nil {
			return nil, err
		}

		switch transaction.Version.String() {
		case "0x3":
			tx, _ = ParseTransactionV3(transaction)
		default:
			tx, _ = ParseTransactionV2(transaction)
		}
		transactions = append(transactions, tx)
	}
	return transactions, nil
}

func ParseTransactionResults(trsRaws *[]interface{}) ([]*TransactionResult, error) {
	var trsArray []*TransactionResult

	for _, raw := range *trsRaws {
		txResult, _ := ParseTransactionResult(raw)
		trsArray = append(trsArray, txResult)
	}
	return trsArray, nil
}

func ParseTransactionResult(tx interface{}) (*TransactionResult, error) {
	bs, _ := json.Marshal(tx)
	txResult := TransactionResult{}
	if err := json.Unmarshal(bs, &txResult); err != nil {
		return nil, err
	}
	bs, _ = txResult.Status.MarshalJSON()
	var status string
	if err := json.Unmarshal(bs, &status); err != nil {
		return nil, err
	}
	if status == "0x1" {
		txResult.StatusFlag = SuccessStatus
	} else {
		txResult.StatusFlag = FailureStatus
	}
	return &txResult, nil
}
