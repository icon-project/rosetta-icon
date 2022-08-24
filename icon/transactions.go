package icon

import (
	"encoding/json"

	"github.com/coinbase/rosetta-sdk-go/types"
)

func ParseGenesisTransaction(txArray []json.RawMessage) ([]*types.Transaction, error) {
	var transactions []*types.Transaction
	for _, raw := range txArray {
		bs, _ := raw.MarshalJSON()
		genesisTransaction := GenesisTransaction{}
		if err := json.Unmarshal(bs, &genesisTransaction); err != nil {
			return nil, err
		}
		operations, _ := ParseGenesisOperationsV2(genesisTransaction)
		tx := &types.Transaction{
			TransactionIdentifier: &types.TransactionIdentifier{
				Hash: "0x0000000000000000000000000000000000000000000000000000000000000000",
			},
			Operations: operations,
			Metadata:   nil,
		}
		transactions = append(transactions, tx)
	}
	return transactions, nil
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

func ParseTransactionV2(transaction Transaction) (*types.Transaction, error) {
	operations, _ := ParseOperationsV2(transaction)
	return &types.Transaction{
		TransactionIdentifier: &types.TransactionIdentifier{
			Hash: transaction.TxHashV2.String(),
		},
		Operations: operations,
		Metadata:   transaction.MetaV2(),
	}, nil
}

func ParseTransactionV3(transaction Transaction) (*types.Transaction, error) {
	operations, _ := ParseOperationsV3(transaction)
	return &types.Transaction{
		TransactionIdentifier: &types.TransactionIdentifier{
			Hash: transaction.TxHashV3.String(),
		},
		Operations: operations,
		Metadata:   transaction.MetaV3(),
	}, nil
}
