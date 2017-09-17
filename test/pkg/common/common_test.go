package common

import (
  "testing"
  "bytes"

	"github.com/blockfreight/blockfreight-alpha/lib/pkg/common"
)

func TestHashByteArrays(t* testing.T){
  first_hash := []byte("firstHash")
  second_hash := []byte("secondHash")
  result_expected := []byte{35, 63, 72, 164, 129, 98, 5, 123, 77, 35, 41, 21, 136, 230, 199, 208, 195, 68, 188, 65, 198, 199, 175, 43, 113, 168, 46, 95, 93, 208, 85, 227}

  resultCommon := common.HashByteArrays(first_hash, second_hash)

  if  bytes.Compare(resultCommon, result_expected) != 0 {
    t.Error("Error on HashByteArrays!")
  }
}
