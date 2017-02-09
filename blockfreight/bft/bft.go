package bft

import (
	"encoding/hex"
	"strings"

	"github.com/tendermint/abci/types"
	. "github.com/tendermint/go-common"
	"github.com/tendermint/go-merkle"
	"github.com/tendermint/go-wire"
)

type BftApplication struct {
	state merkle.Tree
}

func NewBftApplication() *BftApplication {
	state := merkle.NewIAVLTree(0, nil)
	return &BftApplication{state: state}
}

func (app *BftApplication) Info() (resInfo types.ResponseInfo) {
	return types.ResponseInfo{Data: Fmt("{\"size\":%v}", app.state.Size())}
}

func (app *BftApplication) SetOption(key string, value string) (log string) {
	return ""
}

// tx is either "key=value" or just arbitrary bytes
func (app *BftApplication) DeliverTx(tx []byte) types.Result {
	parts := strings.Split(string(tx), "=")
	if len(parts) == 2 {
		app.state.Set([]byte(parts[0]), []byte(parts[1]))
	} else {
		app.state.Set(tx, tx)
	}
	return types.OK
}

func (app *BftApplication) CheckTx(tx []byte) types.Result {
	return types.OK
}

func (app *BftApplication) Commit() types.Result {
	hash := app.state.Hash()
	return types.NewResultOK(hash, "")
}

func (app *BftApplication) Query(query []byte) types.Result {
	index, value, exists := app.state.Get(query)
	queryResult := QueryResult{index, string(value), hex.EncodeToString(value), exists}
	return types.NewResultOK(wire.JSONBytes(queryResult), "")
}

type QueryResult struct {
	Index    int    `json:"index"`
	Value    string `json:"value"`
	ValueHex string `json:"valueHex"`
	Exists   bool   `json:"exists"`
}
