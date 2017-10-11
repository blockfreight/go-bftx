package crypto

import (
	"testing"

	"github.com/blockfreight/go-bftx/lib/app/bf_tx"
	"github.com/blockfreight/go-bftx/lib/pkg/crypto"
)

func TestSignBFTX(t *testing.T) {
	t.Log("Test on SignBFTX function")
	bftx, err := bf_tx.SetBFTX("../../../examples/bf_tx_example.json")
	if err != nil {
		t.Log(err.Error())
	}

	bftx, err = crypto.SignBFTX(bftx)
	if err != nil {
		t.Log(err.Error())
	}

	if bftx.Signature == "" {
		t.Error("Error on bf_tx.Signature")
	}
	if bftx.Verified == false {
		t.Error("Error on bf_tx.Verified")
	}
}
