package route

import (
	"net/http" // Provides HTTP client and server implementations.
	"fmt"
	"log"
	"github.com/gorilla/mux" //Implements a request router and dispatcher for matching incoming requests to their respective handler
	"github.com/blockfreight/go-bftx/config" //Package that handles with the application configutarions.
	"github.com/blockfreight/go-bftx/lib/app/bf_tx"         // Defines the Blockfreight™ Transaction (BF_TX) transaction standard and provides some useful functions to work with the BF_TX.
)

func StartApi() error {
	configuration, _ := config.LoadConfiguration()
	router := mux.NewRouter()
	router.HandleFunc("/transaction/{tx}", apiConstructBfTx).Methods("POST")
	router.HandleFunc("/transaction", apiConstructBfTx).Methods("GET")
	log.Fatal(http.ListenAndServe(configuration.BFTX_API_ADDRESS, router))

	return nil
}


// Construct the Blockfreight™ Transaction [BF_TX]
func apiConstructBfTx(w http.ResponseWriter, r *http.Request) {
	var bftx bf_tx.BF_TX
	test := bftx.Transmitted
	params := mux.Vars(r)
	fmt.Println(params)
	fmt.Println(test)
/*
	// Read JSON and instance the BF_TX structure
	bftx, err := bf_tx.SetBFTX(tx)
	if err != nil {
		return err
	}

	newId, err := cmdGenerateBftxID(bftx)
	if err != nil {
		return err
	}

	// Re-validate a BF_TX before create a BF_TX
	result, err := validator.ValidateBFTX(bftx)
	if err != nil {
		fmt.Println(result)
		return err
	}

	// Get the BF_TX content in string format
	content, err := bf_tx.BFTXContent(bftx)
	if err != nil {
		return err
	}

	// Save on DB
	err = leveldb.RecordOnDB(bftx.Id, content)
	if err != nil {
		return err
	}

	// Result
	printResponse(c, response{
		Result: "BF_TX Id: " + bftx.Id,
	})*/
}