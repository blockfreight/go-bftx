package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/blockfreight/go-bftx/lib/app/bf_tx"
	"github.com/blockfreight/go-bftx/lib/app/response"
	"github.com/blockfreight/go-bftx/lib/app/validator"
	"github.com/blockfreight/go-bftx/lib/pkg/crypto"
	"github.com/blockfreight/go-bftx/lib/pkg/leveldb"
	"github.com/gorilla/mux"

	// ===============
	// Tendermint Core
	// ===============
	"github.com/tendermint/abci/client"
)

var TendermintClient abcicli.Client

// Construct the Blockfreight™ Transaction [BF_TX]
func FullTransactionBfTx(w http.ResponseWriter, r *http.Request) {
	var transaction bf_tx.BF_TX
	_ = json.NewDecoder(r.Body).Decode(&transaction)

	resInfo, err := TendermintClient.InfoSync()
	if err != nil {
		response.Error(w, http.StatusInternalServerError)
		return
	}

	hash, err := bf_tx.HashBFTX(transaction)
	if err != nil {
		response.Error(w, http.StatusInternalServerError)
		return
	}

	// Generate BF_TX id
	transaction.Id = bf_tx.GenerateBFTXUID(hash, resInfo.LastBlockAppHash)

	// Re-validate a BF_TX before create a BF_TX
	_, err = validator.ValidateBFTX(transaction)
	if err != nil {
		response.Error(w, http.StatusBadRequest)
		return
	}

	// Sign BF_TX
	transaction, err = crypto.SignBFTX(transaction)
	if err != nil {
		response.Error(w, http.StatusInternalServerError)
		return
	}

	transaction.Transmitted = true

	// Get the BF_TX content in string format
	content, err := bf_tx.BFTXContent(transaction)
	if err != nil {
		response.Error(w, http.StatusInternalServerError)
		return
	}

	// Save on DB
	if err = leveldb.RecordOnDB(string(transaction.Id), content); err != nil {
		response.Error(w, http.StatusInternalServerError)
		return
	}

	// Deliver / Publish a BF_TX
	TendermintClient.DeliverTxSync([]byte(content))

	// Check the BF_TX hash
	TendermintClient.CommitSync()

	response.SuccessTx(w, transaction)

}

func ConstructBfTx(transaction bf_tx.BF_TX) (bf_tx.BF_TX, error) {
	resInfo, err := TendermintClient.InfoSync()
	if err != nil {
		return bf_tx.BF_TX{}, err
	}

	hash, err := bf_tx.HashBFTX(transaction)
	if err != nil {
		return bf_tx.BF_TX{}, err
	}

	// Generate BF_TX id
	transaction.Id = bf_tx.GenerateBFTXUID(hash, resInfo.LastBlockAppHash)

	// Re-validate a BF_TX before create a BF_TX
	/*_, err = validator.ValidateBFTX(transaction)
	if err != nil {
		return bf_tx.BF_TX{}, err
	}*/

	// Get the BF_TX content in string format
	content, err := bf_tx.BFTXContent(transaction)
	if err != nil {
		return bf_tx.BF_TX{}, err
	}

	// Save on DB
	if err = leveldb.RecordOnDB(transaction.Id, content); err != nil {
		return bf_tx.BF_TX{}, err
	}

	return transaction, nil
}

func SignBfTx(idBftx string) (interface{}, error) {
	// Get a BF_TX by id
	fmt.Println(idBftx)
	transaction, err := leveldb.GetBfTx(idBftx)
	if err != nil {
		return nil, err
	}
	if transaction.Verified {
		return nil, err
	}

	fmt.Printf("%+v\n", transaction)

	// Sign BF_TX
	transaction, err = crypto.SignBFTX(transaction)
	if err != nil {
		return nil, err
	}

	// Get the BF_TX content in string format
	content, err := bf_tx.BFTXContent(transaction)
	if err != nil {
		return nil, err
	}

	// Update on DB
	if err = leveldb.RecordOnDB(string(transaction.Id), content); err != nil {
		return nil, err
	}

	return transaction, nil
}

func BroadcastBfTx(idBftx string) (interface{}, error) {

	// Get a BF_TX by id
	transaction, err := leveldb.GetBfTx(idBftx)
	if err != nil {
		return nil, err
	}

	if !transaction.Verified {
		return nil, errors.New("Transaction not verified.")
	}
	if transaction.Transmitted {
		return nil, errors.New("Transaction already transmitted.")
	}

	// Change the boolean valud for Transmitted attribute
	transaction.Transmitted = true

	// Get the BF_TX content in string format
	content, err := bf_tx.BFTXContent(transaction)
	if err != nil {
		return nil, err
	}

	// Update on DB
	if err = leveldb.RecordOnDB(string(transaction.Id), content); err != nil {
		return nil, err
	}

	// Deliver / Publish a BF_TX
	TendermintClient.DeliverTxSync([]byte(content))

	// Check the BF_TX hash
	TendermintClient.CommitSync()

	return transaction, nil
}

func GetTransaction(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if params["id"] == "" {
		total, _ := leveldb.Total()
		response.SuccessInt(w, total)
		return
	}

	// Get a BF_TX by id
	transaction, err := leveldb.GetBfTx(params["id"])
	if err != nil {
		if err.Error() == "LevelDB Get function: BF_TX not found." {
			response.Error(w, http.StatusNotFound)
			return
		}
		response.Error(w, http.StatusInternalServerError)
		return
	}

	response.SuccessTx(w, transaction)
}
