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
	"github.com/icon-project/rosetta-icon/icon/client_v1"
)

type NetworkAPIService struct {
	config *configuration.Configuration
	client *icon.Client
}

func NewNetworkAPIService(
	config *configuration.Configuration,
	client *icon.Client,
) *NetworkAPIService {
	return &NetworkAPIService{
		config: config,
		client: client,
	}
}

// NetworkList implements the /network/list endpoint
func (s *NetworkAPIService) NetworkList(
	ctx context.Context,
	request *types.MetadataRequest,
) (*types.NetworkListResponse, *types.Error) {
	return &types.NetworkListResponse{
		NetworkIdentifiers: []*types.NetworkIdentifier{
			s.config.Network,
		},
	}, nil
}

// NetworkOptions implements the /network/options endpoint.
func (s *NetworkAPIService) NetworkOptions(
	ctx context.Context,
	request *types.NetworkRequest,
) (*types.NetworkOptionsResponse, *types.Error) {
	return &types.NetworkOptionsResponse{
		Version: &types.Version{
			NodeVersion:       client_v1.NodeVersion,
			RosettaVersion:    types.RosettaAPIVersion,
			MiddlewareVersion: types.String(configuration.MiddlewareVersion),
		},
		Allow: &types.Allow{
			Errors:                  Errors,
			OperationTypes:          client_v1.OperationTypes,
			OperationStatuses:       client_v1.OperationStatuses,
			HistoricalBalanceLookup: client_v1.HistoricalBalanceSupported,
		},
	}, nil
}

// NetworkStatus implements the /network/status endpoint.
func (s *NetworkAPIService) NetworkStatus(
	ctx context.Context,
	request *types.NetworkRequest,
) (*types.NetworkStatusResponse, *types.Error) {
	if s.config.Mode != configuration.Online {
		return nil, wrapErr(ErrUnavailableOffline, nil)
	}

	gh := int64(0)
	params := &types.PartialBlockIdentifier{
		Index: &gh,
	}
	genesisBlock, err := s.client.GetBlock(params)
	if err != nil {
		return nil, wrapErr(ErrWrongBlockHash, err)
	}

	params = &types.PartialBlockIdentifier{}
	lastBlock, err := s.client.GetBlock(params)
	if err != nil {
		return nil, wrapErr(ErrWrongBlockHash, err)
	}

	peers, err := s.client.GetPeer()
	return &types.NetworkStatusResponse{
		CurrentBlockIdentifier: lastBlock.BlockIdentifier,
		CurrentBlockTimestamp:  lastBlock.Timestamp,
		GenesisBlockIdentifier: genesisBlock.BlockIdentifier,
		SyncStatus:             nil,
		Peers:                  peers,
	}, nil
}
