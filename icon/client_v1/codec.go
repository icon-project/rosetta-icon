package client_v1

import (
	"encoding/json"
	"github.com/icon-project/goloop/common"
)

func MarshalJSONMap(i interface{}) (map[string]interface{}, error) {
	b, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}

	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}

	return m, nil
}

func UnmarshalJSONMap(m map[string]interface{}, i interface{}) error {
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, i)
}

func MapNetwork(n string) *common.HexInt {
	switch n {
	case MainnetNetwork:
		return common.NewHexInt(1)
	case TestnetNetwork:
		return common.NewHexInt(2)
	case DevelopNetwork:
		return common.NewHexInt(7)
	default:
		return common.NewHexInt(7)
	}
}
