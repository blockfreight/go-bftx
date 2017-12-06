package transaction

import (
	"fmt"
	"encoding/json"
	"net/http" // Provides HTTP client and server implementations.
	"github.com/gorilla/mux" //Implements a request router and dispatcher for matching incoming requests to their respective handler
	"github.com/blockfreight/go-bftx/lib/app/bf_tx"         // Defines the Blockfreight™ Transaction (BF_TX) transaction standard and provides some useful functions to work with the BF_TX.
	"github.com/blockfreight/go-bftx/lib/app/validator"     // Provides functions to assure the input JSON is correct.
	"github.com/blockfreight/go-bftx/lib/pkg/leveldb"       // Provides some useful functions to work with LevelDB.
	"github.com/blockfreight/go-bftx/lib/pkg/crypto"        // Provides useful functions to sign BF_TX.
	"github.com/blockfreight/lib/app/response"

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
		response.Error(w, err, http.StatusInternalServerError)
		return
	}

	hash, err := bf_tx.HashBFTX(transaction)
	if err != nil {
		response.Error(w, err, http.StatusInternalServerError)
		return
	}

	// Generate BF_TX id
	transaction.Id = fmt.Sprintf("%x", bf_tx.GenerateBFTXSalt(hash, resInfo.LastBlockAppHash))

	// Re-validate a BF_TX before create a BF_TX
	_ , err = validator.ValidateBFTX(transaction)
	if err != nil {
		response.Error(w, err, http.StatusBadRequest)
		return
	}	

	// Sign BF_TX
	transaction, err = crypto.SignBFTX(transaction)
	if err != nil {
		response.Error(w, err, http.StatusInternalServerError)
		return
	}

	transaction.Transmitted = true

	// Get the BF_TX content in string format
	content, err := bf_tx.BFTXContent(transaction)
	if err != nil {
		response.Error(w, err, http.StatusInternalServerError)
		return
	}

	// Save on DB
	err = leveldb.RecordOnDB(transaction.Id, content)
	if err != nil {
		response.Error(w, err, http.StatusInternalServerError)
		return
	}

	// Deliver / Publish a BF_TX
	TendermintClient.DeliverTxSync([]byte(content))
	
	// Check the BF_TX hash
	TendermintClient.CommitSync()
	
	response.Success(w, transaction)
	
}

func ConstructBfTx(w http.ResponseWriter, r *http.Request) {
	var transaction bf_tx.BF_TX
	_ = json.NewDecoder(r.Body).Decode(&transaction)

	resInfo, err := TendermintClient.InfoSync()
	if err != nil {
		response.Error(w, err, http.StatusInternalServerError)
		return
	}

	hash, err := bf_tx.HashBFTX(transaction)
	if err != nil {
		response.Error(w, err, http.StatusInternalServerError)
		return
	}

	// Generate BF_TX id
	transaction.Id = fmt.Sprintf("%x", bf_tx.GenerateBFTXSalt(hash, resInfo.LastBlockAppHash))

	// Re-validate a BF_TX before create a BF_TX
	_ , err = validator.ValidateBFTX(transaction)
	if err != nil {
		response.Error(w, err, http.StatusBadRequest)
		return
	}	

	// Get the BF_TX content in string format
	content, err := bf_tx.BFTXContent(transaction)
	if err != nil {
		response.Error(w, err, http.StatusInternalServerError)
		return
	}

	// Save on DB
	err = leveldb.RecordOnDB(transaction.Id, content)
	if err != nil {
		response.Error(w, err, http.StatusInternalServerError)
		return
	}

	response.Success(w, transaction)
}

func SignBfTx(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println(params["id"])
	// Get a BF_TX by id
	transaction, err := leveldb.GetBfTx(params["id"])
	if err != nil {
		response.Error(w, err, http.StatusInternalServerError)
		return
	}
	if transaction.Verified {
		w.WriteHeader(http.StatusConflict)
		if err := json.NewEncoder(w).Encode(transaction); err != nil {
			panic(err)
		}
		return
	}

	// Sign BF_TX
	transaction, err = crypto.SignBFTX(transaction)
	if err != nil {
		response.Error(w, err, http.StatusInternalServerError)
		return
	}

	// Get the BF_TX content in string format
	content, err := bf_tx.BFTXContent(transaction)
	if err != nil {
		response.Error(w, err, http.StatusInternalServerError)
		return
	}

	// Update on DB
	err = leveldb.RecordOnDB(string(transaction.Id), content)
	if err != nil {
		response.Error(w, err, http.StatusInternalServerError)
		return
	}

	response.Success(w, transaction)
}

func BroadcastBfTx(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println(params["id"])
	// Get a BF_TX by id
	transaction, err := leveldb.GetBfTx(params["id"])
	if err != nil {
		response.Error(w, err, http.StatusInternalServerError)
		return
	}

	if !transaction.Verified {
		response.Error(w, err, http.StatusConflict)
		return
	}
	if transaction.Transmitted {
		response.Error(w, err, http.StatusConflict)
		return
	}

	// Change the boolean valud for Transmitted attribute
	transaction.Transmitted = true

	// Get the BF_TX content in string format
	content, err := bf_tx.BFTXContent(transaction)
	if err != nil {
		response.Error(w, err, http.StatusInternalServerError)
		return
	}

	// Update on DB
	err = leveldb.RecordOnDB(string(transaction.Id), content)
	if err != nil {
		response.Error(w, err, http.StatusInternalServerError)
		return
	}

	// Deliver / Publish a BF_TX
	TendermintClient.DeliverTxSync([]byte(content))

	// Check the BF_TX hash
	TendermintClient.CommitSync()

	response.Success(w, transaction)
}

func GetTransaction(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	// Get a BF_TX by id
	transaction, err := leveldb.GetBfTx(params["id"])
	if err != nil {
		if err.Error() == "LevelDB Get function: BF_TX not found." {
			response.Error(w, err, http.StatusNotFound)
			return
		}
		response.Error(w, err, http.StatusInternalServerError)
		return
	}

	response.Success(w, transaction)
}

