package blockchain

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	rpc "github.com/tendermint/tendermint/rpc/client"
)

type Blockchain struct {
	LastBlockAppHash []byte
	Data             string
}

func GetInfo() (Blockchain, error) {
	var result Blockchain
	rpcClient := rpc.NewHTTP(os.Getenv("LOCAL_RPC_CLIENT_ADDRESS"), "/websocket")
	err := rpcClient.Start()
	if err != nil {
		fmt.Println("Error when initializing rpcClient")
		fmt.Println(err.Error())
		return result, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}
	defer rpcClient.Stop()

	abciInfo, err := rpcClient.ABCIInfo()
	if err != nil {
		fmt.Println("Error when initializing rpcClient")
		fmt.Println(err.Error())
		return result, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}

	result.Data = abciInfo.Response.Data
	result.LastBlockAppHash = abciInfo.Response.LastBlockAppHash

	return result, nil
}
