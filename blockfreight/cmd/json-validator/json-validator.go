package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "encoding/json"
    "reflect"

    "github.com/blockfreight/blockfreight-alpha/blockfreight/bft/bol"
    "github.com/davecgh/go-spew/spew"
)

func main(){
    printJson := true
    examplePath := "./files/bf_tx_example.json"
    var bol bol.BoL
    json.Unmarshal(readJSON(examplePath), &bol)
    if printJson { spew.Dump(bol) }
    valid, err := ValidateBoL(bol)
    if valid {
        fmt.Println("Success! [OK]")
    } else {
        fmt.Println(`
    Blockfreight, Inc. © 2017. Open Source (MIT) License.

    Error [01]:

    Invalid structure in JSON provided. JSON 结构无效.
    Struttura JSON non valido. هيكل JSON صالح. 無効なJSON構造

    support: support@blockfreight.com`)
    }
    if err != "" {
        fmt.Println(`
    Specific Error [01]:

    `+err)
    }
}

func readJSON(path string) []byte {
    fmt.Println("\nReading "+path+"\n")
    file, e := ioutil.ReadFile(path)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }
    return file
}

func ValidateBoL(bol bol.BoL) (bool, string){
    if (reflect.TypeOf(bol.Type) != reflect.TypeOf("s")){
        return false, "bol.Type is not a string."
    }
    if (reflect.TypeOf(bol.Properties.Shipper.Type) != reflect.TypeOf("s")){
        return false, "bol.Properties.Shipper.Type is not a string."
    }
    if (reflect.TypeOf(bol.Properties.Bol_Num.Type) != reflect.TypeOf(1)) || bol.Properties.Bol_Num.Type == 0 {
        return false, "bol.Properties.Bol_Num.Type is not a nunber."
    }
    if (reflect.TypeOf(bol.Properties.Ref_Num.Type) != reflect.TypeOf(1)) || bol.Properties.Ref_Num.Type == 0 {
        return false, "bol.Properties.Ref_Num.Type is not a number."
    }
    if (reflect.TypeOf(bol.Properties.Consignee.Type) != reflect.TypeOf("s")){
        return false, "bol.Properties.Consignee.Type is not a string."
    }
    if (reflect.TypeOf(bol.Properties.Vessel.Type) != reflect.TypeOf(1)) || bol.Properties.Vessel.Type == 0 {
        return false, "bol.Properties.Vessel.Type is not a number."
    }
    if (reflect.TypeOf(bol.Properties.Port_of_Loading.Type) != reflect.TypeOf(1)) || bol.Properties.Port_of_Loading.Type == 0 {
        return false, "bol.Properties.Port_of_Loading.Type is not a number."
    }
    if (reflect.TypeOf(bol.Properties.Port_of_Discharge.Type) != reflect.TypeOf(1)) || bol.Properties.Port_of_Discharge.Type == 0 {
        return false, "bol.Properties.Port_of_Discharge.Type is not a number."
    }
    if (reflect.TypeOf(bol.Properties.Notify_Address.Type) != reflect.TypeOf("s")){
        return false, "bol.Properties.Notify_Address.Type is not a string."
    }
    if (reflect.TypeOf(bol.Properties.Desc_of_Goods.Type) != reflect.TypeOf("s")){
        return false, "bol.Properties.Desc_of_Goods.Type is not a string."
    }
    if (reflect.TypeOf(bol.Properties.Gross_Weight.Type) != reflect.TypeOf(1)) || bol.Properties.Gross_Weight.Type == 0 {
        return false, "bol.Properties.Gross_Weight.Type is not a number."
    }
    if (reflect.TypeOf(bol.Properties.Freight_Payable_Amt.Type) != reflect.TypeOf(1)) || bol.Properties.Freight_Payable_Amt.Type == 0 {
        return false, "bol.Properties.Freight_Payable_Amt.Type is not a number."
    }
    if (reflect.TypeOf(bol.Properties.Freight_Adv_Amt.Type) != reflect.TypeOf(1)) || bol.Properties.Freight_Adv_Amt.Type == 0 {
        return false, "bol.Properties.Freight_Adv_Amt.Type is not a number."
    }
    if (reflect.TypeOf(bol.Properties.General_Instructions.Type) != reflect.TypeOf("s")){
        return false, "bol.Properties.General_Instructions.Type is not a string."
    }
    if (reflect.TypeOf(bol.Properties.Date_Shipped.Type) != reflect.TypeOf(1)) || bol.Properties.Date_Shipped.Type == 0 {
        return false, "bol.Properties.Date_Shipped.Type is not a number."
    }
    if (reflect.TypeOf(bol.Properties.Date_Shipped.Format) != reflect.TypeOf("s")){
        return false, "bol.Properties.Date_Shipped.Format is not a date format."
    }
    if (reflect.TypeOf(bol.Properties.Issue_Details.Type) != reflect.TypeOf("s")){
        return false, "bol.Properties.Issue_Details.Type is not a string."
    }
    if (reflect.TypeOf(bol.Properties.Issue_Details.Properties.Place_of_Issue.Type) != reflect.TypeOf("s")){
        return false, "bol.Properties.Issue_Details.Properties.Place_of_Issue.Type is not a string."
    }
    if (reflect.TypeOf(bol.Properties.Issue_Details.Properties.Date_of_Issue.Type) != reflect.TypeOf(1)) || bol.Properties.Issue_Details.Properties.Date_of_Issue.Type == 0 {
        return false, "bol.Properties.Issue_Details.Properties.Date_of_Issue.Type is not a number."
    }
    if (reflect.TypeOf(bol.Properties.Issue_Details.Properties.Date_of_Issue.Format) != reflect.TypeOf("s")){
        return false, "bol.Properties.Issue_Details.Properties.Place_of_Issue.Format is not a date format."
    }
    if (reflect.TypeOf(bol.Properties.Num_Bol.Type) != reflect.TypeOf(1)) || bol.Properties.Num_Bol.Type == 0 {
        return false, "bol.Properties.Num_Bol.Type is not a number."
    }
    if (reflect.TypeOf(bol.Properties.Master_Info.Type) != reflect.TypeOf("s")){
        return false, "bol.Properties.Master_Info.Type is not a string."
    }
    if (reflect.TypeOf(bol.Properties.Master_Info.Properties.First_Name.Type) != reflect.TypeOf("s")){
        return false, "bol.Properties.Master_Info.Properties.First_Name.Type is not a string."
    }
    if (reflect.TypeOf(bol.Properties.Master_Info.Properties.Last_Name.Type) != reflect.TypeOf("s")){
        return false, "bol.Properties.Master_Info.Properties.Last_Name.Type is not a string."
    }
    if (reflect.TypeOf(bol.Properties.Master_Info.Properties.Sig.Type) != reflect.TypeOf("s")){
        return false, "bol.Properties.Master_Info.Properties.Sig.Type is not a string."
    }
    if (reflect.TypeOf(bol.Properties.Agent_for_Master.Type) != reflect.TypeOf("s")){
        return false, "bol.Properties.Agent_for_Master.Type is not a string."
    }
    if (reflect.TypeOf(bol.Properties.Agent_for_Master.Properties.First_Name.Type) != reflect.TypeOf("s")){
        return false, "bol.Properties.Agent_for_Master.Properties.First_Name.Type is not a string."
    }
    if (reflect.TypeOf(bol.Properties.Agent_for_Master.Properties.Last_Name.Type) != reflect.TypeOf("s")){
        return false, "bol.Properties.Agent_for_Master.Properties.Last_Name.Type is not a string."
    }
    if (reflect.TypeOf(bol.Properties.Agent_for_Master.Properties.Sig.Type) != reflect.TypeOf("s")){
        return false, "bol.Properties.Agent_for_Master.Properties.Sig.Type is not a string."
    }
    if (reflect.TypeOf(bol.Properties.Agent_for_Owner.Type) != reflect.TypeOf("s")){
        return false, "bol.Properties.Agent_for_Owner.Type is not a string."
    }
    if (reflect.TypeOf(bol.Properties.Agent_for_Owner.Properties.First_Name.Type) != reflect.TypeOf("s")){
        return false, "bol.Properties.Agent_for_Owner.Properties.First_Name.Type is not a string."
    }
    if (reflect.TypeOf(bol.Properties.Agent_for_Owner.Properties.Last_Name.Type) != reflect.TypeOf("s")){
        return false, "bol.Properties.Agent_for_Owner.Properties.Last_Name.Type is not a string."
    }
    if (reflect.TypeOf(bol.Properties.Agent_for_Owner.Properties.Sig.Type) != reflect.TypeOf("s")){
        return false, "bol.Properties.Agent_for_Owner.Properties.Sig.Type is not a string."
    }
    if (reflect.TypeOf(bol.Properties.Agent_for_Owner.Properties.Conditions_for_Carriage.Type) != reflect.TypeOf("s")){
        return false, "bol.Properties.Agent_for_Owner.Properties.Conditions_for_Carriage.Type is not a string."
    }
    return true, ""
}