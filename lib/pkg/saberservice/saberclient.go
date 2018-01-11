package saberservice

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

type saberinput struct {
	mode         string
	address      string
	txpath       string
	txconfigpath string
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
func SaberEncoding(content string) ([]byte, error) {
	r := []byte{}
	return r, nil
}
