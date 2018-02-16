package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"time"

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

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

// simpleLogger writes errors and the function name that generated the error to bftx.log
func simpleLogger(i interface{}, currentError error) {
	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile(os.Getenv("GOPATH")+"/src/github.com/blockfreight/go-bftx/logs/api.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte(time.Now().Format("2006/01/02 15:04") + ", " + getFunctionName(i) + ", " + currentError.Error() + "\n\n")); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

// queryLogger writes errors, the function name that generated the error, and the transaction body to bftx.log for cmdQuery only
func queryLogger(i interface{}, currentError string, id string) {
	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile(os.Getenv("GOPATH")+"/src/github.com/blockfreight/go-bftx/logs/bftx.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte(time.Now().Format("2006/01/02 15:04") + ", " + getFunctionName(i) + ", " + currentError + ", " + id + "\n\n")); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

// transLogger writes errors, the function name that generated the error, and the transaction body to bftx.log
func transLogger(i interface{}, currentError error, id string) {
	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile(os.Getenv("GOPATH")+"/src/github.com/blockfreight/go-bftx/logs/api.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte(time.Now().Format("2006/01/02 15:04") + ", " + getFunctionName(i) + ", " + currentError.Error() + ", " + id + "\n\n")); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func ConstructBfTx(transaction bf_tx.BF_TX) (interface{}, error) {

	resInfo, err := TendermintClient.InfoSync(abciTypes.RequestInfo{})
	if err != nil {
		simpleLogger(ConstructBfTx, err)
		return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}

	hash, err := bf_tx.HashBFTX(transaction)
	if err != nil {
		simpleLogger(ConstructBfTx, err)
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
		transLogger(ConstructBfTx, err, transaction.Id)
		return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}

	/* TODO: ENCRYPT TRANSACTION */

	// Save on DB
	if err = leveldb.RecordOnDB(transaction.Id, content); err != nil {
		transLogger(ConstructBfTx, err, transaction.Id)
		return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}

	return transaction, nil
}

func SignBfTx(idBftx string) (interface{}, error) {
	transaction, err := leveldb.GetBfTx(idBftx)

	if err != nil {
		if err.Error() == "LevelDB Get function: BF_TX not found." {
			transLogger(SignBfTx, err, idBftx)
			return nil, errors.New(strconv.Itoa(http.StatusNotFound))
		}
		transLogger(SignBfTx, err, idBftx)
		return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}

	if transaction.Verified {
		return nil, errors.New(strconv.Itoa(http.StatusNotAcceptable))
	}

	// Sign BF_TX
	transaction, err = crypto.SignBFTX(transaction)
	if err != nil {
		transLogger(SignBfTx, err, idBftx)
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
		transLogger(SignBfTx, err, idBftx)
		return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}

	// Update on DB
	if err = leveldb.RecordOnDB(string(transaction.Id), content); err != nil {
		transLogger(SignBfTx, err, idBftx)
		return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}

	return transaction, nil
}

func EncryptBfTx(idBftx string) (interface{}, error) {
	transaction, err := leveldb.GetBfTx(idBftx)

	if err != nil {
		if err.Error() == "LevelDB Get function: BF_TX not found." {
			transLogger(EncryptBfTx, err, idBftx)
			return nil, errors.New(strconv.Itoa(http.StatusNotFound))
		}
		transLogger(EncryptBfTx, err, idBftx)
		return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}

	if transaction.Verified {
		return nil, errors.New(strconv.Itoa(http.StatusNotAcceptable))
	}

	nwbftx, err := saberservice.BftxStructConverstionON(&transaction)
	if err != nil {
		log.Fatalf("Conversion error, can not convert old bftx to new bftx structure")
		transLogger(EncryptBfTx, err, idBftx)
		return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}
	st := saberservice.SaberDefaultInput()
	saberbftx, err := saberservice.SaberEncoding(nwbftx, st)
	if err != nil {
		transLogger(EncryptBfTx, err, idBftx)
		return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}
	bftxold, err := saberservice.BftxStructConverstionNO(saberbftx)
	//update the encoded transaction to database
	// Get the BF_TX content in string format
	content, err := bf_tx.BFTXContent(*bftxold)
	if err != nil {
		transLogger(EncryptBfTx, err, idBftx)
		return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}

	// Update on DB
	err = leveldb.RecordOnDB(string(bftxold.Id), content)
	if err != nil {
		transLogger(EncryptBfTx, err, idBftx)
		return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}

	return bftxold, nil
}

func DecryptBfTx(idBftx string) (interface{}, error) {
	transaction, err := leveldb.GetBfTx(idBftx)

	if err != nil {
		if err.Error() == "LevelDB Get function: BF_TX not found." {
			transLogger(DecryptBfTx, err, idBftx)
			return nil, errors.New(strconv.Itoa(http.StatusNotFound))
		}
		transLogger(DecryptBfTx, err, idBftx)
		return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}

	if transaction.Verified {
		return nil, errors.New(strconv.Itoa(http.StatusNotAcceptable))
	}

	nwbftx, err := saberservice.BftxStructConverstionON(&transaction)
	if err != nil {
		log.Fatalf("Conversion error, can not convert old bftx to new bftx structure")
		transLogger(DecryptBfTx, err, idBftx)
		return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}
	st := saberservice.SaberDefaultInput()
	saberbftx, err := saberservice.SaberDecoding(nwbftx, st)
	if err != nil {
		transLogger(DecryptBfTx, err, idBftx)
		return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}
	bftxold, err := saberservice.BftxStructConverstionNO(saberbftx)
	//update the encoded transaction to database
	// Get the BF_TX content in string format
	content, err := bf_tx.BFTXContent(*bftxold)
	if err != nil {
		transLogger(DecryptBfTx, err, idBftx)
		return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}

	// Update on DB
	err = leveldb.RecordOnDB(string(bftxold.Id), content)
	if err != nil {
		transLogger(DecryptBfTx, err, idBftx)
		return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}

	return bftxold, nil
}

func BroadcastBfTx(idBftx string) (interface{}, error) {
	rpcClient := rpc.NewHTTP(os.Getenv("LOCAL_RPC_CLIENT_ADDRESS"), "/websocket")
	err := rpcClient.Start()
	if err != nil {
		fmt.Println("Error when initializing rpcClient")
		transLogger(BroadcastBfTx, err, idBftx)
		log.Fatal(err.Error())
	}

	// Get a BF_TX by id
	transaction, err := leveldb.GetBfTx(idBftx)
	if err != nil {
		if err.Error() == "LevelDB Get function: BF_TX not found." {
			transLogger(BroadcastBfTx, err, idBftx)
			return nil, errors.New(strconv.Itoa(http.StatusNotFound))
		}
		transLogger(BroadcastBfTx, err, idBftx)
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
		transLogger(BroadcastBfTx, err, idBftx)
		return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}

	// Update on DB
	if err = leveldb.RecordOnDB(string(transaction.Id), content); err != nil {
		transLogger(BroadcastBfTx, err, idBftx)
		return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}

	var tx tmTypes.Tx
	tx = []byte(content)

	_, rpcErr := rpcClient.BroadcastTxSync(tx)
	if rpcErr != nil {
		transLogger(BroadcastBfTx, rpcErr, idBftx)
		fmt.Printf("%+v\n", rpcErr)
		return nil, rpcErr
	}

	defer rpcClient.Stop()

	return transaction, nil
}

func GetInfo() (interface{}, error) {
	rpcClient := rpc.NewHTTP(os.Getenv("LOCAL_RPC_CLIENT_ADDRESS"), "/websocket")
	err := rpcClient.Start()
	if err != nil {
		fmt.Println("Error when initializing rpcClient")
		fmt.Println(err.Error())
		simpleLogger(GetInfo, err)
		return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}

	abciInfo, err := rpcClient.ABCIInfo()
	if err != nil {
		fmt.Println("Error when initializing rpcClient")
		fmt.Println(err.Error())
		simpleLogger(GetInfo, err)
		return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}

	defer rpcClient.Stop()

	return abciInfo.Response, nil
}

func GetTotal() (interface{}, error) {
	// Query the total of BF_TX in DB
	total, err := leveldb.Total()
	if err != nil {
		simpleLogger(GetTotal, err)
		return nil, err
	}

	return total, nil
}

func GetTransaction(idBftx string) (interface{}, error) {
	transaction, err := leveldb.GetBfTx(idBftx)
	if err != nil {
		if err.Error() == "LevelDB Get function: BF_TX not found." {
			transLogger(GetTransaction, err, idBftx)
			return nil, errors.New(strconv.Itoa(http.StatusNotFound))
		}
		transLogger(GetTransaction, err, idBftx)
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
		transLogger(QueryTransaction, err, idBftx)
	}
	defer rpcClient.Stop()
	query := "bftx.id='" + idBftx + "'"
	resQuery, err := rpcClient.TxSearch(query, true)
	if err != nil {
		transLogger(QueryTransaction, err, idBftx)
		return nil, err
	}

	if len(resQuery) > 0 {
		var transaction bf_tx.BF_TX
		err := json.Unmarshal(resQuery[0].Tx, &transaction)
		if err != nil {
			transLogger(QueryTransaction, err, idBftx)
			return nil, err
		}

		return transaction, nil
	}

	queryLogger(QueryTransaction, "Blockfreight Transaction not found.", idBftx)
	return nil, errors.New(strconv.Itoa(http.StatusNotFound))
}
