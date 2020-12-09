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
	"context"
	"github.com/coinbase/rosetta-sdk-go/types"
	"math/big"
)

// Client is used by the servicers to get block
// data and to submit transactions.
type Client interface {
	GetBlock(
		context.Context,
		*types.PartialBlockIdentifier,
	) (*types.Block, error)
}

type options struct {
	From string `json:"from"`
}

type metadata struct {
	StepPrice *big.Int `json:"stepPrice"`
}
