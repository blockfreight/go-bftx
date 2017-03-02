package validator

import (
    "fmt"
    "reflect"
    "strconv"
    
    "github.com/blockfreight/blockfreight-alpha/blockfreight/bft/bf_tx"
    "github.com/blockfreight/blockfreight-alpha/blockfreight/bft/leveldb"
)

func ValidateBf_Tx(bf_tx bf_tx.BF_TX) string {
    //printJson := true
    //examplePath := "./files/bf_tx_example.json"
    espErr := ""
    
    valid, err := ValidateFields(bf_tx)
    if valid {
        return "Success! [OK]"
    } else {
        if err != "" {
            espErr = `
    Specific Error [01]:
    `+err
    }
        return `
    Blockfreight, Inc. © 2017. Open Source (MIT) License.

    Error [01]:

    Invalid structure in JSON provided. JSON 结构无效.
    Struttura JSON non valido. هيكل JSON صالح. 無効なJSON構造

    support: support@blockfreight.com`+espErr
    }
}

func RecordOnDB(/*id string, */json string) bool {  //TODO: Check the id
    db_path := "bft-db"
    db, err := leveldb.OpenDB(db_path)
    defer leveldb.CloseDB(db)
    
    //Get the number of bf_tx on DB
    var n int
    n, err = leveldb.Iterate(db)

    leveldb.HandleError(err, "Create or Open Database")
    //fmt.Println("Database created / open on "+db_path)
    
    err = leveldb.InsertBF_TX(strconv.Itoa(n+1), json, db)
    //err = leveldb.InsertBF_TX(id, json, db)    //TODO: Check the id

    //Iteration
    n, err = leveldb.Iterate(db)
    leveldb.HandleError(err, "Iteration")
    fmt.Println("Total: "+strconv.Itoa(n))  

    return true
}

