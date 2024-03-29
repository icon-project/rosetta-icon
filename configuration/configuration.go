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

package configuration

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/icon-project/rosetta-icon/icon"
)

// Mode is the setting that determines if
// the implementation is "online" or "offline".
type Mode string

const (
	// Online is when the implementation is permitted
	// to make outbound connections.
	Online Mode = "ONLINE"

	// Offline is when the implementation is not permitted
	// to make outbound connections.
	Offline Mode = "OFFLINE"

	// Mainnet is the ICON Mainnet.
	Mainnet string = "MAINNET"

	// Lisbon is the ICON Lisbon testnet.
	Lisbon string = "LISBON"

	// Berlin is the ICON Berlin testnet.
	Berlin string = "BERLIN"

	// Localnet is the ICON local testnet.
	Localnet string = "LOCALNET"

	// ModeEnv is the environment variable read
	// to determine mode.
	ModeEnv = "MODE"

	// NetworkEnv is the environment variable
	// read to determine network.
	NetworkEnv = "NETWORK"

	// EndpointEnv is the environment variable
	// read to determine endpoint.
	EndpointEnv = "ENDPOINT"

	// DefaultEndPoint is the default endpoint for a running node.
	DefaultEndPoint = "http://localhost:9080"

	// PortEnv is the environment variable
	// read to determine the port for the Rosetta
	// implementation.
	PortEnv = "PORT"

	// MiddlewareVersion is the version of rosetta-icon
	MiddlewareVersion = "0.0.4"
)

// Configuration determines how
type Configuration struct {
	Mode         Mode
	Network      *types.NetworkIdentifier
	GenesisBlock *types.BlockIdentifier
	Endpoint     string
	Port         int
}

// LoadConfiguration attempts to create a new Configuration
// using the ENVs in the environment.
func LoadConfiguration() (*Configuration, error) {
	config := &Configuration{}

	modeValue := Mode(os.Getenv(ModeEnv))
	switch modeValue {
	case Online:
		config.Mode = Online
	case Offline:
		config.Mode = Offline
	case "":
		return nil, errors.New("MODE must be populated")
	default:
		return nil, fmt.Errorf("%s is not a valid mode", modeValue)
	}

	networkValue := os.Getenv(NetworkEnv)
	switch networkValue {
	case Mainnet:
		config.Network = &types.NetworkIdentifier{
			Blockchain: icon.Blockchain,
			Network:    icon.MainnetNetwork,
		}
	case Lisbon:
		config.Network = &types.NetworkIdentifier{
			Blockchain: icon.Blockchain,
			Network:    icon.LisbonNetwork,
		}
	case Berlin:
		config.Network = &types.NetworkIdentifier{
			Blockchain: icon.Blockchain,
			Network:    icon.BerlinNetwork,
		}
	case Localnet:
		config.Network = &types.NetworkIdentifier{
			Blockchain: icon.Blockchain,
			Network:    icon.LocalNetwork,
		}
	case "":
		return nil, errors.New("NETWORK must be populated")
	default:
		return nil, fmt.Errorf("%s is not a valid network", networkValue)
	}

	envEndpoint := os.Getenv(EndpointEnv)
	if len(envEndpoint) > 0 {
		config.Endpoint = envEndpoint
	} else {
		config.Endpoint = DefaultEndPoint
	}

	envPort := os.Getenv(PortEnv)
	if len(envPort) == 0 {
		return nil, errors.New("PORT must be populated")
	}

	port, err := strconv.Atoi(envPort)
	if err != nil || len(envPort) == 0 || port <= 0 {
		return nil, fmt.Errorf("%w: unable to parser port %s", err, envPort)
	}
	config.Port = port

	return config, nil
}
