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

	for index, raw := range txArray {
		var tx *types.Transaction

		bs, _ := raw.MarshalJSON()
		transaction := Transaction{}
		if err := json.Unmarshal(bs, &transaction); err != nil {
			return nil, err
		}

		switch transaction.Version.String() {
		case "0x3":
			tx, _ = ParseTransactionV3(int64(index), transaction)
		default:
			tx, _ = ParseTransactionV2(int64(index), transaction)
		}
		transactions = append(transactions, tx)
	}
	return transactions, nil
}
