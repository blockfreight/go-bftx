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

	btx "github.com/blockfreight/go-bftx/lib/app/bf_tx"
	"google.golang.org/grpc"
	yaml "gopkg.in/yaml.v2"
)

// Saberinput structure for go-bftx use and test
type Saberinput struct {
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

// NVCsvConverterNew is a function that
// convert an array of bftx parameters to BFTXTransaction structure.
// This is used for the converting the Lading.csv to bftx.BFTX
//--------------------------------------------------------------------------------
func NVCsvConverterNew(line []string) *BFTXTransaction {
	msg := BFTXTransaction{
		Properties: &BFTX_Payload{
			Shipper:         line[0],
			Consignee:       line[1],
			ReceiveAgent:    line[2],
			HouseBill:       line[3],
			PortOfLoading:   line[4],
			PortOfDischarge: line[5],
			Destination:     line[6],
			MarksAndNumbers: line[7],
			DescOfGoods:     nvparsedesc(line[8]),
			GrossWeight:     line[9],
			UnitOfWeight:    line[10],
			Volume:          line[11],
			UnitOfVolume:    line[12],
			Container:       line[13],
			ContainerSeal:   line[14],
			ContainerMode:   line[15],
			ContainerType:   line[16],
			Packages:        line[17],
			PackType:        line[18],
			INCOTerms:       line[19],
			DeliverAgent:    line[20],
		},
	}
	return &msg
}

// NVCsvConverterOld is a function that
// convert an array of bftx parameters to BF_TX structure.
// This is used for the converting the Lading.csv to bftx.BFTX
//--------------------------------------------------------------------------------
func NVCsvConverterOld(line []string) btx.BF_TX {
	msg := btx.BF_TX{
		Properties: btx.Properties{
			Shipper:         line[0],
			Consignee:       line[1],
			ReceiveAgent:    line[2],
			HouseBill:       line[3],
			PortOfLoading:   line[4],
			PortOfDischarge: line[5],
			Destination:     line[6],
			MarksAndNumbers: line[7],
			DescOfGoods:     nvparsedesc(line[8]),
			GrossWeight:     line[9],
			UnitOfWeight:    line[10],
			Volume:          line[11],
			UnitOfVolume:    line[12],
			Container:       line[13],
			ContainerSeal:   line[14],
			ContainerMode:   line[15],
			ContainerType:   line[16],
			Packages:        line[17],
			PackType:        line[18],
			INCOTerms:       line[19],
			DeliverAgent:    line[20],
		},
	}
	return msg
}

// NVCsvConverter helper functions
// nvparseasfloat provides error handling necessary for bf_tx.Properties single-value float context
// func nvparseasfloat(num string) float32 {
// 	c, err := strconv.ParseFloat(num, 32)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return float32(c)
// }

// nvparseasint provides error handling necessary for bf_tx.Properties single-value int context
// func nvparseasint(num string) int64 {
// 	c, err := strconv.Atoi(num)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return int64(c)
// }

// nvparseasint provides error handling necessary for bf_tx.Properties single-value string context
func nvparsedesc(desc string) string {
	item := desc
	item = strings.Replace(item, "\n", " ", -1)
	item = strings.Replace(item, "\t", " ", -1)
	item = strings.Replace(item, "\r", " ", -1)
	return item
}

//--------------------------------------------------------------------------------

// BftxStructConverstionNO (New to Old) is a function that convert the new *BFTXTransaction
// structure to old struct *BF_TX. These two structure is duplicated somehow. This function is used
// for temporal conversion.
// since this is just for temporal usage, I will just use json marshal and unmarshal
// to convert structures
func BftxStructConverstionNO(tx *BFTXTransaction) (*btx.BF_TX, error) {
	var oldbftx btx.BF_TX
	bfjs, err := json.Marshal(*tx)
	if err != nil {
		log.Fatal("\nBftxStructConverstionNO convertion error\n", err)
	}
	err = json.Unmarshal(bfjs, &oldbftx)
	if err != nil {
		log.Fatal("BftxStructConverstionNO Converstion failed. Maybe because of different structure. ", err)
	}
	return &oldbftx, err
}

// BftxStructConverstionON (Old to New) is a function that convert the old structure
// *BF_TX to new structure *BFTXTransaction. These two structure is duplicated somehow. This function is used
// for temporal conversion.
// since this is just for temporal usage, I will just use json marshal and unmarshal
// to convert structures
func BftxStructConverstionON(tx *btx.BF_TX) (*BFTXTransaction, error) {
	var newbftx BFTXTransaction
	bfjs, err := json.Marshal(tx)
	if err != nil {
		log.Fatal("\nBftxStructConverstionON convertion error\n", err)
	}
	err = json.Unmarshal(bfjs, &newbftx)
	if err != nil {
		log.Fatal("BftxStructConverstionON Converstion failed. Maybe because of different structure. ", err)
	}
	return &newbftx, err
}

// SaberDefaultInput provides the saberinput structure with default value
func SaberDefaultInput() Saberinput {
	var st Saberinput
	_gopath := os.Getenv("GOPATH")
	_bftxpath := "/src/github.com/blockfreight/go-bftx"

	st.address = "localhost:22222"
	st.txconfigpath = _gopath + _bftxpath + "/examples/config.yaml"
	st.KeyName = "./Data/carol_pri_key.json"

	return st
}

// Saberinputcli provides interactive interaction for user to use bftx interface
func Saberinputcli(in *os.File) (st Saberinput) {
	fmt.Println("Please type your mode: 't' for test mode, 'm' for massconstruct, otherwise type your settings")
	if in == nil {
		in = os.Stdin
	}
	reader := bufio.NewReader(in)
	txt, _ := reader.ReadString('\n')
	txt = strings.Replace(txt, "\n", "", -1)
	_gopath := os.Getenv("GOPATH")
	_bftxpath := "/src/github.com/blockfreight/go-bftx"
	if txt == "t" {
		st.mode = "test"
		st.address = "localhost:22222"
		st.txpath = _gopath + _bftxpath + "/examples/bftx.json"
		st.txconfigpath = _gopath + _bftxpath + "/examples/config.yaml"
		// For server run on localhost
		st.KeyName = _gopath + _bftxpath + "/examples/carol_pri_key.json"
		// For server run on docker
		st.KeyName = "./Data/carol_pri_key.json"
	} else if txt == "m" {
		st.mode = "massconstruct"
		st.address = "localhost:22222"
		st.txconfigpath = _gopath + _bftxpath + "/examples/config.yaml"
		st.KeyName = "./Data/carol_pri_key.json"
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
		fmt.Println("Please type your decryption key path:")
		txt, _ = reader.ReadString('\n')
		st.KeyName = strings.Replace(txt, "\n", "", -1)
	}
	return st
}

// SaberEncoding takes an Bftx transaction and returns the saber encoded messages
func SaberEncoding(tx *BFTXTransaction, st Saberinput) (*BFTXTransaction, error) {
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

// SaberEncodingTestCase is the function that enable it to connect to a container which realizing the
// saber encoding service
func SaberEncodingTestCase(st Saberinput) (*BFTXTransaction, error) {
	switch st.mode {
	//	case "massconstruct":
	//		err := massSaberEncoding(st)
	//		return nil, err
	default:
		tx := loadtransaction(st.txpath)
		// txconfig := loadconfiguration(st.txconfigpath)

		// bfencreq := BFTX_EncodeRequest{
		// 	Bftxtrans:  tx,
		// 	Bftxconfig: txconfig,
		// }

		// conn, err := grpc.Dial(st.address, grpc.WithInsecure())
		// if err != nil {
		// 	log.Fatalf("%s cannot connected by program: %v", st.address, err)
		// }
		// defer conn.Close()
		// c := NewBFSaberServiceClient(conn)

		// encr, err := c.BFTX_Encode(context.Background(), &bfencreq)
		// check(err)
		encr, err := SaberEncoding(tx, st)
		return encr, err
	}
}

// massSaberEncoding is used for massively load the transaction from the lading.csv file
/*
func massSaberEncoding(st Saberinput) error {
	// define the index i
	i := 0
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	// fmt.Printf(wd + "/examples/Lading.csv")
	csvFile, err := os.Open(wd + "/examples/Lading.csv")
	if err != nil {
		log.Fatal("csv read error:\n", err)
	}

	// Define the abci client
	abciClient, err := abcicli.NewClient("tcp://127.0.0.1:46658", "socket", true)
	if err != nil {
		fmt.Println("Error when starting abci client")
		log.Fatalf("\n massSaberEncoding Error: %+v \n", err)
	}
	err = abciClient.Start()
	if err != nil {
		fmt.Println("Error when initializing abciClient")
		log.Fatal(err.Error())
	}
	defer abciClient.Stop()

	// Define the rpc client
	rpcClient := rpc.NewHTTP("tcp://127.0.0.1:46657", "/websocket")

	err = rpcClient.Start()
	if err != nil {
		fmt.Println("Error when initializing rpcClient")
		log.Fatal(err.Error())
	}
	defer rpcClient.Stop()

	// define the grpc client for saber service
	conn, err := grpc.Dial(st.address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("%s cannot connected by program: %v", st.address, err)
	}
	defer conn.Close()
	bfsaberclient := NewBFSaberServiceClient(conn)

	txconfig := loadconfiguration(st.txconfigpath)

	reader := csv.NewReader(bufio.NewReader(csvFile))

	for {
		i++
		line, err := reader.Read()
		if err == io.EOF {
			log.Fatal("io read error", err)
		}
		if err != nil {
			log.Fatal(err)
		}

		if len(line) != 22 {
			fmt.Printf("breaking line number: %d \n", i)
			fmt.Printf("Line has wrong length:%d \n", len(line))
			fmt.Printf("Line: %+v", line)
			continue
		}
		tx := NVCsvConverterNew(line)

		bfencreq := BFTX_EncodeRequest{
			Bftxtrans:  tx,
			Bftxconfig: txconfig,
		}
		bfencr, err := bfsaberclient.BFTX_Encode(context.Background(), &bfencreq)
		if err != nil {
			log.Printf("Line %d, BFTX_Encode error: %v", i, err)
			continue
		}

		// do the bftx sign--------------------------------------
		oldbf, err := BftxStructConverstionNO(bfencr)
		bfmsg := *oldbf
		if err != nil {
			log.Fatalf("Cannot do the conversion: %v", err)
		}

		salt, err := th.GetBlockAppHash(abciClient)
		if err != nil {
			log.Fatalf("GetBlockAppHash error: %v", err)
		}

		// Hash BF_TX Object
		hash, err := btx.HashBFTX(bfmsg)
		if err != nil {
			log.Fatalf("HashBFTX error: %v", err)
		}
		// Generate BF_TX id
		bftxID := btx.HashByteArray(hash, salt)
		bfmsg.Id = bftxID
		// do the bftx sign--------------------------------------

		if err = crypto.SignBFTX()
		if err != nil {
			return err
		}

		bfmsg, err = crypto.SignBFTX(bfmsg)
		if err != nil {
			return err
		}
		// Change the boolean valud for Transmitted attribute
		bfmsg.Transmitted = true

		// Get the BF_TX content in string format
		content, err := btx.BFTXContent(bfmsg)
		if err != nil {
			log.Fatal("BFTXContent error", err)
			return err
		}
		// Update on database
		err = leveldb.RecordOnDB(string(bfmsg.Id), content)
		if err != nil {
			log.Fatal("BFTXContent error", err)
			return err
		}

		resp, err := rpcClient.BroadcastTxSync([]byte(content))
		if err != nil {
			log.Fatal("rpcclient err:", err)
		}
		// added for flow control

		if i%100 == 0 {
			fmt.Print(i, "\n")
			fmt.Printf(bfmsg.Id, "\n")
			fmt.Printf(": %+v\n", resp)
		} else {
			fmt.Print(i, ",")
		}

	}
	return err

}
*/
// SaberDecodingTestCase is the function that enable it to connect to a container which realizing the
// saber decoding service
func SaberDecoding(tx *BFTXTransaction, st Saberinput) (*BFTXTransaction, error) {

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