func ValidateFields(bf_tx bf_tx.BF_TX) (bool, string){
    if (reflect.TypeOf(bf_tx.Type) != reflect.TypeOf("s")){
        return false, "bf_tx.Type is not a string."
    }
    if (reflect.TypeOf(bf_tx.Properties.Shipper.Type) != reflect.TypeOf("s")){
        return false, "bf_tx.Properties.Shipper.Type is not a string."
    }
    if (reflect.TypeOf(bf_tx.Properties.Bol_Num.Type) != reflect.TypeOf(1)) || bf_tx.Properties.Bol_Num.Type == 0 {
        return false, "bf_tx.Properties.Bol_Num.Type is not a nunber."
    }
    if (reflect.TypeOf(bf_tx.Properties.Ref_Num.Type) != reflect.TypeOf(1)) || bf_tx.Properties.Ref_Num.Type == 0 {
        return false, "bf_tx.Properties.Ref_Num.Type is not a number."
    }
    if (reflect.TypeOf(bf_tx.Properties.Consignee.Type) != reflect.TypeOf("s")){
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
    if (reflect.TypeOf(bf_tx.Properties.Notify_Address.Type) != reflect.TypeOf("s")){
        return false, "bf_tx.Properties.Notify_Address.Type is not a string."
    }
    if (reflect.TypeOf(bf_tx.Properties.Desc_of_Goods.Type) != reflect.TypeOf("s")){
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
    if (reflect.TypeOf(bf_tx.Properties.General_Instructions.Type) != reflect.TypeOf("s")){
        return false, "bf_tx.Properties.General_Instructions.Type is not a string."
    }
    if (reflect.TypeOf(bf_tx.Properties.Date_Shipped.Type) != reflect.TypeOf(1)) || bf_tx.Properties.Date_Shipped.Type == 0 {
        return false, "bf_tx.Properties.Date_Shipped.Type is not a number."
    }
    if (reflect.TypeOf(bf_tx.Properties.Date_Shipped.Format) != reflect.TypeOf("s")){
        return false, "bf_tx.Properties.Date_Shipped.Format is not a date format."
    }
    if (reflect.TypeOf(bf_tx.Properties.Issue_Details.Type) != reflect.TypeOf("s")){
        return false, "bf_tx.Properties.Issue_Details.Type is not a string."
    }
    if (reflect.TypeOf(bf_tx.Properties.Issue_Details.Properties.Place_of_Issue.Type) != reflect.TypeOf("s")){
        return false, "bf_tx.Properties.Issue_Details.Properties.Place_of_Issue.Type is not a string."
    }
    if (reflect.TypeOf(bf_tx.Properties.Issue_Details.Properties.Date_of_Issue.Type) != reflect.TypeOf(1)) || bf_tx.Properties.Issue_Details.Properties.Date_of_Issue.Type == 0 {
        return false, "bf_tx.Properties.Issue_Details.Properties.Date_of_Issue.Type is not a number."
    }
    if (reflect.TypeOf(bf_tx.Properties.Issue_Details.Properties.Date_of_Issue.Format) != reflect.TypeOf("s")){
        return false, "bf_tx.Properties.Issue_Details.Properties.Place_of_Issue.Format is not a date format."
    }
    if (reflect.TypeOf(bf_tx.Properties.Num_Bol.Type) != reflect.TypeOf(1)) || bf_tx.Properties.Num_Bol.Type == 0 {
        return false, "bf_tx.Properties.Num_Bol.Type is not a number."
    }
    if (reflect.TypeOf(bf_tx.Properties.Master_Info.Type) != reflect.TypeOf("s")){
        return false, "bf_tx.Properties.Master_Info.Type is not a string."
    }
    if (reflect.TypeOf(bf_tx.Properties.Master_Info.Properties.First_Name.Type) != reflect.TypeOf("s")){
        return false, "bf_tx.Properties.Master_Info.Properties.First_Name.Type is not a string."
    }
    if (reflect.TypeOf(bf_tx.Properties.Master_Info.Properties.Last_Name.Type) != reflect.TypeOf("s")){
        return false, "bf_tx.Properties.Master_Info.Properties.Last_Name.Type is not a string."
    }
    if (reflect.TypeOf(bf_tx.Properties.Master_Info.Properties.Sig.Type) != reflect.TypeOf("s")){
        return false, "bf_tx.Properties.Master_Info.Properties.Sig.Type is not a string."
    }
    if (reflect.TypeOf(bf_tx.Properties.Agent_for_Master.Type) != reflect.TypeOf("s")){
        return false, "bf_tx.Properties.Agent_for_Master.Type is not a string."
    }
    if (reflect.TypeOf(bf_tx.Properties.Agent_for_Master.Properties.First_Name.Type) != reflect.TypeOf("s")){
        return false, "bf_tx.Properties.Agent_for_Master.Properties.First_Name.Type is not a string."
    }
    if (reflect.TypeOf(bf_tx.Properties.Agent_for_Master.Properties.Last_Name.Type) != reflect.TypeOf("s")){
        return false, "bf_tx.Properties.Agent_for_Master.Properties.Last_Name.Type is not a string."
    }
    if (reflect.TypeOf(bf_tx.Properties.Agent_for_Master.Properties.Sig.Type) != reflect.TypeOf("s")){
        return false, "bf_tx.Properties.Agent_for_Master.Properties.Sig.Type is not a string."
    }
    if (reflect.TypeOf(bf_tx.Properties.Agent_for_Owner.Type) != reflect.TypeOf("s")){
        return false, "bf_tx.Properties.Agent_for_Owner.Type is not a string."
    }
    if (reflect.TypeOf(bf_tx.Properties.Agent_for_Owner.Properties.First_Name.Type) != reflect.TypeOf("s")){
        return false, "bf_tx.Properties.Agent_for_Owner.Properties.First_Name.Type is not a string."
    }
    if (reflect.TypeOf(bf_tx.Properties.Agent_for_Owner.Properties.Last_Name.Type) != reflect.TypeOf("s")){
        return false, "bf_tx.Properties.Agent_for_Owner.Properties.Last_Name.Type is not a string."
    }
    if (reflect.TypeOf(bf_tx.Properties.Agent_for_Owner.Properties.Sig.Type) != reflect.TypeOf("s")){
        return false, "bf_tx.Properties.Agent_for_Owner.Properties.Sig.Type is not a string."
    }
    if (reflect.TypeOf(bf_tx.Properties.Agent_for_Owner.Properties.Conditions_for_Carriage.Type) != reflect.TypeOf("s")){
        return false, "bf_tx.Properties.Agent_for_Owner.Properties.Conditions_for_Carriage.Type is not a string."
    }
    return true, ""
}