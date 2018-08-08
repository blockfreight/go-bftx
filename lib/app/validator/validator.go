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
	"errors"
	"reflect"

	"github.com/blockfreight/go-bftx/lib/app/bf_tx"
)

// Implements run-time reflection, allowing a program to manipulate objects with arbitrary types.
// ======================
// Blockfreight™ packages
// ======================
// Defines the Blockfreight™ Transaction (BF_TX) transaction standard and provides some useful functions to work with the BF_TX.

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
	if reflect.TypeOf(bftx.Properties.Shipment.Consignee) != reflect.TypeOf("s") {
		return false, "bftx.Properties.Consignee is not a string."
	}
	// if reflect.TypeOf(bftx.Properties.Vessel).Kind() != reflect.Int {
	// 	return false, "bftx.Properties.Vessel is not a number."
	// }
	if reflect.TypeOf(bftx.Properties.Consol.PortOfLoading) != reflect.TypeOf("s") {
		return false, "bftx.Properties.PortOfLoading is not a number."
	}
	if reflect.TypeOf(bftx.Properties.Consol.PortOfDischarge) != reflect.TypeOf("s") {
		return false, "bftx.Properties.PortOfDischarge is not a number."
	}
	if reflect.TypeOf(bftx.Properties.Shipment.GoodsDescription) != reflect.TypeOf("s") {
		return false, "bftx.Properties.GoodsDescription is not a string."
	}
	if reflect.TypeOf(bftx.Properties.Shipment.Housebill).Kind() != reflect.String {
		return false, "bftx.Properties.HouseBill is not a string."
	}
	if reflect.TypeOf(bftx.Properties.Shipment.MarksAndNumbers).Kind() != reflect.String {
		return false, "bftx.Properties.MarksAndNumbers is not a string."
	}
	if reflect.TypeOf(bftx.Properties.Shipment.PackType).Kind() != reflect.String {
		return false, "bftx.Properties.PackType is not a string."
	}
	if reflect.TypeOf(bftx.Properties.Shipment.INCOTERM).Kind() != reflect.String {
		return false, "bftx.Properties.INCOTerms is not a string."
	}
	if reflect.TypeOf(bftx.Properties.Consol.ContainerMode).Kind() != reflect.String {
		return false, "bftx.Properties.ContainerMode is not a string."
	}
	if reflect.TypeOf(bftx.Properties.Consol.ContainerType).Kind() != reflect.String {
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
