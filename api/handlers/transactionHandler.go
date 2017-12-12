package handlers

import (
	"encoding/json"
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
	transaction.Id = fmt.Sprintf("%x", bf_tx.GenerateBFTXSalt(hash, resInfo.LastBlockAppHash))

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
	transaction.Id = fmt.Sprintf("%x", bf_tx.GenerateBFTXSalt(hash, resInfo.LastBlockAppHash))

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

func SignBfTx(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println(params["id"])
	// Get a BF_TX by id
	transaction, err := leveldb.GetBfTx(params["id"])
	if err != nil {
		response.Error(w, http.StatusInternalServerError)
		return
	}
	if transaction.Verified {
		w.WriteHeader(http.StatusConflict)
		if err := json.NewEncoder(w).Encode(transaction); err != nil {
			panic(err)
		}

	}

	// Sign BF_TX
	transaction, err = crypto.SignBFTX(transaction)
	if err != nil {
		response.Error(w, http.StatusInternalServerError)
		return
	}

	// Get the BF_TX content in string format
	content, err := bf_tx.BFTXContent(transaction)
	if err != nil {
		response.Error(w, http.StatusInternalServerError)
		return
	}

	// Update on DB
	if err = leveldb.RecordOnDB(string(transaction.Id), content); err != nil {
		response.Error(w, http.StatusInternalServerError)
		return
	}

	response.SuccessTx(w, transaction)
}

func BroadcastBfTx(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println(params["id"])
	// Get a BF_TX by id
	transaction, err := leveldb.GetBfTx(params["id"])
	if err != nil {
		response.Error(w, http.StatusInternalServerError)
		return
	}

	if !transaction.Verified {
		response.Error(w, http.StatusConflict)
		return
	}
	if transaction.Transmitted {
		response.Error(w, http.StatusConflict)
		return
	}

	// Change the boolean valud for Transmitted attribute
	transaction.Transmitted = true

	// Get the BF_TX content in string format
	content, err := bf_tx.BFTXContent(transaction)
	if err != nil {
		response.Error(w, http.StatusInternalServerError)
		return
	}

	// Update on DB
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
