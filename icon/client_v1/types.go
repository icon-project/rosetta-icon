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
	"math/big"

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
		BaseOpType,
		IssueOpType,
		BurnOpType,
		ICXTransferOpType,
		ClaimOpType,
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
	NodeVersion = "1.2.10"

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

	ICXSymbol   = "ICX"
	ICXDecimals = 18

	GenesisBlockIndex          = int64(0)
	HistoricalBalanceSupported = false
	IncludeMempoolCoins        = false

	TreasuryAddress    = "hx1000000000000000000000000000000000000000"
	SystemScoreAddress = "cx0000000000000000000000000000000000000000"

	GenesisOpType     = "GENESIS"
	TransferOpType    = "TRANSFER"
	FeeOpType         = "FEE"
	BaseOpType        = "BASE"
	BurnOpType        = "BURN"
	WithdrawnType     = "WITHDRAWN"
	ICXTransferOpType = "ICXTRANSFER"
	ClaimOpType       = "CLAIM"
	IssueOpType       = "ISSUE"
	MessageOpType     = "MESSAGE"
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

	FeeOpFromIndex = 2
	GenesisTxHash  = "0x0000000000000000000000000000000000000000000000000000000000000000"
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
}

type Block01a struct {
	ID                 common.HexBytes   `json:"block_hash"`
	Version            string            `json:"version"`
	Height             common.HexInt64   `json:"height"`
	Timestamp          common.HexInt64   `json:"time_stamp"`
	Proposer           common.Address    `json:"peer_id"`
	PrevID             common.HexBytes   `json:"prev_block_hash"`
	MerkleTreeRootHash common.HexBytes   `json:"merkle_tree_root_hash"`
	NextLeader         common.Address    `json:"next_leader"`
	Transactions       []json.RawMessage `json:"confirmed_transaction_list" `
}

func (b *Block01a) Number() int64 {
	return b.Height.Value
}

func (b *Block01a) Hash() string {
	return b.ID.String()
}

func (b *Block01a) PrevHash() string {
	return b.PrevID.String()
}

func (b *Block01a) Time() int64 {
	return b.Timestamp.Value
}

func (b *Block01a) TimestampMilli() int64 {
	return b.Timestamp.Value / 1000
}

func (b *Block01a) GenesisMeta() map[string]interface{} {
	return map[string]interface{}{
		"version":               b.Version,
		"peer_id":               b.Proposer,
		"signature":             "",
		"next_leader":           b.NextLeader,
		"merkle_tree_root_hash": b.MerkleTreeRootHash,
	}
}

func (b *Block01a) Meta() map[string]interface{} {
	return map[string]interface{}{
		"version":               b.Version,
		"peer_id":               b.Proposer,
		"next_leader":           b.NextLeader,
		"merkle_tree_root_hash": b.MerkleTreeRootHash,
	}
}

type Block03 struct {
	ID               common.HexBytes   `json:"hash"`
	Version          string            `json:"version"`
	Height           common.HexInt64   `json:"height"`
	Timestamp        common.HexInt64   `json:"timestamp"`
	Leader           common.Address    `json:"leader"`
	PrevID           common.HexBytes   `json:"prevHash"`
	TransactionsHash common.HexBytes   `json:"transactionsHash"`
	NextLeader       common.Address    `json:"nextLeader"`
	Transactions     []json.RawMessage `json:"transactions"`
	StateHash        common.HexBytes   `json:"stateHash"`
	ReceiptsHash     common.HexBytes   `json:"receiptsHash"`
	RepsHash         common.HexBytes   `json:"repsHash"`
	NextRepsHash     common.HexBytes   `json:"nextRepsHash"`
	LeaderVotesHash  common.HexBytes   `json:"leaderVotesHash"`
	PrevVotesHash    common.HexBytes   `json:"prevVotesHash"`
	LogsBloom        common.HexBytes   `json:"logsBloom"`
}

func (b *Block03) Number() int64 {
	return b.Height.Value
}

func (b *Block03) Hash() string {
	return b.ID.String()
}

func (b *Block03) Time() int64 {
	return b.Timestamp.Value
}

func (b *Block03) TimestampMilli() int64 {
	return b.Timestamp.Value / 1000
}

func (b *Block03) PrevHash() string {
	return b.PrevID.String()
}

func (b *Block03) Meta() map[string]interface{} {
	return map[string]interface{}{
		"version":          b.Version,
		"transactionsHash": b.TransactionsHash,
		"stateHash":        b.StateHash,
		"receiptsHash":     b.ReceiptsHash,
		"repsHash":         b.RepsHash,
		"nextRepsHash":     b.NextRepsHash,
		"leaderVotesHash":  b.LeaderVotesHash,
		"prevVotesHash":    b.PrevVotesHash,
		"logsBloom":        b.LogsBloom,
		"leader":           b.Leader,
		"nextLeader":       b.NextLeader,
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

type BalanceWithBlockId struct {
	ID     common.HexBytes `json:"block_hash"`
	Height common.HexInt64 `json:"height"`
}

func (info *BalanceWithBlockId) Hash() string {
	return info.ID.String()
}

func (info *BalanceWithBlockId) Number() int64 {
	return info.Height.Value
}

type StakeInfo struct {
	Stake    *common.HexInt `json:"stake"`
	UnStakes []*Unstake     `json:"unstakes,omitempty"`
}

type Unstake struct {
	Value  *common.HexInt `json:"unstake"`
	Expire int64          `json:"unstakeBlockHeight"`
	Remain int64          `json:"remainingBlocks"`
}

func (da *StakeInfo) Total() *big.Int {
	totalStake := new(big.Int)
	for _, u := range da.UnStakes {
		totalStake.Add(totalStake, &u.Value.Int)
	}
	if da.Stake == nil {
		da.Stake = common.NewHexInt(0)
	}
	totalStake.Add(totalStake, &da.Stake.Int)
	return totalStake
}
