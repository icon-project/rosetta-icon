package client_v1

import (
	"github.com/coinbase/rosetta-sdk-go/types"
	"math/big"
)

const (
	icxTransferSig   = "ICXTransfer(Address,Address,int)"
	issueSig         = "ICXIssued(int,int,int,int)"
	claimSig         = "IScoreClaimed(int,int)"
	claimSig2        = "IScoreClaimedV2(Address,int,int)"
	burnSig1         = "ICXBurned"
	burnSig2         = "ICXBurned(int)"
	burnSig3         = "ICXBurnedV2(Address,int,int)"
	depositWithdrawn = "DepositWithdrawn(bytes,Address,int,int)"
)

func ParseGenesisOperationsV2(tx GenesisTransaction) ([]*types.Operation, error) {
	var ops []*types.Operation
	for _, account := range tx.Accounts {
		accountOp := &types.Operation{
			OperationIdentifier: &types.OperationIdentifier{
				Index: int64(len(ops)),
			},
			Type:   GenesisOpType,
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
		Type:   MessageOpType,
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

	opType := TransferOpType
	fromOp := &types.Operation{
		OperationIdentifier: &types.OperationIdentifier{
			Index: 0,
		},
		Type:   opType,
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
		Type:   opType,
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

	if dataType == BaseDataType {
		baseOp, _ := MakeBaseOperations()
		ops = append(ops, baseOp)
		return ops, nil
	}
	opType := getOPType(dataType, transaction.To.String())

	fromOp := &types.Operation{
		OperationIdentifier: &types.OperationIdentifier{
			Index: 0,
		},
		Type:   opType,
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
		Type:   opType,
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

func GetOperations(fa string, els []*EventLog, lastOpIndex int64) []*types.Operation {
	ops := make([]*types.Operation, 0)
	for _, el := range els {
		switch *el.Indexed[0] {
		case icxTransferSig:
			value := new(big.Int)
			value.SetString((*el.Indexed[3])[2:], 16)
			ops = append(ops, &types.Operation{
				OperationIdentifier: &types.OperationIdentifier{
					Index: lastOpIndex + 1,
				},
				Type:   ICXTransferOpType,
				Status: SuccessStatus,
				Account: &types.AccountIdentifier{
					Address: *el.Indexed[1],
				},
				Amount: &types.Amount{
					Value:    "-" + value.Text(10),
					Currency: ICXCurrency,
				},
			})
			lastOpIndex += 1
			ops = append(ops, &types.Operation{
				OperationIdentifier: &types.OperationIdentifier{
					Index: lastOpIndex + 1,
				},
				RelatedOperations: []*types.OperationIdentifier{
					{
						Index: lastOpIndex,
					},
				},
				Type:   ICXTransferOpType,
				Status: SuccessStatus,
				Account: &types.AccountIdentifier{
					Address: *el.Indexed[2],
				},
				Amount: &types.Amount{
					Value:    value.Text(10),
					Currency: ICXCurrency,
				},
			})
			lastOpIndex += 1
		case issueSig:
			value := new(big.Int)
			value.SetString((*el.Data[2])[2:], 16)
			ops = append(ops, &types.Operation{
				OperationIdentifier: &types.OperationIdentifier{
					Index: lastOpIndex + 1,
				},
				Type:   IssueOpType,
				Status: SuccessStatus,
				Account: &types.AccountIdentifier{
					Address: TreasuryAddress,
				},
				Amount: &types.Amount{
					Value:    value.Text(10),
					Currency: ICXCurrency,
				},
			})
			lastOpIndex += 1
		case claimSig:
			op := getClaimOps(fa, el, lastOpIndex)
			ops = append(ops, op...)
			lastOpIndex += 2
		case claimSig2:
			op := getClaimOps(fa, el, lastOpIndex)
			ops = append(ops, op...)
			lastOpIndex += 2
		case burnSig1:
			op := getBurnOps(el, lastOpIndex)
			ops = append(ops, op)
			lastOpIndex += 1
		case burnSig2:
			op := getBurnOps(el, lastOpIndex)
			ops = append(ops, op)
			lastOpIndex += 1
		case burnSig3:
			op := getBurnOps(el, lastOpIndex)
			ops = append(ops, op)
			lastOpIndex += 1
		case depositWithdrawn:
			op := getDepositWithdrawn(el, lastOpIndex)
			ops = append(ops, op)
			lastOpIndex += 1
		}
	}
	return ops
}

func getClaimOps(fa string, el *EventLog, lastOpIndex int64) []*types.Operation {
	value := new(big.Int)
	value.SetString((*el.Data[1])[2:], 16)
	ops := make([]*types.Operation, 0)
	ops = append(ops, &types.Operation{
		OperationIdentifier: &types.OperationIdentifier{
			Index: lastOpIndex + 1,
		},
		Type:   ClaimOpType,
		Status: SuccessStatus,
		Account: &types.AccountIdentifier{
			Address: TreasuryAddress,
		},
		Amount: &types.Amount{
			Value:    "-" + value.Text(10),
			Currency: ICXCurrency,
		},
	})
	lastOpIndex += 1
	ops = append(ops, &types.Operation{
		OperationIdentifier: &types.OperationIdentifier{
			Index: lastOpIndex + 1,
		},
		RelatedOperations: []*types.OperationIdentifier{
			&types.OperationIdentifier{
				Index: lastOpIndex,
			},
		},
		Type:   ClaimOpType,
		Status: SuccessStatus,
		Account: &types.AccountIdentifier{
			Address: fa,
		},
		Amount: &types.Amount{
			Value:    value.Text(10),
			Currency: ICXCurrency,
		},
	})
	return ops
}

func getBurnOps(el *EventLog, lastOpIndex int64) *types.Operation {
	value := new(big.Int)
	value.SetString((*el.Data[0])[2:], 16)
	op := &types.Operation{
		OperationIdentifier: &types.OperationIdentifier{
			Index: lastOpIndex + 1,
		},
		Type:   BurnOpType,
		Status: SuccessStatus,
		Account: &types.AccountIdentifier{
			Address: SystemScoreAddress,
		},
		Amount: &types.Amount{
			Value:    "-" + value.Text(10),
			Currency: ICXCurrency,
		},
	}
	return op
}

func getDepositWithdrawn(el *EventLog, lastOpIndex int64) *types.Operation {
	value := new(big.Int)
	value.SetString((*el.Data[0])[2:], 16)
	fa := *el.Indexed[2]
	op := &types.Operation{
		OperationIdentifier: &types.OperationIdentifier{
			Index: lastOpIndex + 1,
		},
		Type:   WithdrawnType,
		Status: SuccessStatus,
		Account: &types.AccountIdentifier{
			Address: fa,
		},
		Amount: &types.Amount{
			Value:    value.Text(10),
			Currency: ICXCurrency,
		},
	}
	return op
}

func getOPType(dataType string, toAddress string) string {
	switch dataType {
	case DeployDataType:
		return DeployOpType
	case CallDataType:
		return CallOpType
	case TransferDataType:
		if IsContract(toAddress) {
			return CallOpType
		}
		return TransferOpType
	case MessageDataType:
		return MessageOpType
	case BaseDataType:
		return BaseOpType
	case DepositDataType:
		return DepositOpType
	default:
		return TransferOpType
	}
}
