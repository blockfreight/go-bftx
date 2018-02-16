package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	rpc "github.com/tendermint/tendermint/rpc/client"
)

func GetInfo() (interface{}, error) {
	rpcClient := rpc.NewHTTP(os.Getenv("LOCAL_RPC_CLIENT_ADDRESS"), "/websocket")
	err := rpcClient.Start()
	if err != nil {
		fmt.Println("Error when initializing rpcClient")
		fmt.Println(err.Error())
		return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}

	abciInfo, err := rpcClient.ABCIInfo()
	if err != nil {
		fmt.Println("Error when initializing rpcClient")
		fmt.Println(err.Error())
		return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}

	defer rpcClient.Stop()

	return abciInfo.Response, nil
}
