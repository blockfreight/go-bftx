// File: ./blockfreight/lib/bf_tx/bf_tx.go
// Summary: Application code for Blockfreight™ | The blockchain of global freight.
// License: MIT License
// Company: Blockfreight, Inc.
// Author: Julian Nunez, Neil Tran, Julian Smith, Gian Felipe & contributors
// Site: https://blockfreight.com
// Support: <support@blockfreight.com>

// Copyright © 2017 Blockfreight, Inc. All Rights Reserved.

// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the "Software"),
// to deal in the Software without restriction, including without limitation
// the rights to use, copy, modify, merge, publish, distribute, sublicense,
// and/or sell copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS
// OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// =================================================================================================================================================
// =================================================================================================================================================
//
// BBBBBBBBBBBb     lll                                kkk             ffff                         iii                  hhh            ttt
// BBBB``````BBBB   lll                                kkk            fff                           ```                  hhh            ttt
// BBBB      BBBB   lll      oooooo        ccccccc     kkk    kkkk  fffffff  rrr  rrr    eeeee      iii     gggggg ggg   hhh  hhhhh   tttttttt
// BBBBBBBBBBBB     lll    ooo    oooo    ccc    ccc   kkk   kkk    fffffff  rrrrrrrr eee    eeee   iii   gggg   ggggg   hhhh   hhhh  tttttttt
// BBBBBBBBBBBBBB   lll   ooo      ooo   ccc           kkkkkkk        fff    rrrr    eeeeeeeeeeeee  iii  gggg      ggg   hhh     hhh    ttt
// BBBB       BBB   lll   ooo      ooo   ccc           kkkk kkkk      fff    rrr     eeeeeeeeeeeee  iii   ggg      ggg   hhh     hhh    ttt
// BBBB      BBBB   lll   oooo    oooo   cccc    ccc   kkk   kkkk     fff    rrr      eee      eee  iii    ggg    gggg   hhh     hhh    tttt    ....
// BBBBBBBBBBBBB    lll     oooooooo       ccccccc     kkk     kkkk   fff    rrr       eeeeeeeee    iii     gggggg ggg   hhh     hhh     ttttt  ....
//                                                                                                        ggg      ggg
//   Blockfreight™ | The blockchain of global freight.                                                      ggggggggg
//
// =================================================================================================================================================
// =================================================================================================================================================

// Package bf_tx is a package that defines the Blockfreight™ Transaction (BF_TX) transaction standard
// and provides some useful functions to work with the BF_TX.
package bf_tx

import (
	// =======================
	// Golang Standard library
	// =======================
	"crypto/ecdsa" // Implements the Elliptic Curve Digital Signature Algorithm, as defined in FIPS 186-3.
	"crypto/elliptic"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"errors" // Implements the SHA256 Algorithm for Hash.
	"hash"
	"io"
	"log"
	"math/big"
	"os"
	// Implements encoding and decoding of JSON as defined in RFC 4627.
	"net/http"
	"strconv"

	"fmt" // Implements formatted I/O with functions analogous to C's printf and scanf.

	// ===============
	// Tendermint Core
	// ===============
	"github.com/tendermint/abci/client"
	abciTypes "github.com/tendermint/abci/types"
	rpc "github.com/tendermint/tendermint/rpc/client"
	tmTypes "github.com/tendermint/tendermint/types"

	// ====================
	// Third-party packages
	// ====================
	"github.com/davecgh/go-spew/spew" // Implements a deep pretty printer for Go data structures to aid in debugging.

	// ======================
	// Blockfreight™ packages
	// ======================

	"github.com/blockfreight/go-bftx/lib/app/bftx_logger"
	"github.com/blockfreight/go-bftx/lib/pkg/common"
	"github.com/blockfreight/go-bftx/lib/pkg/leveldb" // Implements common functions for Blockfreight™
)

var TendermintClient abcicli.Client

// SetBFTX receives the path of a JSON, reads it and returns the BF_TX structure with all attributes.
func SetBFTX(jsonpath string) (BF_TX, error) {
	var bftx BF_TX
	file, err := common.ReadJSON(jsonpath)
	if err != nil {
		bftx_logger.SimpleLogger("SetBFTX", err)
		return bftx, err
	}
	json.Unmarshal(file, &bftx)
	return bftx, nil
}

//HashBFTX hashes the BF_TX object
func HashBFTX(bftx BF_TX) ([]byte, error) {
	bftxBytes := []byte(fmt.Sprintf("%v", bftx))

	hash := sha256.New()
	hash.Write(bftxBytes)

	return hash.Sum(nil), nil
}

