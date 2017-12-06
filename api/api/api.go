package api

import (
	"net/http" // Provides HTTP client and server implementations.
	"github.com/gorilla/mux" //Implements a request router and dispatcher for matching incoming requests to their respective handler
	"github.com/blockfreight/go-bftx/api/transaction"
)


func Start() error {
	//configuration, _ := config.LoadConfiguration()
	router := mux.NewRouter()
	router.HandleFunc("/fulltransaction", transaction.FullTransactionBfTx).Methods("POST")
	router.HandleFunc("/transaction/construct", transaction.ConstructBfTx).Methods("POST")
	router.HandleFunc("/transaction/sign/{id}", transaction.SignBfTx).Methods("PUT")
	router.HandleFunc("/transaction/broadcast/{id}", transaction.BroadcastBfTx).Methods("PUT")
	router.HandleFunc("/transaction/{id}", transaction.GetTransaction).Methods("GET")
	return http.ListenAndServe(":12345", router)	
}


