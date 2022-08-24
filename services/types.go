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

package services

import (
	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/icon-project/goloop/common"
	"github.com/icon-project/rosetta-icon/icon"
)

// Client is used by the services to get block
// data and to submit transactions.
type Client interface {
	Status() (*types.BlockIdentifier, int64, []*types.Peer, error)

	GetBlock(
		identifier *types.PartialBlockIdentifier,
	) (*types.Block, error)

	GetTransaction(
		identifier *types.TransactionIdentifier,
	) (*types.Transaction, error)

	GetBalance(
		account *types.AccountIdentifier,
		block *types.PartialBlockIdentifier,
	) (*types.AccountBalanceResponse, error)

	GetDefaultStepCost() (*common.HexInt, error)

	SendTransaction(
		tx icon.Transaction,
	) error
}

type options struct {
	From string `json:"from"`
}

type metadata struct {
	DefaultStepCost *common.HexInt `json:"default_step_cost"`
}