//HashByteArray hashes two byte arrays and returns it.
func HashByteArray(hash []byte, salt []byte) string {
	return "BFTX" + fmt.Sprintf("%x", common.HashByteArrays(hash, salt))
}

// BFTXContent receives the BF_TX structure, applies it the json.Marshal procedure and return the content of the BF_TX JSON.
func BFTXContent(bftx BF_TX) (string, error) {
	jsonContent, err := json.Marshal(bftx)
	return string(jsonContent), err
}

// PrintBFTX receives a BF_TX and prints it clearly.
func PrintBFTX(bftx BF_TX) {
	spew.Dump(bftx)
}

// State reports the current state of a BF_TX
func State(bftx BF_TX) string {
	if bftx.Transmitted {
		return "Transmitted!"
	} else if bftx.Verified {
		return "Signed!"
	} else {
		return "Constructed!"
	}
}

func ByteArrayToBFTX(obj []byte) BF_TX {
	var bftx BF_TX
	json.Unmarshal(obj, &bftx)
	return bftx
}

func (bftx *BF_TX) GenerateBFTX(origin string) error {
	resInfo, err := TendermintClient.InfoSync(abciTypes.RequestInfo{})
	if err != nil {
		bftx_logger.SimpleLogger("GenerateBFTX", err)
		return handleResponse(origin, err, strconv.Itoa(http.StatusInternalServerError))
	}

	hash, err := HashBFTX(*bftx)
	if err != nil {
		bftx_logger.SimpleLogger("GenerateBFTX", err)
		return handleResponse(origin, err, strconv.Itoa(http.StatusInternalServerError))
	}

	// Generate BF_TX id
	bftx.Id = HashByteArray(hash, resInfo.LastBlockAppHash)

	// Get the BF_TX content in string format
	content, err := BFTXContent(*bftx)
	if err != nil {
		bftx_logger.TransLogger("GenerateBFTX", err, bftx.Id)
		return handleResponse(origin, err, strconv.Itoa(http.StatusInternalServerError))
	}

	// Save on DB
	if err = leveldb.RecordOnDB(bftx.Id, content); err != nil {
		bftx_logger.TransLogger("GenerateBFTX", err, bftx.Id)
		return handleResponse(origin, err, strconv.Itoa(http.StatusInternalServerError))
	}

	return nil
}

func (bftx *BF_TX) SignBFTX(idBftx, origin string) error {
	data, err := leveldb.GetBfTx(idBftx)
	if err != nil {
		if err.Error() == "LevelDB Get function: BF_TX not found." {
			bftx_logger.TransLogger("SignBFTX", err, idBftx)
			return handleResponse(origin, err, strconv.Itoa(http.StatusNotFound))
		}
		bftx_logger.TransLogger("SignBFTX", err, idBftx)
		return handleResponse(origin, err, strconv.Itoa(http.StatusInternalServerError))
	}

	if err = json.Unmarshal(data, &bftx); err != nil {
		bftx_logger.TransLogger("SignBFTX", err, idBftx)
		return handleResponse(origin, err, strconv.Itoa(http.StatusInternalServerError))
	}

	if bftx.Verified {
		return handleResponse(origin, errors.New("Transaction already signed"), strconv.Itoa(http.StatusNotAcceptable))
	}

	// Sign BF_TX
	if err = bftx.setSignature(); err != nil {
		bftx_logger.TransLogger("SignBFTX", err, idBftx)
		return handleResponse(origin, err, strconv.Itoa(http.StatusInternalServerError))
	}

	// Get the BF_TX content in string format
	content, err := BFTXContent(*bftx)
	if err != nil {
		bftx_logger.TransLogger("SignBFTX", err, idBftx)
		return handleResponse(origin, err, strconv.Itoa(http.StatusInternalServerError))
	}

	// Update on DB
	if err = leveldb.RecordOnDB(bftx.Id, content); err != nil {
		bftx_logger.TransLogger("SignBFTX", err, idBftx)
		return handleResponse(origin, err, strconv.Itoa(http.StatusInternalServerError))
	}

	return nil

}

