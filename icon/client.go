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
	RosettaTypes "github.com/coinbase/rosetta-sdk-go/types"
	"github.com/icon-project/goloop/common"
	"github.com/icon-project/rosetta-icon/icon/client_v1"
)

// Client is used to fetch blocks from ICON Node and
// to parser ICON block data into Rosetta types.
//
// We opted not to use existing ICON RPC libraries
// because they don't allow providing context
// in each request.

type Client struct {
	Currency *RosettaTypes.Currency
	IconV1   *client_v1.ClientV3
}

func NewClient(
	client *client_v1.ClientV3,
	currency *RosettaTypes.Currency,
) *Client {
	return &Client{
		currency,
		client,
	}
}

func (ic *Client) GetPeer() ([]*RosettaTypes.Peer, error) {
	return ic.IconV1.GetPeer()
}

func (ic *Client) SendTransaction(tx client_v1.Transaction) error {
	js, err := tx.ToJSON()
	if err != nil {
		return err
	}
	if err := ic.IconV1.SendTransaction(js); err != nil {
		return err
	}
	return nil
}

func (ic *Client) GetDefaultStepCost() (*common.HexInt, error) {
	res, err := ic.IconV1.GetStepDefaultStepCost()
	if err != nil {
		return nil, err
	}
	return res, nil
}
