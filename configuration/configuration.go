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
	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/leeheonseung/rosetta-icon/icon/client_v1"
	"os"
	"strconv"
	"strings"
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

	// Testnet is ICON Testnet3.
	Testnet string = "TESTNET"

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
	DefaultEndPoint = "http://localhost:9000"

	EndpointPrefix        = "api"
	EndpointVersionPrefix = "v3"
	DebugPrefix           = "debug"

	// PortEnv is the environment variable
	// read to determine the port for the Rosetta
	// implementation.
	PortEnv = "PORT"
)

// Configuration determines how
type Configuration struct {
	Mode     Mode
	Network  *types.NetworkIdentifier
	URL      string
	DebugURL string
	Port     int
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
			Blockchain: client_v1.Blockchain,
			Network:    client_v1.MainnetNetwork,
		}
	case Testnet:
		config.Network = &types.NetworkIdentifier{
			Blockchain: client_v1.Blockchain,
			Network:    client_v1.TestnetNetwork,
		}
	case "":
		return nil, errors.New("NETWORK must be populated")
	default:
		return nil, fmt.Errorf("%s is not a valid network", networkValue)
	}

	envEndpoint := os.Getenv(EndpointEnv)
	if len(envEndpoint) > 0 {
		url := []string{
			envEndpoint,
			EndpointPrefix,
			EndpointVersionPrefix,
		}
		config.URL = strings.Join(url, "/")

		debugUrl := []string{
			envEndpoint,
			EndpointPrefix,
			DebugPrefix,
			EndpointVersionPrefix,
		}
		config.DebugURL = strings.Join(debugUrl, "/")
	} else {
		url := []string{
			DefaultEndPoint,
			EndpointPrefix,
			EndpointVersionPrefix,
		}
		config.URL = strings.Join(url, "/")

		debugUrl := []string{
			DefaultEndPoint,
			EndpointPrefix,
			DebugPrefix,
			EndpointVersionPrefix,
		}
		config.DebugURL = strings.Join(debugUrl, "/")
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
