package common

import (
	"bytes"
	"testing"

	"github.com/blockfreight/go-bftx/lib/pkg/common"
)

func TestHashByteArrays(t *testing.T) {
	t.Log("Test HashBytesArrays on Common Lib")
	firstHash := []byte("firstHash")
	secondHash := []byte("secondHash")
	resultExpected := []byte{35, 63, 72, 164, 129, 98, 5, 123, 77, 35, 41, 21, 136, 230, 199, 208, 195, 68, 188, 65, 198, 199, 175, 43, 113, 168, 46, 95, 93, 208, 85, 227}

	resultGenerated := common.HashByteArrays(firstHash, secondHash)

	if bytes.Compare(resultGenerated, resultExpected) != 0 {
		t.Error("Error on HashByteArrays!")
	}
}
