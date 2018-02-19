// File: ./blockfreight/lib/validator/validator.go
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

// Package validator is a package that provides functions to assure the input JSON is correct.
package validator

import (
	// =======================
	// Golang Standard library
	// =======================
	"errors"  // Implements functions to manipulate errors.
	"reflect" // Implements run-time reflection, allowing a program to manipulate objects with arbitrary types.

	// ======================
	// Blockfreight™ packages
	// ======================
	"github.com/blockfreight/go-bftx/lib/app/bf_tx" // Defines the Blockfreight™ Transaction (BF_TX) transaction standard and provides some useful functions to work with the BF_TX.
)

// ValidateBFTX is a function that receives the BF_TX and return the proper message according with the result of ValidateFields function.
func ValidateBFTX(bftx bf_tx.BF_TX) (string, error) {
	var espErr error

	valid, err := ValidateFields(bftx)
	if valid {
		return "Success! [OK]", nil
	}

	if err != "" {
		espErr = errors.New(`
    Specific Error [01]:
    ` + err)
	}
	return `
    Blockfreight, Inc. © 2017. Open Source (MIT) License.

    Error [01]:

    Invalid structure in JSON provided. JSON 结构无效.
    Struttura JSON non valido. هيكل JSON صالح. 無効なJSON構造.
	Estructura inválida en el JSON dado.
	Estrutura inválida no JSON enviado.

    support: support@blockfreight.com`, espErr
}

