package tenderhelper

import (
	"github.com/tendermint/abci/client"
	"github.com/tendermint/abci/types"
)

var client abcicli.Client

// GetBlockAppHash uses the abcicli to get the last block app hash
func GetBlockAppHash() ([]byte, error) {
	resInfo, err := client.InfoSync(types.RequestInfo{})
	return resInfo.LastBlockAppHash, err
}
