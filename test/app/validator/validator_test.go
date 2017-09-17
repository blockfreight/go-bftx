package validator

import (
  "testing"

  "github.com/blockfreight/blockfreight-alpha/lib/app/bf_tx"
  "github.com/blockfreight/blockfreight-alpha/lib/app/validator"
)

func TestValidator(t *testing.T) {
  t.Log("Test on validator function")
  bftx, err := bf_tx.SetBF_TX("../../../examples/bf_tx_example.json")
  if err != nil {
    t.Log(err.Error())
  }
  result, err := validator.ValidateBf_Tx(bftx)
  if err != nil {
    t.Log(err.Error())
  }
  if result != "Success! [OK]" {
    t.Error("Error on result of TestValidator")
    t.Error(result)
  }
}
