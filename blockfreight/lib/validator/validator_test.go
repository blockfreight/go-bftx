package validator

import (
	"fmt"
	"testing"

	"github.com/blockfreight/blockfreight-alpha/blockfreight/lib/bf_tx"
)

func TestValidator(t *testing.T) {
	t.Log("Test on validator function")
	bf_tx, err := bf_tx.SetBF_TX("../.././files/bf_tx_example.json")
	if err != nil {
        t.Log(err.Error())
    }
    result, err := ValidateBf_Tx(bf_tx)
    if err != nil {
        fmt.Println(result)
        t.Log(err.Error())
    }
    if result != "Success! [OK]" {
    	t.Error("Error on result of TestValidator")
    	t.Error(result)
    }
}