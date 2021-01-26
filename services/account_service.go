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
	"github.com/leeheonseung/rosetta-icon/configuration"
	"github.com/leeheonseung/rosetta-icon/icon/client_v1"
)

type AccountAPIService struct {
	config *configuration.Configuration
	i      Indexer
}

func NewAccountAPIService(
	config *configuration.Configuration,
	i Indexer,
) *AccountAPIService {
	return &AccountAPIService{
		config: config,
		i:      i,
	}
}

// AccountBalance implements /account/balance.
func (s *AccountAPIService) AccountBalance(
	ctx context.Context,
	request *types.AccountBalanceRequest,
) (*types.AccountBalanceResponse, *types.Error) {
	if s.config.Mode != configuration.Online {
		return nil, wrapErr(ErrUnavailableOffline, nil)
	}

	// TODO: filter balances by request currencies

	// If we are fetching a historical balance,
	// use balance storage and don't return coins.
	amount, block, err := s.i.GetBalance(
		ctx,
		request.AccountIdentifier,
		client_v1.ICXCurrency,
		request.BlockIdentifier,
	)
	if err != nil {
		return nil, wrapErr(ErrUnableToGetBalance, err)
	}

	return &types.AccountBalanceResponse{
		BlockIdentifier: block,
		Balances: []*types.Amount{
			amount,
		},
	}, nil
}

// AccountBalance implements /account/coin.
func (s *AccountAPIService) AccountCoins(ctx context.Context, request *types.AccountCoinsRequest) (*types.AccountCoinsResponse, *types.Error) {
	return nil, nil
}
