package crypto

import(
	"testing"

	"github.com/blockfreight/blockfreight-alpha/lib/app/bf_tx"
)

func TestSign_BF_TX(t *testing.T) {
	t.Log("Test on Sign_BF_TX function")
	bftx, err := bf_tx.SetBF_TX("../../../examples/bf_tx_example.json")
    if err != nil {
        t.Log(err.Error())
    }
    
	bftx, err = Sign_BF_TX(bftx)
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