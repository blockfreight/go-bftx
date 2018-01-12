package saberservice

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"testing"

	"gopkg.in/yaml.v2"
)

// TestLoadtransaction is the unit test of loadtransaction
func TestLoadtransaction(t *testing.T) {
	var tx *BFTXTransaction
	_gopath := os.Getenv("GOPATH")
	_bftxpath := "/src/github.com/blockfreight/go-bftx"
	txpath := "/examples/bftx.json"
	b, err := ioutil.ReadFile(_gopath + _bftxpath + txpath)

	err = json.Unmarshal(b, &tx)
	if err != nil {
		t.Errorf("%s", b)
		t.Errorf("the file %s cannot be unmarshaled.\n %v", _gopath+_bftxpath+txpath, err)
	}
}

// Unit test of the loadconfiguration funtion
func TestLoadconfiguration(t *testing.T) {
	var tx *BFTXEncryptionConfig

	_gopath := os.Getenv("GOPATH")
	_bftxpath := "/src/github.com/blockfreight/go-bftx"
	// txpath := "/examples/bftx.json"
	txpath := "/examples/config.yaml"
	b, err := ioutil.ReadFile(txpath)
	err = yaml.UnmarshalStrict(b, &tx)
	if err != nil {
		t.Errorf("%s", b)
		t.Errorf("================\n")
		t.Errorf("yaml file %s cannot be unmarshaled.\n %v", _gopath+_bftxpath+txpath, err)
	}
}

func TestSaberinputcli(t *testing.T) {
	var st, st2 Saberinput
	_gopath := os.Getenv("GOPATH")
	_bftxpath := "/src/github.com/blockfreight/go-bftx"
	st.mode = "test"
	st.address = "localhost:22222"
	st.txpath = _gopath + _bftxpath + "/examples/bftx.json"
	st.txconfigpath = _gopath + _bftxpath + "/examples/config.yaml"
	st.KeyName = _gopath + _bftxpath + "/examples/carol_pri_key.json"

	st2.mode = "mm"
	st2.address = "add"
	st2.txpath = "path"
	st2.txconfigpath = "configpath"

	in, err := ioutil.TempFile("", "")
	check(err)
	defer in.Close()

	_, err = io.WriteString(in, "t\n")
	check(err)

	_, err = in.Seek(0, 0)
	r := Saberinputcli(in)
	if r != st {
		t.Errorf("result unmatch, %s,%s", st, r)
	}

	in, err = ioutil.TempFile("", "")
	check(err)

	_, err = io.WriteString(in, "mm\nadd\npath\nconfigpath\n")
	check(err)
	_, err = in.Seek(0, 0)
	r2 := Saberinputcli(in)
	if r2 != st2 {
		t.Errorf("result unmatch, %s,%s", st2, r2)
	}

}

func TestSaberEncoding(t *testing.T) {
	var st Saberinput

	_gopath := os.Getenv("GOPATH")
	_bftxpath := "/src/github.com/blockfreight/go-bftx"

	st.mode = "test"
	st.address = "localhost:22222"
	st.txpath = _gopath + _bftxpath + "/examples/bftx.json"
	st.txconfigpath = _gopath + _bftxpath + "/examples/config.yaml"

	encr, err := SaberEncoding(st)
	if err != nil {
		t.Errorf("error: %v", err)
		t.Errorf("Returned message: %v", encr)
	}
}

func TestSaberDecoding(t *testing.T) {
	var st Saberinput

	_gopath := os.Getenv("GOPATH")
	_bftxpath := "/src/github.com/blockfreight/go-bftx"

	st.mode = "test"
	st.address = "localhost:22222"
	st.txpath = _gopath + _bftxpath + "/examples/bftx.json"
	st.txconfigpath = _gopath + _bftxpath + "/examples/config.yaml"
	st.KeyName = _gopath + _bftxpath + "/examples/carol_pri_key.json"

	encr, err := SaberEncoding(st)
	if err != nil {
		t.Errorf("error: %v", err)
		t.Errorf("Returned message: %v", encr)
	}
}
