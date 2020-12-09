// Copyright 2020 ICON Foundation, Inc.
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

package client_v1

import (
	"bytes"
	"encoding/json"
	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/icon-project/goloop/common"
	"github.com/icon-project/goloop/common/crypto"
	"github.com/icon-project/goloop/service/transaction"
	"math/big"
	"time"
)

var (
	ICXCurrency = &types.Currency{
		Symbol:   ICXSymbol,
		Decimals: ICXDecimals,
	}

	OperationTypes = []string{
		"TEST",
	}

	// OperationStatuses are all supported operation statuses.
	OperationStatuses = []*types.OperationStatus{
		{
			Status:     SuccessStatus,
			Successful: true,
		},
		{
			Status:     FailureStatus,
			Successful: false,
		},
	}

	MiddlewareVersion = "0.0.1"
	RosettaVersion    = "1.4.0"
	NodeVersion       = "1.8.0"
	StepPrice         = big.NewInt(100000)
	TransferStepPrice = big.NewInt(12500000000)
	TransferStepLimit = big.NewInt(1000000000000000)
)

const (
	// Blockchain is ICON.
	Blockchain string = "ICON"

	// MainnetNetwork is the value of the network
	// in MainnetNetworkIdentifier.
	MainnetNetwork string = "Mainnet"

	// TestnetNetwork is the value of the network
	// in TestnetNetworkIdentifier.
	TestnetNetwork string = "Testnet"

	ICXSymbol   = "ICX"
	ICXDecimals = 18

	GenesisBlockIndex          = int64(0)
	HistoricalBalanceSupported = false

	TreasuryAddress = "hx1000000000000000000000000000000000000000"

	CallOpType = "CALL"
	FeeOpType  = "FEE"

	SuccessStatus = "SUCCESS"
	FailureStatus = "FAIL"
)

type BlockRPCRequest struct {
	Hash   string `json:"hash,omitempty"`
	Height string `json:"height,omitempty"`
}

type TransactionRPCRequest struct {
	Hash string `json:"Hash"`
}

type TransactionV3 struct {
	Version   common.HexUint16 `json:"version"`
	From      common.Address   `json:"from"`
	To        common.Address   `json:"to"`
	Value     *common.HexInt   `json:"value,omitempty"`
	StepLimit common.HexInt    `json:"stepLimit"`
	Timestamp common.HexInt64  `json:"timestamp"`
	NID       *common.HexInt64 `json:"nid"`
	Nonce     *common.HexInt   `json:"nonce,omitempty"`
	DataType  *string          `json:"dataType,omitempty"`
	Data      json.RawMessage  `json:"data,omitempty"`
}

func (tx *TransactionV3) CalcHash() ([]byte, error) {
	// sha := sha3.New256()
	sha := bytes.NewBuffer(nil)
	sha.Write([]byte("icx_sendTransaction"))

	// data
	if tx.Data != nil {
		sha.Write([]byte(".data."))
		if len(tx.Data) > 0 {
			var obj interface{}
			if err := json.Unmarshal(tx.Data, &obj); err != nil {
				return nil, err
			}
			if bs, err := transaction.SerializeValue(obj); err != nil {
				return nil, err
			} else {
				sha.Write(bs)
			}
		}
	}

	// dataType
	if tx.DataType != nil {
		sha.Write([]byte(".dataType."))
		sha.Write([]byte(*tx.DataType))
	}

	// from
	sha.Write([]byte(".from."))
	sha.Write([]byte(tx.From.String()))

	// nid
	if tx.NID != nil {
		sha.Write([]byte(".nid."))
		sha.Write([]byte(tx.NID.String()))
	}

	// nonce
	if tx.Nonce != nil {
		sha.Write([]byte(".nonce."))
		sha.Write([]byte(tx.Nonce.String()))
	}

	// stepLimit
	sha.Write([]byte(".stepLimit."))
	sha.Write([]byte(tx.StepLimit.String()))

	// timestamp
	sha.Write([]byte(".timestamp."))
	sha.Write([]byte(tx.Timestamp.String()))

	// to
	sha.Write([]byte(".to."))
	sha.Write([]byte(tx.To.String()))

	// value
	if tx.Value != nil {
		sha.Write([]byte(".value."))
		sha.Write([]byte(tx.Value.String()))
	}

	// version
	sha.Write([]byte(".version."))
	sha.Write([]byte(tx.Version.String()))

	return crypto.SHA3Sum256(sha.Bytes()), nil
}

type TransactionV3WithSig struct {
	TransactionV3
	Signature common.Signature `json:"signature"`
	Hash      []byte           `json:"txHash"`
}

func (tx *TransactionV3WithSig) VerifySignature() error {
	pk, err := tx.Signature.RecoverPublicKey(tx.TxHash())
	if err != nil {
		return transaction.InvalidSignatureError.Wrap(err, "fail to recover public key")
	}
	addr := common.NewAccountAddressFromPublicKey(pk)
	if tx.From.Equal(addr) {
		return nil
	}
	return transaction.InvalidSignatureError.New("fail to verify signature")
}

func (tx *TransactionV3WithSig) TxHash() []byte {
	if tx.Hash == nil {
		h, err := tx.CalcHash()
		if err != nil {
			tx.Hash = []byte{}
		} else {
			tx.Hash = h
		}
	}
	return tx.Hash
}

func (tx *TransactionV3WithSig) UnmarshalJSON(input []byte) error {
	if err := json.Unmarshal(input, tx); err != nil {
		return err
	}
	if tx.VerifySignature() != nil {
		return transaction.InvalidSignatureError.New("failed to verify signature")
	}
	tx.Timestamp = common.HexInt64{Value: time.Now().UTC().UnixNano()}
	return nil
}

func ParseV3JSON(js []byte, raw bool) (*TransactionV3, error) {
	tx := new(TransactionV3)

	if err := json.Unmarshal(js, tx); err != nil {
		return nil, transaction.InvalidFormat.Wrapf(err, "Invalid json for transactionV3(%s)", string(js))
	}
	return tx, nil
}
