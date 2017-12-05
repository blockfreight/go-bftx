package api

import (
	"net/http" // Provides HTTP client and server implementations.
	"github.com/gorilla/mux" //Implements a request router and dispatcher for matching incoming requests to their respective handler
	"fmt"
	"encoding/json"
	//"github.com/blockfreight/go-bftx/config" //Package that handles with the application configutarions.
	"github.com/blockfreight/go-bftx/lib/app/bf_tx"         // Defines the Blockfreight™ Transaction (BF_TX) transaction standard and provides some useful functions to work with the BF_TX.
	"github.com/blockfreight/go-bftx/lib/app/validator"     // Provides functions to assure the input JSON is correct.
	"github.com/blockfreight/go-bftx/lib/pkg/leveldb"       // Provides some useful functions to work with LevelDB.
	"github.com/blockfreight/go-bftx/lib/pkg/crypto"        // Provides useful functions to sign BF_TX.

	// ===============
	// Tendermint Core
	// ===============
	"github.com/tendermint/abci/client"
)

// client is a global variable so it can be reused by the console
var TendermintClient abcicli.Client

func Start() error {
	//configuration, _ := config.LoadConfiguration()
	router := mux.NewRouter()
	router.HandleFunc("/fulltransaction", apiFullTransactionBfTx).Methods("POST")
	router.HandleFunc("/transaction/construct", apiConstructBfTx).Methods("POST")
	router.HandleFunc("/transaction/sign/{id}", apiSignBfTx).Methods("PUT")
	router.HandleFunc("/transaction/broadcast/{id}", apiBroadcastBfTx).Methods("PUT")
	router.HandleFunc("/transaction/{id}", apiGetTransaction).Methods("GET")
	return http.ListenAndServe(":12345", router)	
}


// Construct the Blockfreight™ Transaction [BF_TX]
func apiFullTransactionBfTx(w http.ResponseWriter, r *http.Request) {
	var transaction bf_tx.BF_TX
	_ = json.NewDecoder(r.Body).Decode(&transaction)

	resInfo, err := TendermintClient.InfoSync()
	if err != nil {
		errorResponse(w, err)
		return
	}

	hash, err := bf_tx.HashBFTX(transaction)
	if err != nil {
		errorResponse(w, err)
		return
	}

	// Generate BF_TX id
	transaction.Id = fmt.Sprintf("%x", bf_tx.GenerateBFTXSalt(hash, resInfo.LastBlockAppHash))

	// Re-validate a BF_TX before create a BF_TX
	_ , err = validator.ValidateBFTX(transaction)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}	

	// Sign BF_TX
	transaction, err = crypto.SignBFTX(transaction)
	if err != nil {
		errorResponse(w, err)
		return
	}

	transaction.Transmitted = true

	// Get the BF_TX content in string format
	content, err := bf_tx.BFTXContent(transaction)
	if err != nil {
		errorResponse(w, err)
		return
	}

	// Save on DB
	err = leveldb.RecordOnDB(transaction.Id, content)
	if err != nil {
		errorResponse(w, err)
		return
	}

	// Deliver / Publish a BF_TX
	TendermintClient.DeliverTxSync([]byte(content))
	
	// Check the BF_TX hash
	TendermintClient.CommitSync()
	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(transaction); err != nil {
        panic(err)
	}
	
}

func apiConstructBfTx(w http.ResponseWriter, r *http.Request) {
	var transaction bf_tx.BF_TX
	_ = json.NewDecoder(r.Body).Decode(&transaction)

	resInfo, err := TendermintClient.InfoSync()
	if err != nil {
		errorResponse(w, err)
		return
	}

	hash, err := bf_tx.HashBFTX(transaction)
	if err != nil {
		errorResponse(w, err)
		return
	}

	// Generate BF_TX id
	transaction.Id = fmt.Sprintf("%x", bf_tx.GenerateBFTXSalt(hash, resInfo.LastBlockAppHash))

	// Re-validate a BF_TX before create a BF_TX
	_ , err = validator.ValidateBFTX(transaction)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}	

	// Get the BF_TX content in string format
	content, err := bf_tx.BFTXContent(transaction)
	if err != nil {
		errorResponse(w, err)
		return
	}

	// Save on DB
	err = leveldb.RecordOnDB(transaction.Id, content)
	if err != nil {
		errorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(transaction); err != nil {
        panic(err)
	}
}

func apiSignBfTx(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println(params["id"])
	// Get a BF_TX by id
	transaction, err := leveldb.GetBfTx(params["id"])
	if err != nil {
		errorResponse(w, err)
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
		errorResponse(w, err)
		return
	}

	// Get the BF_TX content in string format
	content, err := bf_tx.BFTXContent(transaction)
	if err != nil {
		errorResponse(w, err)
		return
	}

	// Update on DB
	err = leveldb.RecordOnDB(string(transaction.Id), content)
	if err != nil {
		errorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(transaction); err != nil {
        panic(err)
	}
}

func apiBroadcastBfTx(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println(params["id"])
	// Get a BF_TX by id
	transaction, err := leveldb.GetBfTx(params["id"])
	if err != nil {
		errorResponse(w, err)
		return
	}

	if !transaction.Verified {
		w.WriteHeader(http.StatusConflict)
		if err := json.NewEncoder(w).Encode(transaction); err != nil {
			panic(err)
		}
		return
	}
	if transaction.Transmitted {
		w.WriteHeader(http.StatusConflict)
		if err := json.NewEncoder(w).Encode(transaction); err != nil {
			panic(err)
		}
		return
	}

	// Change the boolean valud for Transmitted attribute
	transaction.Transmitted = true

	// Get the BF_TX content in string format
	content, err := bf_tx.BFTXContent(transaction)
	if err != nil {
		errorResponse(w, err)
		return
	}

	// Update on DB
	err = leveldb.RecordOnDB(string(transaction.Id), content)
	if err != nil {
		errorResponse(w, err)
		return
	}

	// Deliver / Publish a BF_TX
	TendermintClient.DeliverTxSync([]byte(content))

	// Check the BF_TX hash
	TendermintClient.CommitSync()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(transaction); err != nil {
        panic(err)
	}	
}

func apiGetTransaction(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	// Get a BF_TX by id
	transaction, err := leveldb.GetBfTx(params["id"])
	if err != nil {
		if err.Error() == "LevelDB Get function: BF_TX not found." {
			w.WriteHeader(http.StatusNotFound)
			if err := json.NewEncoder(w).Encode(err); err != nil {
				panic(err)
			}
			return
		}
		errorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(transaction); err != nil {
        panic(err)
	}
}

func errorResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
    if err := json.NewEncoder(w).Encode(err); err != nil {
        panic(err)
    }
}