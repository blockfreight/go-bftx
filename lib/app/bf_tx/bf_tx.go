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
	"crypto/ecdsa"  // Implements the Elliptic Curve Digital Signature Algorithm, as defined in FIPS 186-3.
	"crypto/sha256" // Implements the SHA256 Algorithm for Hash.
	"encoding/json" // Implements encoding and decoding of JSON as defined in RFC 4627.
	"fmt"           // Implements formatted I/O with functions analogous to C's printf and scanf.

	// ====================
	// Third-party packages
	// ====================
	"github.com/davecgh/go-spew/spew" // Implements a deep pretty printer for Go data structures to aid in debugging.

	// ======================
	// Blockfreight™ packages
	// ======================
	"github.com/blockfreight/blockfreight-alpha/lib/pkg/common" // Implements common functions for Blockfreight™
)

// SetBF_TX receives the path of a JSON, reads it and returns the BF_TX structure with all attributes.
func SetBF_TX(jsonpath string) (BF_TX, error) {
	var bf_tx BF_TX
	file, err := common.ReadJSON(jsonpath)
	if err != nil {
		return bf_tx, err
	}
	json.Unmarshal(file, &bf_tx)
	return bf_tx, nil
}

//HashBF_TX hashes the BF_TX object
func HashBF_TX(bf_tx BF_TX) ([]byte, error) {
	bf_tx_bytes := []byte(fmt.Sprintf("%v", bf_tx))

	hash := sha256.New()
	hash.Write(bf_tx_bytes)

	return hash.Sum(nil), nil
}

func HashBF_TX_salt(hash []byte, salt []byte) []byte {
	sha := sha256.New()
	sha.Write(append(hash[:], salt[:]...))
	return sha.Sum(nil)
}

// BF_TXContent receives the BF_TX structure, applies it the json.Marshal procedure and return the content of the BF_TX JSON.
func BF_TXContent(bf_tx BF_TX) (string, error) {
	jsonContent, err := json.Marshal(bf_tx)
	return string(jsonContent), err
}

// PrintBF_TX receives a BF_TX and prints it clearly.
func PrintBF_TX(bf_tx BF_TX) {
	spew.Dump(bf_tx)
}

// State reports the current state of a BF_TX
func State(bf_tx BF_TX) string {
	if bf_tx.Transmitted {
		return "Transmitted!"
	} else if bf_tx.Verified {
		return "Signed!"
	} else {
		return "Constructed!"
	}
}

// Reinitialize set the default values to the Blockfreight attributes of BF_TX
func Reinitialize(bf_tx BF_TX) BF_TX {
	bf_tx.PrivateKey.Curve = nil
	bf_tx.PrivateKey.X = nil
	bf_tx.PrivateKey.Y = nil
	bf_tx.PrivateKey.D = nil
	bf_tx.Signhash = nil
	bf_tx.Signature = ""
	bf_tx.Verified = false
	bf_tx.Transmitted = false
	return bf_tx
}

// BF_TX structure respresents an logical abstraction of a Blockfreight™ Transaction.
type BF_TX struct {
	// =========================
	// Bill of Lading attributes
	// =========================
	Type       string
	Properties Properties

	// ===================================
	// Blockfreight Transaction attributes
	// ===================================
	Id          string
	PrivateKey  ecdsa.PrivateKey
	Signhash    []uint8
	Signature   string
	Verified    bool
	Transmitted bool
	Amendment   string
}

type Properties struct {
	Shipper              Shipper
	Bol_Num              BolNum
	Ref_Num              RefNum
	Consignee            Consignee
	Vessel               Vessel
	Port_of_Loading      PortLoading
	Port_of_Discharge    PortDischarge
	Notify_Address       NotifyAddress
	Desc_of_Goods        DescGoods
	Gross_Weight         GrossWeight
	Freight_Payable_Amt  FreightPayableAmt
	Freight_Adv_Amt      FreightAdvAmt
	General_Instructions GeneralInstructions
	Date_Shipped         Date
	Issue_Details        IssueDetails
	Num_Bol              NumBol
	Master_Info          MasterInfo
	Agent_for_Master     AgentMaster
	Agent_for_Owner      AgentOwner
}

type Shipper struct {
	Type string
}

type BolNum struct {
	Type int
}

type RefNum struct {
	Type int
}

type Consignee struct {
	Type string //Null
}

type Vessel struct {
	Type int
}

type PortLoading struct {
	Type int
}

type PortDischarge struct {
	Type int
}

type NotifyAddress struct {
	Type string
}

type DescGoods struct {
	Type string
}

type GrossWeight struct {
	Type int
}

type FreightPayableAmt struct {
	Type int
}

type FreightAdvAmt struct {
	Type int
}

type GeneralInstructions struct {
	Type string
}

type Date struct {
	Type   int
	Format string
}

type IssueDetails struct {
	Type       string
	Properties IssueDetailsProperties
}

type IssueDetailsProperties struct {
	Place_of_Issue PlaceIssue
	Date_of_Issue  Date
}

type PlaceIssue struct {
	Type string
}

type NumBol struct {
	Type int
}

type MasterInfo struct {
	Type       string
	Properties MasterInfoProperties
}

type MasterInfoProperties struct {
	First_Name FirstName
	Last_Name  LastName
	Sig        Sig
}

type AgentMaster struct {
	Type       string
	Properties AgentMasterProperties
}

type AgentMasterProperties struct {
	First_Name FirstName
	Last_Name  LastName
	Sig        Sig
}

type AgentOwner struct {
	Type       string
	Properties AgentOwnerProperties
}

type AgentOwnerProperties struct {
	First_Name              FirstName
	Last_Name               LastName
	Sig                     Sig
	Conditions_for_Carriage ConditionsCarriage
}

type FirstName struct {
	Type string
}

type LastName struct {
	Type string
}

type Sig struct {
	Type string
}

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
