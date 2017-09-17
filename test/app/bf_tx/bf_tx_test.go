package bf_tx

import (
	"reflect"
	"testing"
  "github.com/blockfreight/blockfreight-alpha/lib/app/bf_tx"
)

func TestSetBF_TX(t *testing.T) {
	t.Log("Test on SetBF_TX function")
	var prot bf_tx.BF_TX
	bf_tx, err := bf_tx.SetBF_TX("../../../examples/bf_tx_example.json")
	if err != nil {
		t.Log(err.Error())
	}
	if reflect.TypeOf(bf_tx) != reflect.TypeOf(prot) {
		t.Error("Error on type of result of SetBF_TX")
	}
}