// ValidateFields is a function that receives the BF_TX, validates every field in the BF_TX and return true or false, and a message if some field is wrong.
func ValidateFields(bftx bf_tx.BF_TX) (bool, string) {
	if reflect.TypeOf(bftx.Properties.Shipper) != reflect.TypeOf("s") {
		return false, "bftx.Properties.Shipperis not a string."
	}
	// if reflect.TypeOf(bftx.Properties.BolNum).Kind() != reflect.Int {
	// 	return false, "bftx.Properties.BolNum is not a number."
	// }
	// if reflect.TypeOf(bftx.Properties.RefNum).Kind() != reflect.Int {
	// 	return false, "bftx.Properties.RefNum is not a number."
	// }
	if reflect.TypeOf(bftx.Properties.Consignee) != reflect.TypeOf("s") {
		return false, "bftx.Properties.Consignee is not a string."
	}
	// if reflect.TypeOf(bftx.Properties.Vessel).Kind() != reflect.Int {
	// 	return false, "bftx.Properties.Vessel is not a number."
	// }
	if reflect.TypeOf(bftx.Properties.PortOfLoading) != reflect.TypeOf("s") {
		return false, "bftx.Properties.PortOfLoading is not a number."
	}
	if reflect.TypeOf(bftx.Properties.PortOfDischarge) != reflect.TypeOf("s") {
		return false, "bftx.Properties.PortOfDischarge is not a number."
	}
	if reflect.TypeOf(bftx.Properties.UnitOfVolume) != reflect.TypeOf("s") {
		return false, "bftx.Properties.UnitOfVolume is not a string."
	}
	if reflect.TypeOf(bftx.Properties.NotifyAddress) != reflect.TypeOf("s") {
		return false, "bftx.Properties.NotifyAddress is not a string."
	}
	if reflect.TypeOf(bftx.Properties.DescOfGoods) != reflect.TypeOf("s") {
		return false, "bftx.Properties.DescOfGoods is not a string."
	}
	// if reflect.TypeOf(bftx.Properties.GrossWeight).Kind() != reflect.Float64 {
	// 	return false, "bftx.Properties.GrossWeight is not a float64."
	// }
	// if reflect.TypeOf(bftx.Properties.FreightPayableAmt).Kind() != reflect.Int {
	// 	return false, "bftx.Properties.FreightPayableAmt is not an int."
	// }
	// if reflect.TypeOf(bftx.Properties.FreightAdvAmt).Kind() != reflect.Int {
	// 	return false, "bftx.Properties.FreightAdvAmt is not an int."
	// }
	if reflect.TypeOf(bftx.Properties.GeneralInstructions) != reflect.TypeOf("s") {
		return false, "bftx.Properties.GeneralInstructions is not a string."
	}
	if reflect.TypeOf(bftx.Properties.DateShipped) != reflect.TypeOf("s") {
		return false, "bftx.Properties.DateShipped is not a string."
	}
	if reflect.TypeOf(bftx.Properties.IssueDetails.DateOfIssue) != reflect.TypeOf("s") {
		return false, "bftx.Properties.IssueDetails.Properties.PlaceOfIssue is not a string."
	}
	if reflect.TypeOf(bftx.Properties.IssueDetails.PlaceOfIssue) != reflect.TypeOf("s") {
		return false, "bftx.Properties.IssueDetails.Properties.PlaceOfIssue is not a string."
	}
	// if reflect.TypeOf(bftx.Properties.NumBol).Kind() != reflect.Int {
	// 	return false, "bftx.Properties.NumBol is not a number."
	// }

	if reflect.TypeOf(bftx.Properties.MasterInfo.FirstName) != reflect.TypeOf("s") {
		return false, "bftx.Properties.MasterInfo.Properties.FirstName is not a string."
	}
	if reflect.TypeOf(bftx.Properties.MasterInfo.LastName) != reflect.TypeOf("s") {
		return false, "bftx.Properties.MasterInfo.Properties.LastName is not a string."
	}
	if reflect.TypeOf(bftx.Properties.MasterInfo.Sig) != reflect.TypeOf("s") {
		return false, "bftx.Properties.MasterInfo.Properties.Sig is not a string."
	}

	if reflect.TypeOf(bftx.Properties.AgentForMaster.FirstName) != reflect.TypeOf("s") {
		return false, "bftx.Properties.AgentForMaster.Properties.FirstName is not a string."
	}
	if reflect.TypeOf(bftx.Properties.AgentForMaster.LastName) != reflect.TypeOf("s") {
		return false, "bftx.Properties.AgentForMaster.Properties.LastName is not a string."
	}
	if reflect.TypeOf(bftx.Properties.AgentForMaster.Sig) != reflect.TypeOf("s") {
		return false, "bftx.Properties.AgentForMaster.Properties.Sig is not a string."
	}

	if reflect.TypeOf(bftx.Properties.AgentForOwner.FirstName) != reflect.TypeOf("s") {
		return false, "bftx.Properties.AgentForOwner.Properties.FirstName is not a string."
	}
	if reflect.TypeOf(bftx.Properties.AgentForOwner.LastName) != reflect.TypeOf("s") {
		return false, "bftx.Properties.AgentForOwner.Properties.LastName is not a string."
	}
	if reflect.TypeOf(bftx.Properties.AgentForOwner.Sig) != reflect.TypeOf("s") {
		return false, "bftx.Properties.AgentForOwner.Properties.Sig is not a string."
	}
	if reflect.TypeOf(bftx.Properties.AgentForOwner.ConditionsForCarriage) != reflect.TypeOf("s") {
		return false, "bftx.Properties.AgentForOwner.Properties.ConditionsForCarriage is not a string."
	}

	// ------------------------------------------------------------------------
	// New fields
	// ------------------------------------------------------------------------
	if reflect.TypeOf(bftx.Properties.EncryptionMetaData).Kind() != reflect.String {
		return false, "bftx.Properties.EncryptionMetaData is not a string."
	}
	if reflect.TypeOf(bftx.Properties.Consignee).Kind() != reflect.String {
		return false, "bftx.Properties.Consignee is not a string."
	}
	if reflect.TypeOf(bftx.Properties.HouseBill).Kind() != reflect.String {
		return false, "bftx.Properties.HouseBill is not a string."
	}
	if reflect.TypeOf(bftx.Properties.ReceiveAgent).Kind() != reflect.String {
		return false, "bftx.Properties.ReceiveAgent is not a string."
	}
	if reflect.TypeOf(bftx.Properties.Destination).Kind() != reflect.String {
		return false, "bftx.Properties.Destination is not a string."
	}
	if reflect.TypeOf(bftx.Properties.MarksAndNumbers).Kind() != reflect.String {
		return false, "bftx.Properties.MarksAndNumbers is not a string."
	}
	if reflect.TypeOf(bftx.Properties.UnitOfWeight).Kind() != reflect.String {
		return false, "bftx.Properties.UnitOfWeight is not a string."
	}
	// if reflect.TypeOf(bftx.Properties.Volume).Kind() != reflect.Float64 {
	// 	return false, "bftx.Properties.Volume is not a float."
	// }
	if reflect.TypeOf(bftx.Properties.Container).Kind() != reflect.String {
		return false, "bftx.Properties.Container is not a string."
	}
	if reflect.TypeOf(bftx.Properties.ContainerSeal).Kind() != reflect.String {
		return false, "bftx.Properties.ContainerSeal is not a string."
	}
	// if reflect.TypeOf(bftx.Properties.Packages).Kind() != reflect.Int {
	// 	return false, "bftx.Properties.Packages is not a string."
	// }
	if reflect.TypeOf(bftx.Properties.PackType).Kind() != reflect.String {
		return false, "bftx.Properties.PackType is not a string."
	}
	if reflect.TypeOf(bftx.Properties.INCOTerms).Kind() != reflect.String {
		return false, "bftx.Properties.INCOTerms is not a string."
	}
	if reflect.TypeOf(bftx.Properties.DeliverAgent).Kind() != reflect.String {
		return false, "bftx.Properties.DeliverAgent is not a string."
	}
	if reflect.TypeOf(bftx.Properties.ContainerMode).Kind() != reflect.String {
		return false, "bftx.Properties.ContainerMode is not a string."
	}
	if reflect.TypeOf(bftx.Properties.ContainerType).Kind() != reflect.String {
		return false, "bftx.Properties.ContainerType is not a string."
	}

	return true, ""
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
