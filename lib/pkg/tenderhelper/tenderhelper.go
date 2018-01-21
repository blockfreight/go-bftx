package tenderhelper

import (
	"github.com/tendermint/abci/client"
	"github.com/tendermint/abci/types"
)

// GetBlockAppHash uses the abcicli to get the last block app hash
func GetBlockAppHash(client abcicli.Client) ([]byte, error) {
	resInfo, err := client.InfoSync(types.RequestInfo{})
	return resInfo.LastBlockAppHash, err
}
