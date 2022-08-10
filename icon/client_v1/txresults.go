package client_v1

import (
	"encoding/json"

	"github.com/coinbase/rosetta-sdk-go/types"
)

func ParseTransactionResults(trsRaws *[]interface{}) ([]*TransactionResult, error) {
	var trsArray []*TransactionResult
	for _, raw := range *trsRaws {
		txResult, err := ParseTransactionResult(raw)
		if err != nil {
			return nil, err
		}
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
