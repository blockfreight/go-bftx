package handlers

import (
	"errors"
	"strconv"

	"net/http" // Provides HTTP client and server implementations.

	"github.com/blockfreight/go-bftx/lib/app/bf_tx"
	"github.com/blockfreight/go-bftx/lib/pkg/crypto"
	"github.com/blockfreight/go-bftx/lib/pkg/leveldb"

	// Provides HTTP client and server implementations.
	// ===============
	// Tendermint Core
	// ===============
	"github.com/tendermint/abci/client"
)

var TendermintClient abcicli.Client

func ConstructBfTx(transaction bf_tx.BF_TX) (interface{}, error) {
	resInfo, err := TendermintClient.InfoSync()
	if err != nil {
		return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}

	hash, err := bf_tx.HashBFTX(transaction)
	if err != nil {
		return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}

	// Generate BF_TX id
	transaction.Id = bf_tx.GenerateBFTXUID(hash, resInfo.LastBlockAppHash)

	// Get the BF_TX content in string format
	content, err := bf_tx.BFTXContent(transaction)
	if err != nil {
		return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}

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

	// Deliver / Publish a BF_TX
	TendermintClient.DeliverTxSync([]byte(content))

	// Check the BF_TX hash
	TendermintClient.CommitSync()

	return transaction, nil
}

func GetTransaction(idBftx string) (interface{}, error) {
	transaction, err := leveldb.GetBfTx(idBftx)
	if err != nil {
		if err.Error() == "LevelDB Get function: BF_TX not found." {
			return nil, errors.New(strconv.Itoa(http.StatusNotFound))
		}
		return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}

	return transaction, nil
}
