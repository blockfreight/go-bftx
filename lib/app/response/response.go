package response

import (
	"net/http" // Provides HTTP client and server implementations.
	"encoding/json"
	"github.com/blockfreight/go-bftx/lib/app/bf_tx"         // Defines the Blockfreight™ Transaction (BF_TX) transaction standard and provides some useful functions to work with the BF_TX.
)

func Error(w http.ResponseWriter, err error, responseStatusCode int) {
	w.WriteHeader(responseStatusCode)
    if err := json.NewEncoder(w).Encode(err); err != nil {
        panic(err)
    }
}

func Success(w http.ResponseWriter, transaction bf_tx.BF_TX) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(transaction); err != nil {
        panic(err)
	}
}
