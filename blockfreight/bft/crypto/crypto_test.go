package crypto

import(
	"testing"

	"github.com/blockfreight/blockfreight-alpha/blockfreight/bft/bf_tx"
)

func TestSign_BF_TX(t *testing.T) {
	t.Log("Test on Sign_BF_TX function")
	bftx := bf_tx.SetBF_TX("../.././files/bf_tx_example.json")
    
	bftx = Sign_BF_TX(bftx)

    if bftx.Signature == "" {
    	t.Error("Error on bf_tx.Signature")
    }
    if bftx.Verified == false {
    	t.Error("Error on bf_tx.Verified")
    }
}