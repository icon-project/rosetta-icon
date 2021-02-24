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
	"github.com/icon-project/rosetta-icon/configuration"
	"github.com/icon-project/rosetta-icon/icon"
)

type BlockAPIService struct {
	config *configuration.Configuration
	client *icon.Client
}

func NewBlockAPIService(
	config *configuration.Configuration,
	client *icon.Client,
) *BlockAPIService {
	return &BlockAPIService{
		config: config,
		client: client,
	}
}

// Block implements the /block endpoint.
func (s *BlockAPIService) Block(
	ctx context.Context,
	request *types.BlockRequest,
) (*types.BlockResponse, *types.Error) {
	if s.config.Mode != configuration.Online {
		return nil, ErrUnavailableOffline
	}

	block, err := s.client.GetBlock(request.BlockIdentifier)
	if err != nil {
		return nil, wrapErr(ErrWrongBlockHash, err)
	}
	return &types.BlockResponse{
		Block: block,
	}, nil
}

// BlockTransaction implements the /block/transaction endpoint.
func (s *BlockAPIService) BlockTransaction(
	ctx context.Context,
	request *types.BlockTransactionRequest,
) (*types.BlockTransactionResponse, *types.Error) {
	if s.config.Mode != configuration.Online {
		return nil, ErrUnavailableOffline
	}

	tx, err := s.client.GetTransaction(request.TransactionIdentifier)
	if err != nil {
		return nil, wrapErr(ErrWrongBlockHash, err)
	}

	return &types.BlockTransactionResponse{
		Transaction: tx,
	}, nil
}
