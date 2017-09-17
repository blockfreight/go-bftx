package common

import (
  "testing"
  "bytes"

	"github.com/blockfreight/blockfreight-alpha/lib/app/bf_tx"
	"github.com/blockfreight/blockfreight-alpha/lib/pkg/common"
)

func TestHashByteArrays(t* testing.T){
  first_hash := []byte("firstHash")
  second_hash := []byte("secondHash")

  resultBftx := bf_tx.HashBF_TX_salt(first_hash, second_hash)
  resultCommon := common.HashByteArrays(first_hash, second_hash)

  if  bytes.Compare(resultCommon, resultBftx) != 0 {
    t.Error("Error on HashByteArrays!")
  }
}
