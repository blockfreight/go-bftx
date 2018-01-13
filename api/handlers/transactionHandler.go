package handlers

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"
	"strconv"

	"github.com/tendermint/abci/client"

	"github.com/blockfreight/go-bftx/lib/app/types"

	"net/http" // Provides HTTP client and server implementations.

	"github.com/blockfreight/go-bftx/lib/app/bf_tx"
	"github.com/blockfreight/go-bftx/lib/pkg/crypto"
	"github.com/blockfreight/go-bftx/lib/pkg/leveldb"

	// Provides HTTP client and server implementations.
	// ===============
	// Tendermint Core
	// ===============
	tendermintTypes "github.com/tendermint/abci/types"
)

var TendermintClient abcicli.Client

func ConstructBfTx(transaction bf_tx.BF_TX) (interface{}, error) {

	resInfo, err := TendermintClient.InfoSync(tendermintTypes.RequestInfo{})
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

func BroadcastBfTx(idBftx string) (interface{}, error) {

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

	/* TODO: ENCRYPT TRANSACTION */

	// Change the boolean valud for Transmitted attribute
	transaction.Transmitted = true

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

	// Deliver / Publish a BF_TX
	src := []byte(content)
	encodedStr := hex.EncodeToString(src)
	url := "http://localhost:46657/broadcast_tx_sync?tx=%22" + encodedStr + "%22"

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var broadcastResp types.ResponseBroadcast
	err = json.Unmarshal(body, &broadcastResp)

	return transaction, nil
}

func GetLocalTransaction(idBftx string) (interface{}, error) {
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

func GetBlockchainTransaction(idBftx string) {

}
