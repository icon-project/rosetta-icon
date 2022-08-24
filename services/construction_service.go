// Copyright 2020 Coinbase, Inc.
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

package services

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/coinbase/rosetta-sdk-go/parser"
	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/icon-project/goloop/common"
	"github.com/icon-project/goloop/common/crypto"
	"github.com/icon-project/rosetta-icon/configuration"
	"github.com/icon-project/rosetta-icon/icon"
)

// ConstructionAPIService implements the server.ConstructionAPIServicer interface.
type ConstructionAPIService struct {
	config *configuration.Configuration
	client Client
}

// NewConstructionAPIService creates a new instance of a ConstructionAPIService.
func NewConstructionAPIService(
	cfg *configuration.Configuration,
	client Client,
) *ConstructionAPIService {
	return &ConstructionAPIService{
		config: cfg,
		client: client,
	}
}

// ConstructionDerive implements the /construction/derive endpoint.
func (s *ConstructionAPIService) ConstructionDerive(
	ctx context.Context,
	request *types.ConstructionDeriveRequest,
) (*types.ConstructionDeriveResponse, *types.Error) {
	pubkey, err := crypto.ParsePublicKey(request.PublicKey.Bytes)
	if err != nil {
		return nil, wrapErr(ErrUnableToDecompressPubkey, err)
	}

	addr := common.NewAccountAddressFromPublicKey(pubkey)
	return &types.ConstructionDeriveResponse{
		AccountIdentifier: &types.AccountIdentifier{
			Address: addr.String(),
		},
	}, nil
}

// ConstructionPreprocess implements the /construction/preprocess
// endpoint.
func (s *ConstructionAPIService) ConstructionPreprocess(
	ctx context.Context,
	request *types.ConstructionPreprocessRequest,
) (*types.ConstructionPreprocessResponse, *types.Error) {
	descriptions := &parser.Descriptions{
		OperationDescriptions: []*parser.OperationDescription{
			{
				Type: icon.TransferOpType,
				Account: &parser.AccountDescription{
					Exists: true,
				},
				Amount: &parser.AmountDescription{
					Exists:   true,
					Sign:     parser.NegativeAmountSign,
					Currency: icon.ICXCurrency,
				},
			},
			{
				Type: icon.TransferOpType,
				Account: &parser.AccountDescription{
					Exists: true,
				},
				Amount: &parser.AmountDescription{
					Exists:   true,
					Sign:     parser.PositiveAmountSign,
					Currency: icon.ICXCurrency,
				},
			},
		},
		ErrUnmatched: true,
	}

	m, err := parser.MatchOperations(descriptions, request.Operations)
	if err != nil {
		return nil, wrapErr(ErrUnclearIntent, err)
	}

	f, _ := m[0].First()
	fa := f.Account.Address
	t, _ := m[1].First()
	ta := t.Account.Address

	// Ensure valid from address
	e := icon.CheckAddress(fa)
	if e != nil {
		return nil, wrapErr(ErrInvalidAddress, fmt.Errorf("%s is not a valid address", fa))
	}

	// Ensure valid to address
	e = icon.CheckAddress(ta)
	if e != nil {
		return nil, wrapErr(ErrInvalidAddress, fmt.Errorf("%s is not a valid address", ta))
	}

	preprocessOutput := &options{
		From: fa,
	}

	marshaled, err := icon.MarshalJSONMap(preprocessOutput)
	if err != nil {
		return nil, wrapErr(ErrUnableToParseIntermediateResult, err)
	}
	return &types.ConstructionPreprocessResponse{
		Options:            marshaled,
		RequiredPublicKeys: []*types.AccountIdentifier{{Address: fa}},
	}, nil
}

// ConstructionMetadata implements the /construction/metadata endpoint.
func (s *ConstructionAPIService) ConstructionMetadata(
	ctx context.Context,
	request *types.ConstructionMetadataRequest,
) (*types.ConstructionMetadataResponse, *types.Error) {
	if s.config.Mode != configuration.Online {
		return nil, ErrUnavailableOffline
	}

	var input options
	if err := icon.UnmarshalJSONMap(request.Options, &input); err != nil {
		return nil, wrapErr(ErrUnableToParseIntermediateResult, err)
	}

	res, err := s.client.GetDefaultStepCost()
	if err != nil {
		return nil, wrapErr(ErrUnableToParseIntermediateResult, err)
	}

	metadata := &metadata{
		res,
	}

	metadataMap, err := icon.MarshalJSONMap(metadata)
	if err != nil {
		return nil, wrapErr(ErrUnableToParseIntermediateResult, err)
	}

	return &types.ConstructionMetadataResponse{
		Metadata: metadataMap,
	}, nil
}

