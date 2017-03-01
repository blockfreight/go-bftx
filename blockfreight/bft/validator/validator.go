package validator

import (
    "fmt"
    "os"
    "io/ioutil"
    "encoding/json"
    "reflect"
    "strconv"
    
    "github.com/blockfreight/blockfreight-alpha/blockfreight/bft/bf_tx"
    "github.com/davecgh/go-spew/spew"
    "github.com/blockfreight/blockfreight-alpha/blockfreight/bft/leveldb"
    "github.com/blockfreight/blockfreight-alpha/blockfreight/ecdsa"
)

func ValidateBfTx(jsonpath string, printJson bool) string {
    //printJson := true
    //examplePath := "./files/bf_tx_example.json"
    espErr := ""
    var bftx bf_tx.BF_TX
    file := ReadJSON(jsonpath)
    json.Unmarshal(file, &bftx)

    //Sign BFTX
    bftx = ecdsa.Sign_BFTX(bftx)
    //fmt.Println(bftx.Signhash)

    if printJson { spew.Dump(bftx) }
    
    valid, err := ValidateFields(bftx)
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

func ReadJSON(path string) []byte {
    fmt.Println("\nReading "+path+"\n")
    file, e := ioutil.ReadFile(path)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }
    return file
}

func recordOnDB(json string) bool {
    db_path := "bft-db"
    db, err := leveldb.OpenDB(db_path)
    defer leveldb.CloseDB(db)
    leveldb.HandleError(err, "Create or Open Database")
    //fmt.Println("Database created / open on "+db_path)
    
    err = leveldb.InsertBFTX("1", json, db)

    //Iteration
    var n int
    n, err = leveldb.Iterate(db)
    leveldb.HandleError(err, "Iteration")
    fmt.Println("Total: "+strconv.Itoa(n))  

    return true
}

