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
	"github.com/leeheonseung/rosetta-icon/icon"
)

type AccountAPIService struct {
	config *configuration.Configuration
	client *icon.Client
}

func NewAccountAPIService(
	config *configuration.Configuration,
	client *icon.Client,
) *AccountAPIService {
	return &AccountAPIService{
		config: config,
		client: client,
	}
}

// AccountBalance implements /account/balance.
func (s *AccountAPIService) AccountBalance(
	ctx context.Context,
	request *types.AccountBalanceRequest,
) (*types.AccountBalanceResponse, *types.Error) {
	if s.config.Mode != configuration.Online {
		return nil, ErrUnavailableOffline
	}

	// TODO Account 개발
	//balanceResponse, err := s.client.Balance(
	//	ctx,
	//	request.AccountIdentifier,
	//	request.BlockIdentifier,
	//)
	//if err != nil {
	//	return nil, wrapErr(ErrGeth, err)
	//}

	balance, err := s.client.GetBalance(request.AccountIdentifier)

	if err != nil {
		return nil, wrapErr(ErrInvalidAddress, err)
	}
	return balance, nil
}