// ConstructionPayloads implements the /construction/payloads endpoint.
func (s *ConstructionAPIService) ConstructionPayloads(
	ctx context.Context,
	request *types.ConstructionPayloadsRequest,
) (*types.ConstructionPayloadsResponse, *types.Error) {
	d := &parser.Descriptions{
		OperationDescriptions: []*parser.OperationDescription{
			{
				Type: icon.TransferOpType,
				Account: &parser.AccountDescription{
					Exists: true,
				},
				Amount: &parser.AmountDescription{
					Exists:   true,
					Sign:     parser.NegativeAmountSign,
					Currency: icon.ICXCurrency,
				},
			},
			{
				Type: icon.TransferOpType,
				Account: &parser.AccountDescription{
					Exists: true,
				},
				Amount: &parser.AmountDescription{
					Exists:   true,
					Sign:     parser.PositiveAmountSign,
					Currency: icon.ICXCurrency,
				},
			},
		},
		ErrUnmatched: true,
	}
	m, err := parser.MatchOperations(d, request.Operations)
	if err != nil {
		return nil, wrapErr(ErrUnclearIntent, err)
	}

	// Required Fields for constructing a ICON transaction
	tOp, amount := m[1].First()
	ta := tOp.Account.Address
	nid := icon.MapNetwork(s.config.Network.Network)

	// stepLimit
	bs, err := json.Marshal(request.Metadata)
	if err != nil {
		return nil, wrapErr(ErrUnclearIntent, err)
	}
	var meta metadata
	err = json.Unmarshal(bs, &meta)
	if err != nil {
		return nil, wrapErr(ErrUnclearIntent, err)
	}

	// Additional Fields for constructing custom ICON tx struct
	fOp, _ := m[0].First()
	fa := fOp.Account.Address
	uTx := &icon.Transaction{
		Version:   common.HexUint16{Value: 3},
		From:      *common.MustNewAddressFromString(fa),
		To:        *common.MustNewAddressFromString(ta),
		Value:     &common.HexInt{Int: *amount},
		StepLimit: *meta.DefaultStepCost,
		Timestamp: common.HexInt64{Value: time.Now().UnixNano() / int64(time.Microsecond)},
		NID:       nid,
		Nonce:     common.NewHexInt(1),
	}

	h, err := uTx.CalcHash()
	if err != nil {
		return nil, ErrUnclearIntent
	}

	payload := &types.SigningPayload{
		AccountIdentifier: &types.AccountIdentifier{Address: fa},
		Bytes:             h,
		SignatureType:     types.EcdsaRecovery,
	}

	unsignedTxJSON, err := json.Marshal(uTx)
	if err != nil {
		return nil, wrapErr(ErrUnableToParseIntermediateResult, err)
	}

	return &types.ConstructionPayloadsResponse{
		UnsignedTransaction: string(unsignedTxJSON),
		Payloads:            []*types.SigningPayload{payload},
	}, nil
}

// ConstructionCombine implements the /construction/combine
// endpoint.
func (s *ConstructionAPIService) ConstructionCombine(
	ctx context.Context,
	request *types.ConstructionCombineRequest,
) (*types.ConstructionCombineResponse, *types.Error) {
	var unsignedTx icon.Transaction
	var err error
	if err = json.Unmarshal([]byte(request.UnsignedTransaction), &unsignedTx); err != nil {
		return nil, wrapErr(ErrUnableToParseIntermediateResult, err)
	}

	var sig common.Signature
	sigBytes := request.Signatures[0].Bytes
	sig.Signature, err = crypto.ParseSignature(sigBytes)
	if err != nil {
		return nil, wrapErr(ErrSignatureInvalid, err)
	}

	signedTx := unsignedTx
	signedTx.Signature = &sig

	if err = signedTx.VerifySignature(); err != nil {
		return nil, wrapErr(ErrUnableToParseIntermediateResult, err)
	}

	signedTxJSON, err := json.Marshal(signedTx)
	if err != nil {
		return nil, wrapErr(ErrUnableToParseIntermediateResult, err)
	}

	return &types.ConstructionCombineResponse{
		SignedTransaction: string(signedTxJSON),
	}, nil
}

