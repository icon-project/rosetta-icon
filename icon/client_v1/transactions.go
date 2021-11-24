package client_v1

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
				Hash: "0000000000000000000000000000000000000000000000000000000000000000",
			},
			Operations: operations,
			Metadata:   nil,
		}
		transactions = append(transactions, tx)
	}
	return transactions, nil
}

func ParseTransactionV2(transaction Transaction) (*types.Transaction, error) {
	operations, _ := ParseOperationsV2(transaction)
	return &types.Transaction{
		TransactionIdentifier: &types.TransactionIdentifier{
			Hash: transaction.TxHashV2.String()[2:],
		},
		Operations: operations,
		Metadata:   transaction.MetaV2(),
	}, nil
}

func ParseTransactionV3(transaction Transaction) (*types.Transaction, error) {
	operations, _ := ParseOperationsV3(transaction)

	return &types.Transaction{
		TransactionIdentifier: &types.TransactionIdentifier{
			Hash: transaction.TxHashV3.String()[2:],
		},
		Operations: operations,
		Metadata:   transaction.MetaV3(),
	}, nil
}
