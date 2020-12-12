package services

import (
	"encoding/json"
	"github.com/icon-project/goloop/common"
	"github.com/leeheonseung/rosetta-icon/icon/client_v1"
)

func marshalJSONMap(i interface{}) (map[string]interface{}, error) {
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

func unmarshalJSONMap(m map[string]interface{}, i interface{}) error {
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, i)
}

func mapNetwork(n string) *common.HexInt64 {
	switch n {
	case client_v1.MainnetNetwork:
		return &common.HexInt64{Value: 1}
	case client_v1.TestnetNetwork:
		return &common.HexInt64{Value: 2}
	default:
		return &common.HexInt64{Value: 3}
	}
}
