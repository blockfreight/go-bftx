package bf_tx

import (
	"reflect"
	"testing"

	bftx "github.com/blockfreight/go-bftx/lib/app/bf_tx"
)

func TestSetBFTX(t *testing.T) {
	t.Log("Test on SetBFTX function")
	var prot bftx.BF_TX
	bftx, err := bftx.SetBFTX("../../../examples/bf_tx_example.json")
	if err != nil {
		t.Log(err.Error())
	}
	if reflect.TypeOf(bftx) != reflect.TypeOf(prot) {
		t.Error("Error on type of result of SetBFTX")
	}
}

func TestTransmitedState(t *testing.T) {
	t.Log("Test on State function")
	newBftx, err := bftx.SetBFTX("../../../examples/bf_tx_example.json")
	if err != nil {
		t.Log(err.Error())
	}

	newBftx.Transmitted = true
	result := bftx.State(newBftx)

	if result != "Transmitted!" {
		t.Error("Error on string result of bftx.State() when Transmitted = true")
	}

}

func TestSignedState(t *testing.T) {
	t.Log("Test on State function")
	newBftx, err := bftx.SetBFTX("../../../examples/bf_tx_example.json")
	if err != nil {
		t.Log(err.Error())
	}

	newBftx.Verified = true
	result := bftx.State(newBftx)

	if result != "Signed!" {
		t.Error("Error on string result of bftx.State() when Verified = true")
	}

}

func TestConstructedState(t *testing.T) {
	t.Log("Test on State function")
	newBftx, err := bftx.SetBFTX("../../../examples/bf_tx_example.json")
	if err != nil {
		t.Log(err.Error())
	}

	result := bftx.State(newBftx)

	if result != "Constructed!" {
		t.Error("Error on string result of bftx.State() when Transaction is Constructed")
	}

}

func TestReinitialize(t *testing.T) {
	t.Log("Test on Reinitialize function")
	var prop bftx.BF_TX

	newBftx := bftx.Reinitialize(prop)

	if newBftx.PrivateKey.Curve != nil || newBftx.PrivateKey.X != nil || newBftx.PrivateKey.D != nil || newBftx.Signhash != nil || newBftx.Signature != "" || newBftx.Verified != false || newBftx.Transmitted != false {
		t.Error("Error on BF_TX object returned by function bf_tx.Reinitialize()")
	}
}
