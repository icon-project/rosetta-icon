package icon

import (
	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/icon-project/goloop/common"
)

const AddressStringLength = 42

func CheckAddress(as string) *types.Error {
	if len(as) != AddressStringLength {
		return ErrInvalidAddress
	}
	a := new(common.Address)
	if a.SetString(as) != nil {
		return ErrInvalidAddress
	}
	return nil
}
