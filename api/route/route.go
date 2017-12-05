package route

import (
	"net/http" // Provides HTTP client and server implementations.
	"github.com/gorilla/mux" //Implements a request router and dispatcher for matching incoming requests to their respective handler
	"encoding/json"
	//"github.com/blockfreight/go-bftx/config" //Package that handles with the application configutarions.
	"github.com/blockfreight/go-bftx/lib/app/bf_tx"         // Defines the Blockfreight™ Transaction (BF_TX) transaction standard and provides some useful functions to work with the BF_TX.
)

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
	
	var bftx bf_tx.BF_TX
	_ = json.NewDecoder(r.Body).Decode(&bftx)
	if err := json.NewEncoder(w).Encode(bftx); err != nil {
        panic(err)
    }
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