package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/tendermint/abci/client"

	"net/http" // Provides HTTP client and server implementations.

	"github.com/blockfreight/go-bftx/lib/app/bf_tx"
	"github.com/blockfreight/go-bftx/lib/pkg/crypto"
	"github.com/blockfreight/go-bftx/lib/pkg/leveldb"
	"github.com/blockfreight/go-bftx/lib/pkg/saberservice"
	rpc "github.com/tendermint/tendermint/rpc/client"
	tmTypes "github.com/tendermint/tendermint/types"

	// Provides HTTP client and server implementations.
	// ===============
	// Tendermint Core
	// ===============
	abciTypes "github.com/tendermint/abci/types"
)

var TendermintClient abcicli.Client

// ConstructBfTx function to create a BFTX via API
func ConstructBfTx(transaction bf_tx.BF_TX) (interface{}, error) {

	resInfo, err := TendermintClient.InfoSync(abciTypes.RequestInfo{})
	if err != nil {
		return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}

	hash, err := bf_tx.HashBFTX(transaction)
	if err != nil {
		return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}

	// Generate BF_TX id
	transaction.Id = bf_tx.GenerateBFTXUID(hash, resInfo.LastBlockAppHash)

	/*jsonContent, err := json.Marshal(transaction)
	if err != nil {
		return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}

	transaction.Private = string(crypto.CryptoTransaction(string(jsonContent)))*/

	// Get the BF_TX content in string format
	content, err := bf_tx.BFTXContent(transaction)
	if err != nil {
		return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}

	/* TODO: ENCRYPT TRANSACTION */

	// Save on DB
	if err = leveldb.RecordOnDB(transaction.Id, content); err != nil {
		return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}

	return transaction, nil
}

// SignBfTx function to sign a BFTX via API
func SignBfTx(idBftx string) (interface{}, error) {
	transaction, err := leveldb.GetBfTx(idBftx)

	if err != nil {
		if err.Error() == "LevelDB Get function: BF_TX not found." {
			return nil, errors.New(strconv.Itoa(http.StatusNotFound))
		}
		return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}

	if transaction.Verified {
		return nil, errors.New(strconv.Itoa(http.StatusNotAcceptable))
	}

	// Sign BF_TX
	transaction, err = crypto.SignBFTX(transaction)
	if err != nil {
		return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}

	/*jsonContent, err := json.Marshal(transaction)
	if err != nil {
		return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}

	transaction.Private = string(crypto.CryptoTransaction(string(jsonContent)))*/

	// Get the BF_TX content in string format
	content, err := bf_tx.BFTXContent(transaction)
	if err != nil {
		return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}

	// Update on DB
	if err = leveldb.RecordOnDB(string(transaction.Id), content); err != nil {
		return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}

	return transaction, nil
}

// EncryptBfTx function to encrypt a BFTX via API
func EncryptBfTx(idBftx string) (interface{}, error) {
	transaction, err := leveldb.GetBfTx(idBftx)

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

// DecryptBfTx function to decrypt a BFTX via API
func DecryptBfTx(idBftx string) (interface{}, error) {
	transaction, err := leveldb.GetBfTx(idBftx)

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

// BroadcastBfTx function to broadcast a BFTX via API
func BroadcastBfTx(idBftx string) (interface{}, error) {
	rpcClient := rpc.NewHTTP(os.Getenv("LOCAL_RPC_CLIENT_ADDRESS"), "/websocket")
	err := rpcClient.Start()
	if err != nil {
		fmt.Println("Error when initializing rpcClient")
		log.Fatal(err.Error())
	}

	// Get a BF_TX by id
	transaction, err := leveldb.GetBfTx(idBftx)
	if err != nil {
		if err.Error() == "LevelDB Get function: BF_TX not found." {
			return nil, errors.New(strconv.Itoa(http.StatusNotFound))
		}
		return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}

	if !transaction.Verified {
		return nil, errors.New(strconv.Itoa(http.StatusNotAcceptable))
	}
	if transaction.Transmitted {
		return nil, errors.New(strconv.Itoa(http.StatusNotAcceptable))
	}

	// Change the boolean valud for Transmitted attribute
	transaction.Transmitted = true

	// Get the BF_TX content in string format
	content, err := bf_tx.BFTXContent(transaction)
	if err != nil {
		return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}

	// Update on DB
	if err = leveldb.RecordOnDB(string(transaction.Id), content); err != nil {
		return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}

	var tx tmTypes.Tx
	tx = []byte(content)

	_, rpcErr := rpcClient.BroadcastTxSync(tx)
	if rpcErr != nil {
		fmt.Printf("%+v\n", rpcErr)
		return nil, rpcErr
	}

	defer rpcClient.Stop()

	return transaction, nil
}

// GetInfo function to get info about the network via API
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

// GetTotal function to get the total amount of transactions via API
func GetTotal() (interface{}, error) {
	// Query the total of BF_TX in DB
	total, err := leveldb.Total()
	if err != nil {
		return nil, err
	}

	return total, nil
}

// GetTransaction function to get a transaction, from the local database, by id via API
func GetTransaction(idBftx string) (interface{}, error) {
	transaction, err := leveldb.GetBfTx(idBftx)
	if err != nil {
		if err.Error() == "LevelDB Get function: BF_TX not found." {
			return nil, errors.New(strconv.Itoa(http.StatusNotFound))
		}
		return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}

	/* TODO: DECRYPT TRANSACTION */

	return transaction, nil
}

// QueryTransaction function to query a transaction, from the network, by id via API
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
		return nil, err
	}

	if len(resQuery) > 0 {
		var transaction bf_tx.BF_TX
		err := json.Unmarshal(resQuery[0].Tx, &transaction)
		if err != nil {
			return nil, err
		}

		return transaction, nil
	}

	return nil, errors.New(strconv.Itoa(http.StatusNotFound))
}
