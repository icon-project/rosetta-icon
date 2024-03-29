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

package icon

import (
	"bytes"
	"encoding/json"

	"github.com/coinbase/rosetta-sdk-go/asserter"
	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/icon-project/goloop/common"
	"github.com/icon-project/goloop/common/crypto"
	"github.com/icon-project/goloop/service/transaction"
)

var (
	ICXCurrency = &types.Currency{
		Symbol:   ICXSymbol,
		Decimals: ICXDecimals,
	}

	OperationTypes = []string{
		GenesisOpType,
		TransferOpType,
		FeeOpType,
		IssueOpType,
		BurnOpType,
		LostOpType,
		FSDepositOpType,
		FSWithdrawOpType,
		FSFeeOpType,
		StakeOpType,
		UnstakeOpType,
		ClaimOpType,
		GhostOpType,
		RewardOpType,
		RegPRepOpType,
		MessageOpType,
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
)

const (
	// NodeVersion is the version of goloop we are using.
	NodeVersion = "1.2.13"

	// Blockchain is ICON.
	Blockchain string = "ICON"

	// MainnetNetwork is the value of the network
	// in MainnetNetworkIdentifier.
	MainnetNetwork string = "Mainnet"

	// LisbonNetwork is the value of the network
	// in LisbonNetworkIdentifier.
	LisbonNetwork string = "Lisbon"

	// BerlinNetwork is the value of the network
	// in BerlinNetworkIdentifier.
	BerlinNetwork string = "Berlin"

	// LocalNetwork is the value of the network
	// in LocalNetworkIdentifier.
	LocalNetwork string = "Localnet"

	EndpointPrefix  = "api"
	EndpointVersion = "v3"
	EndpointAdmin   = "admin"
	EndpointRosetta = "rosetta"

	ICXSymbol   = "ICX"
	ICXDecimals = 18

	GenesisBlockIndex          = int64(0)
	HistoricalBalanceSupported = true
	IncludeMempoolCoins        = false

	TreasuryAddress    = "hx1000000000000000000000000000000000000000"
	SystemScoreAddress = "cx0000000000000000000000000000000000000000"

	GenesisOpType    = "GENESIS"
	TransferOpType   = "TRANSFER"
	FeeOpType        = "FEE"
	IssueOpType      = "ISSUE"
	BurnOpType       = "BURN"
	LostOpType       = "LOST"
	FSDepositOpType  = "FS_DEPOSIT"
	FSWithdrawOpType = "FS_WITHDRAW"
	FSFeeOpType      = "FS_FEE"
	StakeOpType      = "STAKE"
	UnstakeOpType    = "UNSTAKE"
	ClaimOpType      = "CLAIM"
	GhostOpType      = "GHOST"
	RewardOpType     = "REWARD"
	RegPRepOpType    = "REG_PREP"
	MessageOpType    = "MESSAGE"

	BaseOpType        = "BASE"
	WithdrawnType     = "WITHDRAWN"
	ICXTransferOpType = "ICXTRANSFER"
	DeployOpType      = "DEPLOY"
	CallOpType        = "CALL"
	DepositOpType     = "DEPOSIT"
	BugOpType         = "BUG"

	BaseDataType     = "base"
	TransferDataType = "transfer"
	MessageDataType  = "message"
	DeployDataType   = "deploy"
	CallDataType     = "call"
	DepositDataType  = "deposit"

	SuccessStatus = "SUCCESS"
	FailureStatus = "FAIL"

	GenesisTxHash = "0x0000000000000000000000000000000000000000000000000000000000000000"
)

type BlockRPCRequest struct {
	Hash   string `json:"hash,omitempty"`
	Height string `json:"height,omitempty"`
}

type TransactionRPCRequest struct {
	Hash string `json:"txHash"`
}

type BalanceRPCRequest struct {
	Address string `json:"address"`
	Height  string `json:"height,omitempty"`
}

type Block struct {
	BlockHash          common.HexBytes   `json:"block_hash"`
	Version            string            `json:"version"`
	Height             int64             `json:"height"`
	Timestamp          int64             `json:"time_stamp"`
	Proposer           common.Address    `json:"peer_id"`
	PrevBlockHash      common.HexBytes   `json:"prev_block_hash"`
	MerkleTreeRootHash common.HexBytes   `json:"merkle_tree_root_hash"`
	Signature          common.HexBytes   `json:"signature"`
	Transactions       []json.RawMessage `json:"confirmed_transaction_list"`
}

func (b Block) Number() int64 {
	return b.Height
}

func (b Block) Hash() string {
	return b.BlockHash.String()
}

func (b Block) TimestampInMillis() int64 {
	return b.Timestamp / 1000
}

func (b Block) PrevHash() string {
	return b.PrevBlockHash.String()
}

func (b Block) GenesisMeta() map[string]interface{} {
	return map[string]interface{}{
		"version":               b.Version,
		"peer_id":               b.Proposer.String(),
		"merkle_tree_root_hash": b.MerkleTreeRootHash,
	}
}

func (b Block) Meta() map[string]interface{} {
	return map[string]interface{}{
		"version":               b.Version,
		"peer_id":               b.Proposer.String(),
		"merkle_tree_root_hash": b.MerkleTreeRootHash,
	}
}

type GenesisAccount struct {
	Name    string         `json:"name"`
	Address common.Address `json:"address"`
	Balance *common.HexInt `json:"balance"`
}

func (ga *GenesisAccount) Addr() string {
	return ga.Address.String()
}

func (ga *GenesisAccount) Balances() string {
	return ga.Balance.Text(10)
}

type GenesisTransaction struct {
	Accounts []GenesisAccount `json:"accounts"`
	Message  string           `json:"message"`
}

type Transaction struct {
	Version   common.HexUint16  `json:"version"`
	From      common.Address    `json:"from"`
	To        common.Address    `json:"to"`
	Value     *common.HexInt    `json:"value,omitempty"`
	StepLimit common.HexInt     `json:"stepLimit"`
	Timestamp common.HexInt64   `json:"timestamp"`
	NID       *common.HexInt    `json:"nid"`
	Nonce     *common.HexInt    `json:"nonce,omitempty"`
	Signature *common.Signature `json:"signature,omitempty"`
	DataType  *string           `json:"dataType,omitempty"`
	Data      json.RawMessage   `json:"data,omitempty"`
	Fee       *common.HexInt    `json:"fee,omitempty"`
	TxHashV3  common.HexBytes   `json:"txHash,omitempty"`
	TxHashV2  common.HexBytes   `json:"tx_hash,omitempty"`
	Method    string            `json:"method,omitempty"`
}

func (tx *Transaction) Values() string {
	if tx.Value != nil {
		return tx.Value.Text(10)
	} else {
		return "0"
	}
}

func (tx *Transaction) FromAddr() string {
	return tx.From.String()
}

func (tx *Transaction) ToAddr() string {
	return tx.To.String()
}

func (tx *Transaction) FeeValue() string {
	return tx.Fee.String()
}

func (tx *Transaction) StepValues() string {
	return tx.StepLimit.Text(10)
}

func (tx *Transaction) MetaV2() map[string]interface{} {
	return map[string]interface{}{
		"nonce":     &tx.Nonce,
		"signature": &tx.Signature,
		"method":    tx.Method,
	}
}

func (tx *Transaction) MetaV3() map[string]interface{} {
	meta := map[string]interface{}{
		"version":   tx.Version,
		"timestamp": tx.Timestamp,
		"data":      tx.Data,
		"dataType":  &tx.DataType,
	}

	if tx.GetDataType() == "Base" {
		return meta
	} else {
		meta["nid"] = &tx.NID
		meta["nonce"] = &tx.Nonce
		meta["signature"] = &tx.Signature
		return meta
	}
}

func (tx *Transaction) GetDataType() string {
	defaultType := [5]string{"call", "deploy", "message", "base", "deposit"}

	if tx.DataType != nil {
		for _, dataType := range defaultType {
			if *tx.DataType == dataType {
				return dataType
			}
		}
	}
	return "transfer"
}

func (tx *Transaction) ToJSON() (map[string]interface{}, error) {
	jso := map[string]interface{}{
		"version":   &tx.Version,
		"from":      &tx.From,
		"to":        &tx.To,
		"stepLimit": &tx.StepLimit,
		"timestamp": &tx.Timestamp,
		"signature": &tx.Signature,
	}
	if tx.Value != nil {
		jso["value"] = tx.Value
	}
	if tx.NID != nil {
		jso["nid"] = tx.NID
	}
	if tx.Nonce != nil {
		jso["nonce"] = tx.Nonce
	}
	if tx.DataType != nil {
		jso["dataType"] = *tx.DataType
	}
	if tx.Data != nil {
		jso["data"] = tx.Data
	}
	return jso, nil
}

func (tx *Transaction) CalcHash() ([]byte, error) {
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

func (tx *Transaction) VerifySignature() error {
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

func (tx *Transaction) TxHash() []byte {
	h, err := tx.CalcHash()
	if err != nil {
		return []byte{}
	}
	return h
}

func ParseV3JSON(js []byte) (*Transaction, error) {
	tx := new(Transaction)

	if err := json.Unmarshal(js, tx); err != nil {
		return nil, transaction.InvalidFormat.Wrapf(err, "Invalid json for transactionV3(%s)", string(js))
	}
	return tx, nil
}

type EventLog struct {
	Addr    string    `json:"scoreAddress"`
	Indexed []*string `json:"indexed"`
	Data    []*string `json:"data"`
}

type TransactionResult struct {
	StatusFlag         *string
	Status             json.RawMessage           `json:"status"`
	BlockHeight        *json.RawMessage          `json:"blockHeight"`
	BlockHash          *json.RawMessage          `json:"blockHash"`
	TxHash             *json.RawMessage          `json:"txHash"`
	TxIndex            *json.RawMessage          `json:"txIndex"`
	To                 *json.RawMessage          `json:"to"`
	ScoreAddress       *json.RawMessage          `json:"scoreAddress"`
	StepUsed           *common.HexInt            `json:"stepUsed"`
	CumulativeStepUsed *common.HexInt            `json:"cumulativeStepUsed"`
	StepPrice          *common.HexInt            `json:"stepPrice"`
	LogsBloom          *json.RawMessage          `json:"logsBloom"`
	EventLogs          []*EventLog               `json:"eventLogs"`
	Failure            *json.RawMessage          `json:"failure"`
	StepDetails        map[string]*common.HexInt `json:"stepUsedDetails"`
}

type RosettaTraceParam struct {
	Tx     string `json:"tx,omitempty"`
	Block  string `json:"block,omitempty"`
	Height string `json:"height,omitempty"`
}

type RosettaTraceResponse struct {
	BlockHash      string           `json:"blockHash"`
	PrevBlockHash  string           `json:"prevBlockHash"`
	BlockHeight    common.HexInt64  `json:"blockHeight"`
	Timestamp      common.HexInt64  `json:"timestamp"`
	BalanceChanges []*BalanceChange `json:"balanceChanges"`
}

type BalanceChange struct {
	TxHash  string       `json:"txHash"`
	TxIndex string       `json:"txIndex"`
	Ops     []*Operation `json:"ops"`
}

type Operation struct {
	OpType string        `json:"opType"`
	From   string        `json:"from"`
	To     string        `json:"to"`
	Amount common.HexInt `json:"amount"`
}

func (rt RosettaTraceResponse) TimestampInMillis() int64 {
	if rt.Timestamp.Value == 0 {
		return asserter.MinUnixEpoch
	}
	return rt.Timestamp.Value / 1000
}

func (rt RosettaTraceResponse) Index() int64 {
	return rt.BlockHeight.Value
}

func (op Operation) IntValue() string {
	return op.Amount.Text(10)
}
