package main

import (
	"encoding/json"
	"fmt"
	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/icon-project/rosetta-icon/configuration"
	"github.com/icon-project/rosetta-icon/icon"
	"github.com/icon-project/rosetta-icon/icon/client_v1"
	"io"
	"os"
)

func main() {
	cfg, err := configuration.LoadConfiguration()
	if err != nil {
		_ = fmt.Errorf("Fail")
		return
	}

	rpcClient := icon.NewClient(cfg.URL, client_v1.ICXCurrency)

	//sampleGetBlock(rpcClient)
	//sampleGetPeer(rpcClient)
	sampleGetTransaction(rpcClient)
}


func sampleGetBlock(client *icon.Client) {
	// 이렇게 할당해야하는가?
	index := int64(5977334)
	params := &types.PartialBlockIdentifier{
		Index: &index,
	}

	block, err := client.GetBlock(params)
	err = JsonPrettyPrintln(os.Stdout, block)
	fmt.Print(err)
}

func sampleGetTransaction(client *icon.Client) {
	// 이렇게 할당해야하는가?

	params := &types.TransactionIdentifier{
		Hash: "0x2c89b69a75ce737ac61b76a6a86ffa233362ae7b05eabd54d067737e282b75c0",
	}

	tx, err := client.GetTransaction(params)
	err = JsonPrettyPrintln(os.Stdout, tx)
	fmt.Print(err)
}

func sampleGetPeer(client *icon.Client) {
	status, err := client.GetPeer()
	err = JsonPrettyPrintln(os.Stdout, status)
	fmt.Print(err)
}

func JsonPrettyPrintln(w io.Writer, v interface{}) error {
	b, err := json.MarshalIndent(v, "", "  ")
	_, err = fmt.Fprintln(w, string(b))
	return err
}
