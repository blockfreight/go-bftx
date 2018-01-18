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
	"crypto/sha256"
	"encoding/json" // Implements the SHA256 Algorithm for Hash.
	// Implements encoding and decoding of JSON as defined in RFC 4627.

	"fmt" // Implements formatted I/O with functions analogous to C's printf and scanf.

	// ====================
	// Third-party packages
	// ====================
	"github.com/davecgh/go-spew/spew" // Implements a deep pretty printer for Go data structures to aid in debugging.

	// ======================
	// Blockfreight™ packages
	// ======================

	"github.com/blockfreight/go-bftx/lib/pkg/common" // Implements common functions for Blockfreight™
)

// SetBFTX receives the path of a JSON, reads it and returns the BF_TX structure with all attributes.
func SetBFTX(jsonpath string) (BF_TX, error) {
	var bftx BF_TX
	file, err := common.ReadJSON(jsonpath)
	if err != nil {
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

//GenerateBFTXUID hashes two byte arrays and returns it.
func GenerateBFTXUID(hash []byte, salt []byte) string {
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
	Id          string           `json:"Id,omitempty"`
	PrivateKey  ecdsa.PrivateKey `json:"PrivateKey,omitempty"`
	Signhash    []uint8          `json:"Signhash,omitempty"`
	Signature   string           `json:"Signature,omitempty"`
	Verified    bool             `json:"Verified,omitempty"`
	Transmitted bool             `json:"Transmitted,omitempty"`
	Amendment   string           `json:"Amendment,omitempty"`
	Private     string           `json:"Private,omitempty"`
}

// Properties struct
type Properties struct {
	Shipper             string       `protobuf:"bytes,1,opt,name=Shipper" json:"Shipper,omitempty"`
	BolNum              int          `protobuf:"varint,1,opt,name=BolNum" json:"BolNum,omitempty"`
	RefNum              int          `protobuf:"varint,2,opt,name=RefNum" json:"RefNum,omitempty"`
	Consignee           string       `protobuf:"bytes,2,opt,name=Consignee" json:"Consignee,omitempty"`
	HouseBill           string       `protobuf:"bytes,3,opt,name=HouseBill" json:"HouseBill,omitempty"`
	Vessel              int          `protobuf:"varint,3,opt,name=Vessel" json:"Vessel,omitempty"`
	Packages            int          `protobuf:"varint,4,opt,name=Packages" json:"Packages,omitempty"`
	PackType            string       `protobuf:"bytes,4,opt,name=PackType" json:"PackType,omitempty"`
	INCOTerms           string       `protobuf:"bytes,5,opt,name=INCOTerms" json:"INCOTerms,omitempty"`
	PortOfLoading       string       `protobuf:"bytes,6,opt,name=PortOfLoading" json:"PortOfLoading,omitempty"`
	PortOfDischarge     string       `protobuf:"bytes,7,opt,name=PortOfDischarge" json:"PortOfDischarge,omitempty"`
	Destination         string       `protobuf:"bytes,8,opt,name=Destination" json:"Destination,omitempty"`
	MarksAndNumbers     string       `protobuf:"bytes,9,opt,name=MarksAndNumbers" json:"MarksAndNumbers,omitempty"`
	UnitOfWeight        string       `protobuf:"bytes,10,opt,name=UnitOfWeight" json:"UnitOfWeight,omitempty"`
	DeliverAgent        string       `protobuf:"bytes,11,opt,name=DeliverAgent" json:"DeliverAgent,omitempty"`
	ReceiveAgent        string       `protobuf:"bytes,12,opt,name=ReceiveAgent" json:"ReceiveAgent,omitempty"`
	Container           string       `protobuf:"bytes,13,opt,name=Container" json:"Container,omitempty"`
	ContainerSeal       string       `protobuf:"bytes,14,opt,name=ContainerSeal" json:"ContainerSeal,omitempty"`
	ContainerMode       string       `protobuf:"bytes,15,opt,name=ContainerMode" json:"ContainerMode,omitempty"`
	ContainerType       string       `protobuf:"bytes,16,opt,name=ContainerType" json:"ContainerType,omitempty"`
	Volume              float64      `protobuf:"bytes,17,opt,name=Volume" json:"Volume,omitempty"`
	UnitOfVolume        string       `protobuf:"bytes,18,opt,name=UnitOfVolume" json:"UnitOfVolume,omitempty`
	NotifyAddress       string       `protobuf:"bytes,19,opt,name=NotifyAddress" json:"NotifyAddress,omitempty"`
	DescOfGoods         string       `protobuf:"bytes,20,opt,name=DescOfGoods" json:"DescOfGoods,omitempty"`
	GrossWeight         float64      `protobuf:"varint,5,opt,name=GrossWeight" json:"GrossWeight,omitempty"`
	FreightPayableAmt   int          `protobuf:"varint,6,opt,name=FreightPayableAmt" json:"FreightPayableAmt,omitempty"`
	FreightAdvAmt       int          `protobuf:"varint,7,opt,name=FreightAdvAmt" json:"FreightAdvAmt,omitempty"`
	GeneralInstructions string       `protobuf:"bytes,21,opt,name=GeneralInstructions" json:"GeneralInstructions,omitempty"`
	DateShipped         string       `protobuf:"bytes,22,opt,name=DateShipped" json:"DateShipped"`
	IssueDetails        IssueDetails `json:"IssueDetails"`
	NumBol              int          `protobuf:"varint,8,opt,name=NumBol" json:"NumBol,omitempty"`
	MasterInfo          MasterInfo   `json:"MasterInfo"`
	AgentForMaster      AgentMaster  `json:"AgentForMaster"`
	AgentForOwner       AgentOwner   `json:"AgentForOwner"`
}

// Shipper struct
type Shipper struct {
	Type string
}

// BolNum struct
type BolNum struct {
	Type int
}

// RefNum struct
type RefNum struct {
	Type int
}

// Consignee struct
type Consignee struct {
	Type string //Null
}

// Vessel struct
type Vessel struct {
	Type int
}

// PortLoading struct
type PortLoading struct {
	Type int
}

// PortDischarge struct
type PortDischarge struct {
	Type int
}

// NotifyAddress struct
type NotifyAddress struct {
	Type string
}

// DescGoods struct
type DescGoods struct {
	Type string
}

// GrossWeight struct
type GrossWeight struct {
	Type int
}

// FreightPayableAmt struct
type FreightPayableAmt struct {
	Type int
}

// FreightAdvAmt struct
type FreightAdvAmt struct {
	Type int
}

// GeneralInstructions struct
type GeneralInstructions struct {
	Type string
}

// Date struct
type Date struct {
	Type   int
	Format string
}

// IssueDetails struct
type IssueDetails struct {
	PlaceOfIssue string `json:"PlaceOfIssue"`
	DateOfIssue  string `json:"DateOfIssue"`
}

// PlaceIssue struct
type PlaceIssue struct {
	Type string
}

// NumBol struct
type NumBol struct {
	Type int
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

// FirstName struct
type FirstName struct {
	Type string
}

// LastName struct
type LastName struct {
	Type string
}

// Sig struct
type Sig struct {
	Type string
}

// ConditionsCarriage struct
type ConditionsCarriage struct {
	Type string
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
