package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"net/http" // Provides HTTP client and server implementations.

	"github.com/blockfreight/go-bftx/lib/app/bf_tx"
	"github.com/blockfreight/go-bftx/lib/pkg/common"
	"github.com/blockfreight/go-bftx/lib/pkg/leveldb"
	"github.com/blockfreight/go-bftx/lib/pkg/saberservice"
	rpc "github.com/tendermint/tendermint/rpc/client"
	// Provides HTTP client and server implementations.
	// ===============
	// Tendermint Core
	// ===============
)

func ConstructBfTx(transaction bf_tx.BF_TX) (interface{}, error) {
	if err := transaction.GenerateBFTXUID(common.ORIGIN_API); err != nil {
		return nil, err
	}

	return transaction, nil
}

func SignBfTx(idBftx string) (interface{}, error) {
	var transaction bf_tx.BF_TX
	if err := transaction.SignBFTX(idBftx, common.ORIGIN_API); err != nil {
		return nil, err
	}

	return transaction, nil
}

func EncryptBfTx(idBftx string) (interface{}, error) {
	var transaction bf_tx.BF_TX
	data, err := leveldb.GetBfTx(idBftx)
	if err != nil {
		return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}

	json.Unmarshal(data, &transaction)

	if err != nil {
		if err.Error() == "LevelDB Get function: BF_TX not found." {
			return nil, errors.New(strconv.Itoa(http.StatusNotFound))
		}
		return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}

	if transaction.Verified {
		return nil, errors.New(strconv.Itoa(http.StatusNotAcceptable))
	}

	nwbftx, err := saberservice.BftxStructConverstionON(&transaction)
	if err != nil {
		log.Fatalf("Conversion error, can not convert old bftx to new bftx structure")
		return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}
	st := saberservice.SaberDefaultInput()
	saberbftx, err := saberservice.SaberEncoding(nwbftx, st)
	if err != nil {
		return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}
	bftxold, err := saberservice.BftxStructConverstionNO(saberbftx)
	//update the encoded transaction to database
	// Get the BF_TX content in string format
	content, err := bf_tx.BFTXContent(*bftxold)
	if err != nil {
		return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}

	// Update on DB
	err = leveldb.RecordOnDB(string(bftxold.Id), content)
	if err != nil {
		return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}

	return bftxold, nil
}

func DecryptBfTx(idBftx string) (interface{}, error) {
	var transaction bf_tx.BF_TX
	data, err := leveldb.GetBfTx(idBftx)
	if err != nil {
		return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}

	json.Unmarshal(data, &transaction)

	if err != nil {
		if err.Error() == "LevelDB Get function: BF_TX not found." {
			return nil, errors.New(strconv.Itoa(http.StatusNotFound))
		}
		return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}

	if transaction.Verified {
		return nil, errors.New(strconv.Itoa(http.StatusNotAcceptable))
	}

	nwbftx, err := saberservice.BftxStructConverstionON(&transaction)
	if err != nil {
		log.Fatalf("Conversion error, can not convert old bftx to new bftx structure")
		return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}
	st := saberservice.SaberDefaultInput()
	saberbftx, err := saberservice.SaberDecoding(nwbftx, st)
	if err != nil {
		return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}
	bftxold, err := saberservice.BftxStructConverstionNO(saberbftx)
	//update the encoded transaction to database
	// Get the BF_TX content in string format
	content, err := bf_tx.BFTXContent(*bftxold)
	if err != nil {
		return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}

	// Update on DB
	err = leveldb.RecordOnDB(string(bftxold.Id), content)
	if err != nil {
		return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}

	return bftxold, nil
}

func BroadcastBfTx(idBftx string) (interface{}, error) {
	var transaction bf_tx.BF_TX

	if err := transaction.BroadcastBfTx(idBftx, common.ORIGIN_API); err != nil {
		return nil, err
	}

	return transaction, nil
}

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

func GetTotal() (interface{}, error) {
	// Query the total of BF_TX in DB
	total, err := leveldb.Total()
	if err != nil {
		return nil, err
	}

	return total, nil
}

func GetTransaction(idBftx string) (interface{}, error) {
	var transaction bf_tx.BF_TX
	data, err := leveldb.GetBfTx(idBftx)
	if err != nil {
		return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}

	json.Unmarshal(data, &transaction)
	if err != nil {
		if err.Error() == "LevelDB Get function: BF_TX not found." {
			return nil, errors.New(strconv.Itoa(http.StatusNotFound))
		}
		return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}

	/* TODO: DECRYPT TRANSACTION */

	return transaction, nil
}

func QueryTransaction(idBftx string) (interface{}, error) {
	rpcClient := rpc.NewHTTP(os.Getenv("LOCAL_RPC_CLIENT_ADDRESS"), "/websocket")
	err := rpcClient.Start()
	if err != nil {
		fmt.Println("Error when initializing rpcClient")
		log.Fatal(err.Error())
	}
	defer rpcClient.Stop()
	query := "bftx.id='" + idBftx + "'"
	resQuery, err := rpcClient.TxSearch(query, true)
	if err != nil {
		return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}

	if len(resQuery) > 0 {
		var transaction bf_tx.BF_TX
		err := json.Unmarshal(resQuery[0].Tx, &transaction)
		if err != nil {
			return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
		}

		return transaction, nil
	}

	return nil, errors.New(strconv.Itoa(http.StatusNotFound))
}
