package icon

import (
	"fmt"

	"github.com/icon-project/goloop/common"
)

const AddressStringLength = 42

func CheckAddress(as string) error {
	if len(as) != AddressStringLength {
		return fmt.Errorf("invalid address")
	}
	a := new(common.Address)
	if a.SetString(as) != nil {
		return fmt.Errorf("invalid address")
	}
	return nil
}

func IsContract(as string) bool {
	if as[:2] == "cx" {
		return true
	} else {
		return false
	}
}
