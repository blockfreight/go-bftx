package response

import (
	"net/http" // Provides HTTP client and server implementations.
	"encoding/json"
	"fmt"
	"github.com/blockfreight/go-bftx/lib/app/bf_tx"         // Defines the Blockfreight™ Transaction (BF_TX) transaction standard and provides some useful functions to work with the BF_TX.
)

type Response struct {
	TotalTransactions int
	Transaction bf_tx.BF_TX
}

func Error(w http.ResponseWriter, responseStatusCode int) {
	w.WriteHeader(responseStatusCode)
}

func SuccessTx(w http.ResponseWriter, transaction bf_tx.BF_TX) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	response := Response {
		Transaction: transaction,
	}
	fmt.Printf("%+v\n", transaction)
	if err := json.NewEncoder(w).Encode(response); err != nil {
        panic(err)
	}
}

func SuccessInt(w http.ResponseWriter, total int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	response := Response {
		TotalTransactions: total,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
        panic(err)
	}
}
