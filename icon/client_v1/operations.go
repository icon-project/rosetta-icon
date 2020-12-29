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

func ParseOperationsV2(transaction Transaction) ([]*types.Operation, error) {
	var ops []*types.Operation

	dataType := transaction.GetDataType()
	dataType = CallOpType
	fromOp := &types.Operation{
		OperationIdentifier: &types.OperationIdentifier{
			Index: 0,
		},
		Type:   dataType,
		Status: SuccessStatus,
		Account: &types.AccountIdentifier{
			Address: transaction.FromAddr(),
		},
		Amount: &types.Amount{
			Value:    "-" + transaction.Values(),
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
		Status: SuccessStatus,
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
		Status: SuccessStatus,
		Account: &types.AccountIdentifier{
			Address: transaction.FromAddr(),
		},
		Amount: &types.Amount{
			Value:    "-" + transaction.Fee.Text(10),
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
		Status: SuccessStatus,
		Account: &types.AccountIdentifier{
			Address: TreasuryAddress,
		},
		Amount: &types.Amount{
			Value:    transaction.Fee.Text(10),
			Currency: ICXCurrency,
		},
	}
	ops = append(ops, feeToOp)
	return ops, nil
}

func ParseOperationsV3(transaction Transaction) ([]*types.Operation, error) {
	var ops []*types.Operation

	dataType := transaction.GetDataType()

	if dataType == BaseOpType {
		baseOp, _ := MakeBaseOperations()
		ops = append(ops, baseOp)
		return ops, nil
	}
	dataType = CallOpType

	fromOp := &types.Operation{
		OperationIdentifier: &types.OperationIdentifier{
			Index: 0,
		},
		Type:   dataType,
		Status: SuccessStatus,
		Account: &types.AccountIdentifier{
			Address: transaction.FromAddr(),
		},
		Amount: &types.Amount{
			Value:    "-" + transaction.Values(),
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
		Status: SuccessStatus,
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
		Status: SuccessStatus,
		Account: &types.AccountIdentifier{
			Address: transaction.FromAddr(),
		},
		Amount: &types.Amount{
			Value:    "-" + transaction.StepLimit.Text(10),
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
		Status: SuccessStatus,
		Account: &types.AccountIdentifier{
			Address: TreasuryAddress,
		},
		Amount: &types.Amount{
			Value:    transaction.StepLimit.Text(10),
			Currency: ICXCurrency,
		},
	}
	ops = append(ops, feeToOp)

	return ops, nil
}

func MakeBaseOperations() (*types.Operation, error) {
	baseOp := &types.Operation{
		OperationIdentifier: &types.OperationIdentifier{
			Index: 0,
		},
		Type:   BaseOpType,
		Status: SuccessStatus,
	}
	return baseOp, nil
}
