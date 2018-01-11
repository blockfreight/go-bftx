package saberservice

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"testing"

	"gopkg.in/yaml.v2"
)

// TestLoadtransaction is the unit test of loadtransaction
func TestLoadtransaction(t *testing.T) {
	var tx *BFTXTransaction

	txpath := "../../../examples/bftx.json"
	b, err := ioutil.ReadFile(txpath)

	err = json.Unmarshal(b, &tx)
	if err != nil {
		t.Errorf("%s", b)
		t.Errorf("the file %s cannot be unmarshaled.\n %v", txpath, err)
	}
}

// Unit test of the loadconfiguration funtion
func TestLoadconfiguration(t *testing.T) {
	var tx *BFTXEncryptionConfig

	txpath := "../../../examples/config.yaml"
	b, err := ioutil.ReadFile(txpath)
	err = yaml.UnmarshalStrict(b, &tx)
	if err != nil {
		t.Errorf("%s", b)
		t.Errorf("================\n")
		t.Errorf("yaml file %s cannot be unmarshaled.\n %v", txpath, err)
	}
}

func TestSaberinputcli(t *testing.T) {
	var st, st2 saberinput
	st.mode = "test"
	st.address = "localhost:22222"
	st.txpath = "../../../examples/bftx.json"
	st.txconfigpath = "../../../examples/config.yaml"

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
	r := saberinputcli(in)
	if r != st {
		t.Errorf("result unmatch, %s,%s", st, r)
	}

	in, err = ioutil.TempFile("", "")
	check(err)

	_, err = io.WriteString(in, "mm\nadd\npath\nconfigpath\n")
	check(err)
	_, err = in.Seek(0, 0)
	r2 := saberinputcli(in)
	if r2 != st2 {
		t.Errorf("result unmatch, %s,%s", st2, r2)
	}

}

// func TestSaberEncoding(t *testing.T) {
// 	const address = "localhost:22222"

// 	tx := loadtransaction("../../../examples/bftx.json")
// 	txconfig := loadconfiguration("../../../examples/config.yaml")

// 	bfencreq := BFTX_EncodeRequest{
// 		Bftxtrans:  tx,
// 		Bftxconfig: txconfig,
// 	}

// 	conn, err := grpc.Dial(address, grpc.WithInsecure())
// 	if err != nil {
// 		log.Fatalf("%s cannot connected by testing program: %v", address, err)
// 	}
// 	defer conn.Close()
// 	c := NewBFSaberServiceClient(conn)

// 	encr, err := c.BFTX_Encode(context.Background(), &bfencreq)
// 	check(err)
// 	_, err = fmt.Print(encr)
// 	check(err)
// }
