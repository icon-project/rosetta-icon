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

func MapNetwork(n string) *common.HexInt64 {
	switch n {
	case MainnetNetwork:
		return &common.HexInt64{Value: 1}
	case TestnetNetwork:
		return &common.HexInt64{Value: 2}
	case ZiconNetwork:
		return &common.HexInt64{Value: 3}
	case DevelopNetwork:
		return &common.HexInt64{Value: 80}
	default:
		return &common.HexInt64{Value: 3}
	}
}
