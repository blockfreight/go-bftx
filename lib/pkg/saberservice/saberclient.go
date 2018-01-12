package saberservice

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"google.golang.org/grpc"
	yaml "gopkg.in/yaml.v2"
)

type saberinput struct {
	mode         string
	address      string
	txpath       string
	txconfigpath string
	KeyName      string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func loadtransaction(s string) *BFTXTransaction {
	var bftx *BFTXTransaction

	jsmsg, err := ioutil.ReadFile(s)
	check(err)

	err = json.Unmarshal(jsmsg, &bftx)
	check(err)
	return bftx
}

func loadconfiguration(s string) *BFTXEncryptionConfig {
	var bfconfig *BFTXEncryptionConfig

	ylconfig, err := ioutil.ReadFile(s)
	check(err)

	err = yaml.UnmarshalStrict(ylconfig, &bfconfig)
	check(err)
	return bfconfig
}

func saberinputcli(in *os.File) (st saberinput) {
	fmt.Println("Please type your mode: 't' for test mode, otherwise type your mode")
	if in == nil {
		in = os.Stdin
	}
	reader := bufio.NewReader(in)
	txt, _ := reader.ReadString('\n')
	txt = strings.Replace(txt, "\n", "", -1)
	if txt == "t" {
		st.mode = "test"
		st.address = "localhost:22222"
		st.txpath = "../../../examples/bftx.json"
		st.txconfigpath = "../../../examples/config.yaml"
	} else {
		st.mode = txt
		fmt.Println("Please type your service host address:")
		txt, _ := reader.ReadString('\n')
		st.address = strings.Replace(txt, "\n", "", -1)
		fmt.Println("Please type your transaction file path:")
		txt, _ = reader.ReadString('\n')
		st.txpath = strings.Replace(txt, "\n", "", -1)
		fmt.Println("Please type your transaction configuration file path:")
		txt, _ = reader.ReadString('\n')
		st.txconfigpath = strings.Replace(txt, "\n", "", -1)
	}
	return st
}

// SaberEncoding is the function that enable it to connect to a container which realizing the
// saber encoding service
func SaberEncoding(st saberinput) (*BFTXTransaction, error) {

	tx := loadtransaction(st.txpath)
	txconfig := loadconfiguration(st.txconfigpath)

	bfencreq := BFTX_EncodeRequest{
		Bftxtrans:  tx,
		Bftxconfig: txconfig,
	}

	conn, err := grpc.Dial(st.address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("%s cannot connected by program: %v", st.address, err)
	}
	defer conn.Close()
	c := NewBFSaberServiceClient(conn)

	encr, err := c.BFTX_Encode(context.Background(), &bfencreq)
	check(err)

	return encr, err
}

// SaberDecoding is the function that enable it to connect to a container which realizing the
// saber decoding service
func SaberDecoding(st saberinput, tx *BFTXTransaction) (*BFTXTransaction, error) {

	bfdcpreq := BFTX_DecodeRequest{}

	conn, err := grpc.Dial(st.address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("%s cannot connected by program: %v", st.address, err)
	}
	defer conn.Close()
	c := NewBFSaberServiceClient(conn)

	bfdcpreq.Bftxtrans = tx
	bfdcpreq.KeyName = st.KeyName

	_, err = fmt.Print("\n==============================\n")
	check(err)

	dcpr, err := c.BFTX_Decode(context.Background(), &bfdcpreq)
	check(err)
	fmt.Print(dcpr)

	return dcpr, err
}