func ValidateFields(bftx bf_tx.BF_TX) (bool, string){
    if (reflect.TypeOf(bftx.Type) != reflect.TypeOf("s")){
        return false, "bftx.Type is not a string."
    }
    if (reflect.TypeOf(bftx.Properties.Shipper.Type) != reflect.TypeOf("s")){
        return false, "bftx.Properties.Shipper.Type is not a string."
    }
    if (reflect.TypeOf(bftx.Properties.Bol_Num.Type) != reflect.TypeOf(1)) || bftx.Properties.Bol_Num.Type == 0 {
        return false, "bftx.Properties.Bol_Num.Type is not a nunber."
    }
    if (reflect.TypeOf(bftx.Properties.Ref_Num.Type) != reflect.TypeOf(1)) || bftx.Properties.Ref_Num.Type == 0 {
        return false, "bftx.Properties.Ref_Num.Type is not a number."
    }
    if (reflect.TypeOf(bftx.Properties.Consignee.Type) != reflect.TypeOf("s")){
        return false, "bftx.Properties.Consignee.Type is not a string."
    }
    if (reflect.TypeOf(bftx.Properties.Vessel.Type) != reflect.TypeOf(1)) || bftx.Properties.Vessel.Type == 0 {
        return false, "bftx.Properties.Vessel.Type is not a number."
    }
    if (reflect.TypeOf(bftx.Properties.Port_of_Loading.Type) != reflect.TypeOf(1)) || bftx.Properties.Port_of_Loading.Type == 0 {
        return false, "bftx.Properties.Port_of_Loading.Type is not a number."
    }
    if (reflect.TypeOf(bftx.Properties.Port_of_Discharge.Type) != reflect.TypeOf(1)) || bftx.Properties.Port_of_Discharge.Type == 0 {
        return false, "bftx.Properties.Port_of_Discharge.Type is not a number."
    }
    if (reflect.TypeOf(bftx.Properties.Notify_Address.Type) != reflect.TypeOf("s")){
        return false, "bftx.Properties.Notify_Address.Type is not a string."
    }
    if (reflect.TypeOf(bftx.Properties.Desc_of_Goods.Type) != reflect.TypeOf("s")){
        return false, "bftx.Properties.Desc_of_Goods.Type is not a string."
    }
    if (reflect.TypeOf(bftx.Properties.Gross_Weight.Type) != reflect.TypeOf(1)) || bftx.Properties.Gross_Weight.Type == 0 {
        return false, "bftx.Properties.Gross_Weight.Type is not a number."
    }
    if (reflect.TypeOf(bftx.Properties.Freight_Payable_Amt.Type) != reflect.TypeOf(1)) || bftx.Properties.Freight_Payable_Amt.Type == 0 {
        return false, "bftx.Properties.Freight_Payable_Amt.Type is not a number."
    }
    if (reflect.TypeOf(bftx.Properties.Freight_Adv_Amt.Type) != reflect.TypeOf(1)) || bftx.Properties.Freight_Adv_Amt.Type == 0 {
        return false, "bftx.Properties.Freight_Adv_Amt.Type is not a number."
    }
    if (reflect.TypeOf(bftx.Properties.General_Instructions.Type) != reflect.TypeOf("s")){
        return false, "bftx.Properties.General_Instructions.Type is not a string."
    }
    if (reflect.TypeOf(bftx.Properties.Date_Shipped.Type) != reflect.TypeOf(1)) || bftx.Properties.Date_Shipped.Type == 0 {
        return false, "bftx.Properties.Date_Shipped.Type is not a number."
    }
    if (reflect.TypeOf(bftx.Properties.Date_Shipped.Format) != reflect.TypeOf("s")){
        return false, "bftx.Properties.Date_Shipped.Format is not a date format."
    }
    if (reflect.TypeOf(bftx.Properties.Issue_Details.Type) != reflect.TypeOf("s")){
        return false, "bftx.Properties.Issue_Details.Type is not a string."
    }
    if (reflect.TypeOf(bftx.Properties.Issue_Details.Properties.Place_of_Issue.Type) != reflect.TypeOf("s")){
        return false, "bftx.Properties.Issue_Details.Properties.Place_of_Issue.Type is not a string."
    }
    if (reflect.TypeOf(bftx.Properties.Issue_Details.Properties.Date_of_Issue.Type) != reflect.TypeOf(1)) || bftx.Properties.Issue_Details.Properties.Date_of_Issue.Type == 0 {
        return false, "bftx.Properties.Issue_Details.Properties.Date_of_Issue.Type is not a number."
    }
    if (reflect.TypeOf(bftx.Properties.Issue_Details.Properties.Date_of_Issue.Format) != reflect.TypeOf("s")){
        return false, "bftx.Properties.Issue_Details.Properties.Place_of_Issue.Format is not a date format."
    }
    if (reflect.TypeOf(bftx.Properties.Num_Bol.Type) != reflect.TypeOf(1)) || bftx.Properties.Num_Bol.Type == 0 {
        return false, "bftx.Properties.Num_Bol.Type is not a number."
    }
    if (reflect.TypeOf(bftx.Properties.Master_Info.Type) != reflect.TypeOf("s")){
        return false, "bftx.Properties.Master_Info.Type is not a string."
    }
    if (reflect.TypeOf(bftx.Properties.Master_Info.Properties.First_Name.Type) != reflect.TypeOf("s")){
        return false, "bftx.Properties.Master_Info.Properties.First_Name.Type is not a string."
    }
    if (reflect.TypeOf(bftx.Properties.Master_Info.Properties.Last_Name.Type) != reflect.TypeOf("s")){
        return false, "bftx.Properties.Master_Info.Properties.Last_Name.Type is not a string."
    }
    if (reflect.TypeOf(bftx.Properties.Master_Info.Properties.Sig.Type) != reflect.TypeOf("s")){
        return false, "bftx.Properties.Master_Info.Properties.Sig.Type is not a string."
    }
    if (reflect.TypeOf(bftx.Properties.Agent_for_Master.Type) != reflect.TypeOf("s")){
        return false, "bftx.Properties.Agent_for_Master.Type is not a string."
    }
    if (reflect.TypeOf(bftx.Properties.Agent_for_Master.Properties.First_Name.Type) != reflect.TypeOf("s")){
        return false, "bftx.Properties.Agent_for_Master.Properties.First_Name.Type is not a string."
    }
    if (reflect.TypeOf(bftx.Properties.Agent_for_Master.Properties.Last_Name.Type) != reflect.TypeOf("s")){
        return false, "bftx.Properties.Agent_for_Master.Properties.Last_Name.Type is not a string."
    }
    if (reflect.TypeOf(bftx.Properties.Agent_for_Master.Properties.Sig.Type) != reflect.TypeOf("s")){
        return false, "bftx.Properties.Agent_for_Master.Properties.Sig.Type is not a string."
    }
    if (reflect.TypeOf(bftx.Properties.Agent_for_Owner.Type) != reflect.TypeOf("s")){
        return false, "bftx.Properties.Agent_for_Owner.Type is not a string."
    }
    if (reflect.TypeOf(bftx.Properties.Agent_for_Owner.Properties.First_Name.Type) != reflect.TypeOf("s")){
        return false, "bftx.Properties.Agent_for_Owner.Properties.First_Name.Type is not a string."
    }
    if (reflect.TypeOf(bftx.Properties.Agent_for_Owner.Properties.Last_Name.Type) != reflect.TypeOf("s")){
        return false, "bftx.Properties.Agent_for_Owner.Properties.Last_Name.Type is not a string."
    }
    if (reflect.TypeOf(bftx.Properties.Agent_for_Owner.Properties.Sig.Type) != reflect.TypeOf("s")){
        return false, "bftx.Properties.Agent_for_Owner.Properties.Sig.Type is not a string."
    }
    if (reflect.TypeOf(bftx.Properties.Agent_for_Owner.Properties.Conditions_for_Carriage.Type) != reflect.TypeOf("s")){
        return false, "bftx.Properties.Agent_for_Owner.Properties.Conditions_for_Carriage.Type is not a string."
    }
    return true, ""
}