// SignBFTX has the whole process of signing each BF_TX.
func (bftx *BF_TX) setSignature() error {
	content, err := BFTXContent(*bftx)
	if err != nil {
		bftx_logger.TransLogger("setSignature", err, bftx.Id)
		return err
	}

	pubkeyCurve := elliptic.P256() //see http://golang.org/pkg/crypto/elliptic/#P256

	privatekey := new(ecdsa.PrivateKey)
	privatekey, err = ecdsa.GenerateKey(pubkeyCurve, rand.Reader) // this generates a public & private key pair
	if err != nil {
		bftx_logger.TransLogger("setSignature", err, bftx.Id)
		return err
	}
	pubkey := privatekey.PublicKey

	// Sign ecdsa style
	var h hash.Hash
	h = md5.New()
	r := big.NewInt(0)
	s := big.NewInt(0)

	io.WriteString(h, content)
	signhash := h.Sum(nil)

	r, s, err = ecdsa.Sign(rand.Reader, privatekey, signhash)
	if err != nil {
		bftx_logger.TransLogger("setSignature", err, bftx.Id)
		return err
	}

	signature := r.Bytes()
	signature = append(signature, s.Bytes()...)

	sign := ""
	for i, _ := range signature {
		sign += strconv.Itoa(int(signature[i]))
	}

	// Verification
	verifystatus := ecdsa.Verify(&pubkey, signhash, r, s)

	//Set Private Key and Sign to BF_TX
	bftx.PrivateKey = *privatekey
	bftx.Signhash = signhash
	bftx.Signature = sign
	bftx.Verified = verifystatus

	return nil
}

func (bftx *BF_TX) BroadcastBFTX(idBftx, origin string) error {
	rpcClient := rpc.NewHTTP(os.Getenv("LOCAL_RPC_CLIENT_ADDRESS"), "/websocket")
	err := rpcClient.Start()
	if err != nil {
		log.Println(err.Error())
		bftx_logger.TransLogger("BroadcastBFTX", err, idBftx)
		return handleResponse(origin, err, strconv.Itoa(http.StatusInternalServerError))
	}

	// Get a BF_TX by id
	data, err := leveldb.GetBfTx(idBftx)
	if err != nil {
		if err.Error() == "LevelDB Get function: BF_TX not found." {
			bftx_logger.TransLogger("BroadcastBFTX", err, idBftx)
			return handleResponse(origin, err, strconv.Itoa(http.StatusNotFound))
		}
		bftx_logger.TransLogger("BroadcastBFTX", err, idBftx)
		return handleResponse(origin, err, strconv.Itoa(http.StatusInternalServerError))
	}

	if err = json.Unmarshal(data, &bftx); err != nil {
		bftx_logger.TransLogger("BroadcastBFTX", err, idBftx)
		return handleResponse(origin, err, strconv.Itoa(http.StatusInternalServerError))
	}

	if !bftx.Verified {
		return handleResponse(origin, err, strconv.Itoa(http.StatusNotAcceptable))
	}
	if bftx.Transmitted {
		return handleResponse(origin, err, strconv.Itoa(http.StatusNotAcceptable))
	}

	// Change the boolean valud for Transmitted attribute
	bftx.Transmitted = true

	// Get the BF_TX content in string format
	content, err := BFTXContent(*bftx)
	if err != nil {
		bftx_logger.TransLogger("BroadcastBFTX", err, idBftx)
		return handleResponse(origin, err, strconv.Itoa(http.StatusInternalServerError))
	}

	// Update on DB
	if err = leveldb.RecordOnDB(string(bftx.Id), content); err != nil {
		bftx_logger.TransLogger("BroadcastBFTX", err, idBftx)
		return handleResponse(origin, err, strconv.Itoa(http.StatusInternalServerError))
	}

	var tx tmTypes.Tx
	tx = []byte(content)

	_, rpcErr := rpcClient.BroadcastTxSync(tx)
	if rpcErr != nil {
		fmt.Printf("%+v\n", rpcErr)
		bftx_logger.TransLogger("BroadcastBFTX", rpcErr, idBftx)
		return handleResponse(origin, err, strconv.Itoa(http.StatusInternalServerError))
	}

	defer rpcClient.Stop()

	return nil
}

func (bftx *BF_TX) GetBFTX(idBftx, origin string) error {
	data, err := leveldb.GetBfTx(idBftx)
	if err != nil {
		if err.Error() == "LevelDB Get function: BF_TX not found." {
			bftx_logger.TransLogger("GetBFTX", err, idBftx)
			return handleResponse(origin, err, strconv.Itoa(http.StatusNotFound))
		}
		bftx_logger.TransLogger("GetBFTX", err, idBftx)
		return handleResponse(origin, err, strconv.Itoa(http.StatusInternalServerError))
	}

	json.Unmarshal(data, &bftx)
	if err != nil {
		bftx_logger.TransLogger("GetBFTX", err, idBftx)
		return handleResponse(origin, err, strconv.Itoa(http.StatusInternalServerError))
	}

	return nil

}