// ConstructionHash implements the /construction/hash endpoint.
func (s *ConstructionAPIService) ConstructionHash(
	ctx context.Context,
	request *types.ConstructionHashRequest,
) (*types.TransactionIdentifierResponse, *types.Error) {
	signedTx := &icon.Transaction{}
	if err := json.Unmarshal([]byte(request.SignedTransaction), signedTx); err != nil {
		return nil, wrapErr(ErrUnableToParseIntermediateResult, err)
	}
	h := "0x" + hex.EncodeToString(signedTx.TxHash())

	return &types.TransactionIdentifierResponse{
		TransactionIdentifier: &types.TransactionIdentifier{
			Hash: h,
		},
	}, nil
}

// ConstructionParse implements the /construction/parse endpoint.
func (s *ConstructionAPIService) ConstructionParse(
	ctx context.Context,
	request *types.ConstructionParseRequest,
) (*types.ConstructionParseResponse, *types.Error) {
	var tx icon.Transaction
	if !request.Signed {
		err := json.Unmarshal([]byte(request.Transaction), &tx)
		if err != nil {
			return nil, wrapErr(ErrUnableToParseIntermediateResult, err)
		}
	} else {
		err := json.Unmarshal([]byte(request.Transaction), &tx)
		if err != nil {
			return nil, wrapErr(ErrUnableToParseIntermediateResult, err)
		}

	}

	err := icon.CheckAddress(tx.From.String())
	if err != nil {
		return nil, wrapErr(ErrInvalidAddress, fmt.Errorf("%s is not a valid address", tx.From.String()))
	}

	err = icon.CheckAddress(tx.From.String())
	if err != nil {
		return nil, wrapErr(ErrInvalidAddress, fmt.Errorf("%s is not a valid address", tx.To))
	}

	ops := []*types.Operation{
		{
			Type: icon.TransferOpType,
			OperationIdentifier: &types.OperationIdentifier{
				Index: 0,
			},
			Account: &types.AccountIdentifier{
				Address: tx.From.String(),
			},
			Amount: &types.Amount{
				Value:    "-" + tx.Values(),
				Currency: icon.ICXCurrency,
			},
		},
		{
			Type: icon.TransferOpType,
			OperationIdentifier: &types.OperationIdentifier{
				Index: 1,
			},
			RelatedOperations: []*types.OperationIdentifier{
				{
					Index: 0,
				},
			},
			Account: &types.AccountIdentifier{
				Address: tx.To.String(),
			},
			Amount: &types.Amount{
				Value:    tx.Values(),
				Currency: icon.ICXCurrency,
			},
		},
	}

	var resp *types.ConstructionParseResponse
	if request.Signed {
		resp = &types.ConstructionParseResponse{
			Operations: ops,
			AccountIdentifierSigners: []*types.AccountIdentifier{
				{
					Address: tx.From.String(),
				},
			},
		}
	} else {
		resp = &types.ConstructionParseResponse{
			Operations:               ops,
			AccountIdentifierSigners: []*types.AccountIdentifier{},
		}
	}
	return resp, nil
}

// ConstructionSubmit implements the /construction/submit endpoint.
func (s *ConstructionAPIService) ConstructionSubmit(
	ctx context.Context,
	request *types.ConstructionSubmitRequest,
) (*types.TransactionIdentifierResponse, *types.Error) {
	if s.config.Mode != configuration.Online {
		return nil, ErrUnavailableOffline
	}

	signedTx, err := icon.ParseV3JSON([]byte(request.SignedTransaction))
	if err != nil {
		return nil, wrapErr(ErrUnableToParseIntermediateResult, err)
	}

	if err := s.client.SendTransaction(*signedTx); err != nil {
		return nil, wrapErr(ErrBroadcastFailed, err)
	}

	h := "0x" + hex.EncodeToString(signedTx.TxHash())
	txIdentifier := &types.TransactionIdentifier{
		Hash: h,
	}

	return &types.TransactionIdentifierResponse{
		TransactionIdentifier: txIdentifier,
	}, nil
}
