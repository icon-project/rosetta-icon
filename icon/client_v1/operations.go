package client_v1

import (
	"github.com/coinbase/rosetta-sdk-go/types"
)

func ParseGenesisOperationsV2(tx GenesisTransaction) ([]*types.Operation, error) {
	var ops []*types.Operation
	for _, account := range tx.Accounts {
		accountOp := &types.Operation{
			OperationIdentifier: &types.OperationIdentifier{
				Index: int64(len(ops)),
			},
			Type:   CallOpType,
			Status: SuccessStatus,
			Account: &types.AccountIdentifier{
				Address: account.Addr(),
			},
			Amount: &types.Amount{
				Value:    account.Balances(),
				Currency: ICXCurrency,
			},
			Metadata: map[string]interface{}{
				"name": account.Name,
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
			"message": tx.Message,
		},
	}
	ops = append(ops, messageOp)
	return ops, nil
}

func ParseOperationsV2(startIndex int64, transaction Transaction) ([]*types.Operation, error) {
	var ops []*types.Operation

	dataType := transaction.GetDataType()

	fromOp := &types.Operation{
		OperationIdentifier: &types.OperationIdentifier{
			Index: int64(len(ops)) + startIndex,
		},
		Type:   dataType,
		Status: "",
		Account: &types.AccountIdentifier{
			Address: transaction.FromAddr(),
		},
		Amount: &types.Amount{
			Value:    transaction.Values(),
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
			&types.OperationIdentifier{
				Index: lastOpIndex,
			},
		},
		Type:   dataType,
		Status: "",
		Account: &types.AccountIdentifier{
			Address: transaction.ToAddr(),
		},
		Amount: &types.Amount{
			Value:    transaction.Values(),
			Currency: ICXCurrency,
		},
	}
	ops = append(ops, toOp)
	lastOpIndex = ops[len(ops)-1].OperationIdentifier.Index

	feeFromOp := &types.Operation{
		OperationIdentifier: &types.OperationIdentifier{
			Index: lastOpIndex + 1,
		},
		RelatedOperations: []*types.OperationIdentifier{
			&types.OperationIdentifier{
				Index: lastOpIndex,
			},
		},
		Type:   FeeOpType,
		Status: "",
		Account: &types.AccountIdentifier{
			Address: transaction.FromAddr(),
		},
		Amount: &types.Amount{
			Value:    transaction.FeeValue(),
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
			&types.OperationIdentifier{
				Index: lastOpIndex,
			},
		},
		Type:   FeeOpType,
		Status: "",
		Account: &types.AccountIdentifier{
			Address: TreasuryAddress,
		},
		Amount: &types.Amount{
			Value:    transaction.Fee.String(),
			Currency: ICXCurrency,
		},
	}
	ops = append(ops, feeToOp)
	return ops, nil
}

func ParseOperationsV3(startIndex int64, transaction Transaction) ([]*types.Operation, error) {
	var ops []*types.Operation

	dataType := transaction.GetDataType()

	if dataType == "Base" {
		baseOp, _ := MakeBaseOperations(startIndex, transaction.GetDataType(), int64(len(ops)))
		ops = append(ops, baseOp)
		return ops, nil
	}

	fromOp := &types.Operation{
		OperationIdentifier: &types.OperationIdentifier{
			Index: int64(len(ops)) + startIndex,
		},
		Type:   dataType,
		Status: "",
		Account: &types.AccountIdentifier{
			Address: transaction.FromAddr(),
		},
		Amount: &types.Amount{
			Value:    transaction.Values(),
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
		Type:   dataType,
		Status: "",
		Account: &types.AccountIdentifier{
			Address: transaction.ToAddr(),
		},
		Amount: &types.Amount{
			Value:    transaction.Values(),
			Currency: ICXCurrency,
		},
	}

	ops = append(ops, toOp)
	lastOpIndex = ops[len(ops)-1].OperationIdentifier.Index

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
			Address: transaction.FromAddr(),
		},
		Amount: &types.Amount{
			Value:    transaction.StepValue(),
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
			Value:    transaction.StepValue(),
			Currency: ICXCurrency,
		},
	}
	ops = append(ops, feeToOp)

	return ops, nil
}

func MakeBaseOperations(startIndex int64, dataType string, size int64) (*types.Operation, error) {
	baseOp := &types.Operation{
		OperationIdentifier: &types.OperationIdentifier{
			Index: size + startIndex,
		},
		Type:   dataType,
		Status: "",
	}
	return baseOp, nil
}