func (bftx *BF_TX) QueryBFTX(idBftx, origin string) error {
	rpcClient := rpc.NewHTTP(os.Getenv("LOCAL_RPC_CLIENT_ADDRESS"), "/websocket")
	err := rpcClient.Start()
	if err != nil {
		log.Println(err.Error())
		// queryLogger("QueryBFTX", err.Error(), idBftx)
		bftx_logger.TransLogger("QueryBFTX", err, idBftx)
		return handleResponse(origin, err, strconv.Itoa(http.StatusInternalServerError))
	}
	defer rpcClient.Stop()
	query := "bftx.id='" + idBftx + "'"
	resQuery, err := rpcClient.TxSearch(query, true)
	if err != nil {
		bftx_logger.TransLogger("QueryBFTX", err, idBftx)
		return handleResponse(origin, err, strconv.Itoa(http.StatusInternalServerError))
	}

	if len(resQuery) > 0 {
		err := json.Unmarshal(resQuery[0].Tx, &bftx)
		if err != nil {
			bftx_logger.TransLogger("QueryBFTX", err, idBftx)
			return handleResponse(origin, err, strconv.Itoa(http.StatusInternalServerError))
		}
		return nil
	}

	bftx_logger.StringLogger("QueryBFTX", "Transaction not found", idBftx)
	return handleResponse(origin, errors.New("Transaction not found"), strconv.Itoa(http.StatusNotFound))
}

func (bftx BF_TX) GetTotal() (int, error) {
	total, err := leveldb.Total()
	if err != nil {
		bftx_logger.SimpleLogger("GetTotal", err)
		return 0, err
	}

	return total, nil
}

func (bftx *BF_TX) FullBFTXCycleWithoutEncryption(origin string) error {
	if err := bftx.GenerateBFTX(origin); err != nil {
		bftx_logger.SimpleLogger("FullBFTXCycleWithoutEncryption", err)
		return err
	}
	if err := bftx.SignBFTX(bftx.Id, origin); err != nil {
		bftx_logger.SimpleLogger("FullBFTXCycleWithoutEncryption", err)
		return err
	}
	if err := bftx.BroadcastBFTX(bftx.Id, origin); err != nil {
		bftx_logger.SimpleLogger("FullBFTXCycleWithoutEncryption", err)
		return err
	}

	return nil
}

func handleResponse(origin string, err error, httpStatusCode string) error {
	if origin == common.ORIGIN_API {
		return errors.New(httpStatusCode)
	}
	return err
}

// Reinitialize set the default values to the Blockfreight attributes of BF_TX
func Reinitialize(bftx BF_TX) BF_TX {
	bftx.PrivateKey.Curve = nil
	bftx.PrivateKey.X = nil
	bftx.PrivateKey.Y = nil
	bftx.PrivateKey.D = nil
	bftx.Signhash = nil
	bftx.Signature = ""
	bftx.Verified = false
	bftx.Transmitted = false
	return bftx
}

// BF_TX structure respresents an logical abstraction of a Blockfreight™ Transaction.
type BF_TX struct {
	// =========================
	// Bill of Lading attributes
	// =========================
	Properties Properties

	// ===================================
	// Blockfreight Transaction attributes
	// ===================================
	Id          string           `json:"Id"`
	PrivateKey  ecdsa.PrivateKey `json:"-"`
	Signhash    []uint8          `json:"Signhash"`
	Signature   string           `json:"Signature"`
	Verified    bool             `json:"Verified"`
	Transmitted bool             `json:"Transmitted"`
	Amendment   string           `json:"Amendment"`
	Private     string           `json:"Private"`
}

