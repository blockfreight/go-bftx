package bf_tx

import (
	"reflect"
	"testing"

	bftx "github.com/blockfreight/go-bftx/lib/app/bf_tx"
)

func TestSetBF_TX(t *testing.T) {
	t.Log("Test on SetBF_TX function")
	var prot bftx.BF_TX
	bf_tx, err := bftx.SetBF_TX("../../../examples/bf_tx_example.json")
	if err != nil {
		t.Log(err.Error())
	}
	if reflect.TypeOf(bf_tx) != reflect.TypeOf(prot) {
		t.Error("Error on type of result of SetBF_TX")
	}
}

func TestTransmitedState(t *testing.T) {
	t.Log("Test on State function")
	bf_tx, err := bftx.SetBF_TX("../../../examples/bf_tx_example.json")
	if err != nil {
		t.Log(err.Error())
	}

	bf_tx.Transmitted = true
	result := bftx.State(bf_tx)

	if result != "Transmitted!" {
		t.Error("Error on string result of bftx.State() when Transmitted = true")
	}

}

func TestSignedState(t *testing.T) {
	t.Log("Test on State function")
	bf_tx, err := bftx.SetBF_TX("../../../examples/bf_tx_example.json")
	if err != nil {
		t.Log(err.Error())
	}

	bf_tx.Verified = true
	result := bftx.State(bf_tx)

	if result != "Signed!" {
		t.Error("Error on string result of bftx.State() when Verified = true")
	}

}

func TestConstructedState(t *testing.T) {
	t.Log("Test on State function")
	bf_tx, err := bftx.SetBF_TX("../../../examples/bf_tx_example.json")
	if err != nil {
		t.Log(err.Error())
	}

	result := bftx.State(bf_tx)

	if result != "Constructed!" {
		t.Error("Error on string result of bftx.State() when Transaction is Constructed")
	}

}

func TestReinitialize(t *testing.T) {
	t.Log("Test on Reinitialize function")
	var prop bftx.BF_TX

	bf_tx := bftx.Reinitialize(prop)

	if bf_tx.PrivateKey.Curve != nil || bf_tx.PrivateKey.X != nil || bf_tx.PrivateKey.D != nil || bf_tx.Signhash != nil || bf_tx.Signature != "" || bf_tx.Verified != false || bf_tx.Transmitted != false {
		t.Error("Error on BF_TX object returned by function bf_tx.Reinitialize()")
	}
}
