package types

import (
	data "github.com/tendermint/go-wire/data"
)

type ResponseBroadcast struct {
	// generic abci response
	JsonRPC string `json:"jsonrpc"`
	id      string `json:"id"`
	Result  resultBroadcastTx
}

type resultBroadcastTx struct {
	Code uint32     `json:"code"`
	Data data.Bytes `json:"data"`
	Log  string     `json:"log"`

	Hash data.Bytes `json:"hash"`
}
