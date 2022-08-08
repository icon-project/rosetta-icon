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
)

var (
	// Errors contains all errors that could be returned
	// by this Rosetta implementation.
	Errors = []*types.Error{
		ErrUnimplemented,
		ErrUnavailableOffline,
		ErrUnableToGetStatus,
		ErrUnableToDecompressPubkey,
		ErrUnclearIntent,
		ErrUnableToParseIntermediateResult,
		ErrSignatureInvalid,
		ErrBroadcastFailed,
		ErrInvalidAddress,
		ErrWrongHashOrIndex,
		ErrUnableToGetBalance,
	}

	// ErrUnimplemented is returned when an endpoint
	// is called that is not implemented.
	ErrUnimplemented = &types.Error{
		Code:    0, //nolint
		Message: "Endpoint not implemented",
	}

	// ErrUnavailableOffline is returned when an endpoint
	// is called that is not available offline.
	ErrUnavailableOffline = &types.Error{
		Code:    1, //nolint
		Message: "Endpoint unavailable offline",
	}

	// ErrUnableToGetStatus is returned when failing to get
	// current network status
	ErrUnableToGetStatus = &types.Error{
		Code:    2,
		Message: "Unable to get network status",
	}

	// ErrUnableToDecompressPubkey is returned when
	// the *types.PublicKey provided in /construction/derive
	// cannot be decompressed.
	ErrUnableToDecompressPubkey = &types.Error{
		Code:    3, //nolint
		Message: "unable to decompress public key",
	}

	// ErrUnclearIntent is returned when operations
	// provided in /construction/preprocess or /construction/payloads
	// are not valid.
	ErrUnclearIntent = &types.Error{
		Code:    4, //nolint
		Message: "Unable to parse intent",
	}

	// ErrUnableToParseIntermediateResult is returned
	// when a data structure passed between Construction
	// API calls is not valid.
	ErrUnableToParseIntermediateResult = &types.Error{
		Code:    5, //nolint
		Message: "Unable to parse intermediate result",
	}

	// ErrSignatureInvalid is returned when a signature
	// cannot be parsed.
	ErrSignatureInvalid = &types.Error{
		Code:    6, //nolint
		Message: "Signature invalid",
	}

	// ErrBroadcastFailed is returned when transaction
	// broadcast fails.
	ErrBroadcastFailed = &types.Error{
		Code:    7, //nolint
		Message: "Unable to broadcast transaction",
	}

	// ErrInvalidAddress is returned when an address
	// is not valid.
	ErrInvalidAddress = &types.Error{
		Code:    12, //nolint
		Message: "Invalid address",
	}

	// ErrWrongHashOrIndex is returned when ICON Node cannot
	// process with the given data
	ErrWrongHashOrIndex = &types.Error{
		Code:      13,
		Message:   "Wrong hash or index",
		Retriable: true,
	}

	// ErrUnableToGetBalance is returned when it is not possible
	// to get the balance of a *types.AccountIdentifier.
	ErrUnableToGetBalance = &types.Error{
		Code:    14,
		Message: "Unable to get balance",
	}
)

// wrapErr adds details to the types.Error provided. We use a function
// to do this so that we don't accidentally overwrite the standard
// errors.
func wrapErr(rErr *types.Error, err error) *types.Error {
	newErr := &types.Error{
		Code:      rErr.Code,
		Message:   rErr.Message,
		Retriable: rErr.Retriable,
	}
	if err != nil {
		newErr.Details = map[string]interface{}{
			"context": err.Error(),
		}
	}

	return newErr
}
