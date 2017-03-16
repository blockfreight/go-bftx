package validator

import (
	"testing"

	"github.com/blockfreight/blockfreight-alpha/blockfreight/bft/bf_tx"
)

func TestValidator(t *testing.T) {
	t.Log("Test on validator function")
	bf_tx := bf_tx.SetBF_TX("../.././files/bf_tx_example.json")
    result := ValidateBf_Tx(bf_tx)
    if result != "Success! [OK]" {
    	t.Error("Error on result of TestValidator")
    	t.Error(result)
    }
}