// Properties struct
type Properties struct {
	Shipper             string       `protobuf:"bytes,1,opt,name=Shipper" json:"Shipper"`
	BolNum              string       `protobuf:"varint,1,opt,name=BolNum" json:"BolNum"`
	RefNum              string       `protobuf:"varint,2,opt,name=RefNum" json:"RefNum"`
	Consignee           string       `protobuf:"bytes,2,opt,name=Consignee" json:"Consignee"`
	HouseBill           string       `protobuf:"bytes,3,opt,name=HouseBill" json:"HouseBill"`
	Vessel              string       `protobuf:"varint,3,opt,name=Vessel" json:"Vessel"`
	Packages            string       `protobuf:"varint,4,opt,name=Packages" json:"Packages"`
	PackType            string       `protobuf:"bytes,4,opt,name=PackType" json:"PackType"`
	INCOTerms           string       `protobuf:"bytes,5,opt,name=INCOTerms" json:"INCOTerms"`
	PortOfLoading       string       `protobuf:"bytes,6,opt,name=PortOfLoading" json:"PortOfLoading"`
	PortOfDischarge     string       `protobuf:"bytes,7,opt,name=PortOfDischarge" json:"PortOfDischarge"`
	Destination         string       `protobuf:"bytes,8,opt,name=Destination" json:"Destination"`
	MarksAndNumbers     string       `protobuf:"bytes,9,opt,name=MarksAndNumbers" json:"MarksAndNumbers"`
	UnitOfWeight        string       `protobuf:"bytes,10,opt,name=UnitOfWeight" json:"UnitOfWeight"`
	DeliverAgent        string       `protobuf:"bytes,11,opt,name=DeliverAgent" json:"DeliverAgent"`
	ReceiveAgent        string       `protobuf:"bytes,12,opt,name=ReceiveAgent" json:"ReceiveAgent"`
	Container           string       `protobuf:"bytes,13,opt,name=Container" json:"Container"`
	ContainerSeal       string       `protobuf:"bytes,14,opt,name=ContainerSeal" json:"ContainerSeal"`
	ContainerMode       string       `protobuf:"bytes,15,opt,name=ContainerMode" json:"ContainerMode"`
	ContainerType       string       `protobuf:"bytes,16,opt,name=ContainerType" json:"ContainerType"`
	Volume              string       `protobuf:"bytes,17,opt,name=Volume" json:"Volume"`
	UnitOfVolume        string       `protobuf:"bytes,18,opt,name=UnitOfVolume" json:"UnitOfVolume"`
	NotifyAddress       string       `protobuf:"bytes,19,opt,name=NotifyAddress" json:"NotifyAddress"`
	DescOfGoods         string       `protobuf:"bytes,20,opt,name=DescOfGoods" json:"DescOfGoods"`
	GrossWeight         string       `protobuf:"varint,5,opt,name=GrossWeight" json:"GrossWeight"`
	FreightPayableAmt   string       `protobuf:"varint,6,opt,name=FreightPayableAmt" json:"FreightPayableAmt"`
	FreightAdvAmt       string       `protobuf:"varint,7,opt,name=FreightAdvAmt" json:"FreightAdvAmt"`
	GeneralInstructions string       `protobuf:"bytes,21,opt,name=GeneralInstructions" json:"GeneralInstructions"`
	DateShipped         string       `protobuf:"bytes,22,opt,name=DateShipped" json:"DateShipped"`
	IssueDetails        IssueDetails `json:"IssueDetails"`
	NumBol              string       `protobuf:"varint,8,opt,name=NumBol" json:"NumBol"`
	MasterInfo          MasterInfo   `json:"MasterInfo"`
	AgentForMaster      AgentMaster  `json:"AgentForMaster"`
	AgentForOwner       AgentOwner   `json:"AgentForOwner"`
	EncryptionMetaData  string       `json:"EncryptionMetaData"`
}

// Date struct
type Date struct {
	Type   string
	Format string
}

// IssueDetails struct
type IssueDetails struct {
	PlaceOfIssue string `json:"PlaceOfIssue"`
	DateOfIssue  string `json:"DateOfIssue"`
}

// MasterInfo struct
type MasterInfo struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Sig       string `json:"Sig"`
}

// AgentMaster struct
type AgentMaster struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Sig       string `json:"Sig"`
}

// AgentOwner struct
type AgentOwner struct {
	FirstName             string `json:"FirstName"`
	LastName              string `json:"LastName"`
	Sig                   string `json:"Sig"`
	ConditionsForCarriage string `json:"ConditionsForCarriage"`
}

// =================================================
// Blockfreight™ | The blockchain of global freight.
// =================================================

// BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB
// BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB
// BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB
// BBBBBBB                    BBBBBBBBBBBBBBBBBBB
// BBBBBBB                       BBBBBBBBBBBBBBBB
// BBBBBBB                        BBBBBBBBBBBBBBB
// BBBBBBB       BBBBBBBBB        BBBBBBBBBBBBBBB
// BBBBBBB       BBBBBBBBB        BBBBBBBBBBBBBBB
// BBBBBBB       BBBBBBB         BBBBBBBBBBBBBBBB
// BBBBBBB                     BBBBBBBBBBBBBBBBBB
// BBBBBBB                        BBBBBBBBBBBBBBB
// BBBBBBB       BBBBBBBBBB        BBBBBBBBBBBBBB
// BBBBBBB       BBBBBBBBBBB       BBBBBBBBBBBBBB
// BBBBBBB       BBBBBBBBBB        BBBBBBBBBBBBBB
// BBBBBBB       BBBBBBBBB        BBB       BBBBB
// BBBBBBB                       BBBB       BBBBB
// BBBBBBB                    BBBBBBB       BBBBB
// BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB
// BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB
// BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB

// ==================================================
// Blockfreight™ | The blockchain for global freight.
// ==================================================
