package client_v1

import (
	"encoding/json"

	"github.com/coinbase/rosetta-sdk-go/types"
)

func ParseBlock(raw map[string]interface{}) (*types.Block, error) {
	blk := &Block01a{}
	if err := UnmarshalJSONMap(raw, blk); err != nil {
		return nil, err
	}
	return ParseBlock01a(blk)
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

		if transaction.Fee == nil {
			tx, _ = ParseTransactionV3(transaction)
		} else {
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
		txResult.StatusFlag = types.String(SuccessStatus)
	} else {
		txResult.StatusFlag = types.String(FailureStatus)
	}
	return &txResult, nil
}
