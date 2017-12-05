package route

import (
	"net/http" // Provides HTTP client and server implementations.
	"github.com/gorilla/mux" //Implements a request router and dispatcher for matching incoming requests to their respective handler
	"fmt"
	"encoding/json"
	//"github.com/blockfreight/go-bftx/config" //Package that handles with the application configutarions.
	"github.com/blockfreight/go-bftx/lib/app/bf_tx"         // Defines the Blockfreight™ Transaction (BF_TX) transaction standard and provides some useful functions to work with the BF_TX.
	
	// ===============
	// Tendermint Core
	// ===============
	"github.com/tendermint/abci/client"
)

// client is a global variable so it can be reused by the console
var TendermintClient abcicli.Client

func StartApi() error {
	//configuration, _ := config.LoadConfiguration()
	router := mux.NewRouter()
	router.HandleFunc("/transaction", apiConstructBfTx).Methods("POST")
	router.HandleFunc("/transaction", apiConstructBfTx).Methods("GET")
	return http.ListenAndServe(":12345", router)	
}


// Construct the Blockfreight™ Transaction [BF_TX]
func apiConstructBfTx(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	
	var transaction bf_tx.BF_TX
	_ = json.NewDecoder(r.Body).Decode(&transaction)

	resInfo, err := TendermintClient.InfoSync()
	if err != nil {
		errorResponse(transaction, w, err)
		return
	}

	hash, err := bf_tx.HashBFTX(transaction)
	if err != nil {
		errorResponse(transaction, w, err)
		return
	}

	// Generate BF_TX id
	transaction.Id = fmt.Sprintf("%x", bf_tx.GenerateBFTXSalt(hash, resInfo.LastBlockAppHash))
	
	if err := json.NewEncoder(w).Encode(transaction); err != nil {
        panic(err)
	}
	
}

func errorResponse(transaction bf_tx.BF_TX, w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
    if err := json.NewEncoder(w).Encode(transaction); err != nil {
        panic(err)
    }
}