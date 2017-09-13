// File: ./blockfreight/lib/validator/validator.go
// Summary: Application code for Blockfreight™ | The blockchain of global freight.
// License: MIT License
// Company: Blockfreight, Inc.
// Author: Julian Nunez, Neil Tran, Julian Smith & contributors
// Site: https://blockfreight.com
// Support: <support@blockfreight.com>

// Copyright 2017 Blockfreight, Inc.

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
    "errors"    // Implements functions to manipulate errors.
    "reflect"   // Implements run-time reflection, allowing a program to manipulate objects with arbitrary types. 
    
    // ======================
    // Blockfreight™ packages
    // ======================
    "github.com/blockfreight/blockfreight-alpha/blockfreight/lib/bf_tx" // Defines the Blockfreight™ Transaction (BF_TX) transaction standard and provides some useful functions to work with the BF_TX.
)

// ValidateBf_Tx is a function that receives the BF_TX and return the proper message according with the result of ValidateFields function.
func ValidateBf_Tx(bf_tx bf_tx.BF_TX) (string, error) {
    var espErr error

    valid, err := ValidateFields(bf_tx)
    if valid {
        return "Success! [OK]", nil
    } else {
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

    support: support@blockfreight.com`, espErr
    }
}

// ValidateFields is a function that receives the BF_TX, validates every field in the BF_TX and return true or false, and a message if some field is wrong.
func ValidateFields(bf_tx bf_tx.BF_TX) (bool, string) {
    if reflect.TypeOf(bf_tx.Type) != reflect.TypeOf("s") {
        return false, "bf_tx.Type is not a string."
    }
    if reflect.TypeOf(bf_tx.Properties.Shipper.Type) != reflect.TypeOf("s") {
        return false, "bf_tx.Properties.Shipper.Type is not a string."
    }
    if (reflect.TypeOf(bf_tx.Properties.Bol_Num.Type) != reflect.TypeOf(1)) || bf_tx.Properties.Bol_Num.Type == 0 {
        return false, "bf_tx.Properties.Bol_Num.Type is not a number."
    }
    if (reflect.TypeOf(bf_tx.Properties.Ref_Num.Type) != reflect.TypeOf(1)) || bf_tx.Properties.Ref_Num.Type == 0 {
        return false, "bf_tx.Properties.Ref_Num.Type is not a number."
    }
    if reflect.TypeOf(bf_tx.Properties.Consignee.Type) != reflect.TypeOf("s") {
        return false, "bf_tx.Properties.Consignee.Type is not a string."
    }
    if (reflect.TypeOf(bf_tx.Properties.Vessel.Type) != reflect.TypeOf(1)) || bf_tx.Properties.Vessel.Type == 0 {
        return false, "bf_tx.Properties.Vessel.Type is not a number."
    }
    if (reflect.TypeOf(bf_tx.Properties.Port_of_Loading.Type) != reflect.TypeOf(1)) || bf_tx.Properties.Port_of_Loading.Type == 0 {
        return false, "bf_tx.Properties.Port_of_Loading.Type is not a number."
    }
    if (reflect.TypeOf(bf_tx.Properties.Port_of_Discharge.Type) != reflect.TypeOf(1)) || bf_tx.Properties.Port_of_Discharge.Type == 0 {
        return false, "bf_tx.Properties.Port_of_Discharge.Type is not a number."
    }
    if reflect.TypeOf(bf_tx.Properties.Notify_Address.Type) != reflect.TypeOf("s") {
        return false, "bf_tx.Properties.Notify_Address.Type is not a string."
    }
    if reflect.TypeOf(bf_tx.Properties.Desc_of_Goods.Type) != reflect.TypeOf("s") {
        return false, "bf_tx.Properties.Desc_of_Goods.Type is not a string."
    }
    if (reflect.TypeOf(bf_tx.Properties.Gross_Weight.Type) != reflect.TypeOf(1)) || bf_tx.Properties.Gross_Weight.Type == 0 {
        return false, "bf_tx.Properties.Gross_Weight.Type is not a number."
    }
    if (reflect.TypeOf(bf_tx.Properties.Freight_Payable_Amt.Type) != reflect.TypeOf(1)) || bf_tx.Properties.Freight_Payable_Amt.Type == 0 {
        return false, "bf_tx.Properties.Freight_Payable_Amt.Type is not a number."
    }
    if (reflect.TypeOf(bf_tx.Properties.Freight_Adv_Amt.Type) != reflect.TypeOf(1)) || bf_tx.Properties.Freight_Adv_Amt.Type == 0 {
        return false, "bf_tx.Properties.Freight_Adv_Amt.Type is not a number."
    }
    if reflect.TypeOf(bf_tx.Properties.General_Instructions.Type) != reflect.TypeOf("s") {
        return false, "bf_tx.Properties.General_Instructions.Type is not a string."
    }
    if (reflect.TypeOf(bf_tx.Properties.Date_Shipped.Type) != reflect.TypeOf(1)) || bf_tx.Properties.Date_Shipped.Type == 0 {
        return false, "bf_tx.Properties.Date_Shipped.Type is not a number."
    }
    if reflect.TypeOf(bf_tx.Properties.Date_Shipped.Format) != reflect.TypeOf("s") {
        return false, "bf_tx.Properties.Date_Shipped.Format is not a date format."
    }
    if reflect.TypeOf(bf_tx.Properties.Issue_Details.Type) != reflect.TypeOf("s") {
        return false, "bf_tx.Properties.Issue_Details.Type is not a string."
    }
    if reflect.TypeOf(bf_tx.Properties.Issue_Details.Properties.Place_of_Issue.Type) != reflect.TypeOf("s") {
        return false, "bf_tx.Properties.Issue_Details.Properties.Place_of_Issue.Type is not a string."
    }
    if (reflect.TypeOf(bf_tx.Properties.Issue_Details.Properties.Date_of_Issue.Type) != reflect.TypeOf(1)) || bf_tx.Properties.Issue_Details.Properties.Date_of_Issue.Type == 0 {
        return false, "bf_tx.Properties.Issue_Details.Properties.Date_of_Issue.Type is not a number."
    }
    if reflect.TypeOf(bf_tx.Properties.Issue_Details.Properties.Date_of_Issue.Format) != reflect.TypeOf("s") {
        return false, "bf_tx.Properties.Issue_Details.Properties.Place_of_Issue.Format is not a date format."
    }
    if (reflect.TypeOf(bf_tx.Properties.Num_Bol.Type) != reflect.TypeOf(1)) || bf_tx.Properties.Num_Bol.Type == 0 {
        return false, "bf_tx.Properties.Num_Bol.Type is not a number."
    }
    if reflect.TypeOf(bf_tx.Properties.Master_Info.Type) != reflect.TypeOf("s") {
        return false, "bf_tx.Properties.Master_Info.Type is not a string."
    }
    if reflect.TypeOf(bf_tx.Properties.Master_Info.Properties.First_Name.Type) != reflect.TypeOf("s") {
        return false, "bf_tx.Properties.Master_Info.Properties.First_Name.Type is not a string."
    }
    if reflect.TypeOf(bf_tx.Properties.Master_Info.Properties.Last_Name.Type) != reflect.TypeOf("s") {
        return false, "bf_tx.Properties.Master_Info.Properties.Last_Name.Type is not a string."
    }
    if reflect.TypeOf(bf_tx.Properties.Master_Info.Properties.Sig.Type) != reflect.TypeOf("s") {
        return false, "bf_tx.Properties.Master_Info.Properties.Sig.Type is not a string."
    }
    if reflect.TypeOf(bf_tx.Properties.Agent_for_Master.Type) != reflect.TypeOf("s") {
        return false, "bf_tx.Properties.Agent_for_Master.Type is not a string."
    }
    if reflect.TypeOf(bf_tx.Properties.Agent_for_Master.Properties.First_Name.Type) != reflect.TypeOf("s") {
        return false, "bf_tx.Properties.Agent_for_Master.Properties.First_Name.Type is not a string."
    }
    if reflect.TypeOf(bf_tx.Properties.Agent_for_Master.Properties.Last_Name.Type) != reflect.TypeOf("s") {
        return false, "bf_tx.Properties.Agent_for_Master.Properties.Last_Name.Type is not a string."
    }
    if reflect.TypeOf(bf_tx.Properties.Agent_for_Master.Properties.Sig.Type) != reflect.TypeOf("s") {
        return false, "bf_tx.Properties.Agent_for_Master.Properties.Sig.Type is not a string."
    }
    if reflect.TypeOf(bf_tx.Properties.Agent_for_Owner.Type) != reflect.TypeOf("s") {
        return false, "bf_tx.Properties.Agent_for_Owner.Type is not a string."
    }
    if reflect.TypeOf(bf_tx.Properties.Agent_for_Owner.Properties.First_Name.Type) != reflect.TypeOf("s") {
        return false, "bf_tx.Properties.Agent_for_Owner.Properties.First_Name.Type is not a string."
    }
    if reflect.TypeOf(bf_tx.Properties.Agent_for_Owner.Properties.Last_Name.Type) != reflect.TypeOf("s") {
        return false, "bf_tx.Properties.Agent_for_Owner.Properties.Last_Name.Type is not a string."
    }
    if reflect.TypeOf(bf_tx.Properties.Agent_for_Owner.Properties.Sig.Type) != reflect.TypeOf("s") {
        return false, "bf_tx.Properties.Agent_for_Owner.Properties.Sig.Type is not a string."
    }
    if reflect.TypeOf(bf_tx.Properties.Agent_for_Owner.Properties.Conditions_for_Carriage.Type) != reflect.TypeOf("s") {
        return false, "bf_tx.Properties.Agent_for_Owner.Properties.Conditions_for_Carriage.Type is not a string."
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
