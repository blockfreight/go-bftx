package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"encoding/json"
	"strconv"
)

type JSONObject struct {
    Type string
    Properties PropertiesType
}

type PropertiesType struct {
    Shipper ShipperType
    Bol_Num BolNumType
    Ref_Num RefNumType
    Consignee ConsigneeType
    Vessel VesselType
    Port_Of_Loading PortOfLoadingType
    Port_Of_Discharge PortOfDischargeType
    Nofity_Address NotifyAddressType		/*Nofity o notify?*/
    Desc_Of_Goods DescOfGoodsType
    Gross_Weight GrossWeightType
    Freight_Payable_Amt FreightPayableAmtType
    Freight_Adv_Amt FreightAdvAmtType
    General_Instructions GeneralInstructionsType
    Date_Shipped DateShippedType
    Issue_Details IssueDetailsType
    Num_Bol NumBolType						/*Is it the same Bol_Num?*/
    Master_Info MasterInfoType
    Agent_For_Master AgentForMasterType
    Agent_For_Owner AgentForOwnerType
}

type ShipperType struct {
    Type string
}

type BolNumType struct {
    Type int
}

type RefNumType struct {
    Type int
}

type ConsigneeType struct {
    Type string								/*Null???*/
}

type VesselType struct {
    Type int
}

type PortOfLoadingType struct {
    Type int
}

type PortOfDischargeType struct {
    Type int
}

type NotifyAddressType struct {
    Type string
}

type DescOfGoodsType struct {
    Type string
}

type GrossWeightType struct {
    Type int								/*Should it be float?*/
}

type FreightPayableAmtType struct {
    Type int
}

type FreightAdvAmtType struct {
    Type int
}

type GeneralInstructionsType struct {
    Type string
}

type DateShippedType struct {
    Type int
    Format string
}

type IssueDetailsType struct {
	Type string
	Issue_Details_Properties IssueDetailsPropertiesType
}

type IssueDetailsPropertiesType struct {
	Place_Of_Issue PlaceOfIssueType
	Date_Of_Issue DateOfIssueType
}

type PlaceOfIssueType struct {
	Type string
}

type DateOfIssueType struct {
	Type int
	Format string
}

type NumBolType struct {
	Type int
}

type MasterInfoType struct {
	Type string
	Master_Info_Properties MasterInfoPropertiesType
}

type MasterInfoPropertiesType struct {
	First_Name FirstNameType
	Last_Name LastNameType
	Sig SigType
}

type AgentForMasterType struct {
	Type string
	Agent_For_Master_Properties AgentForMasterPropertiesType
}

type AgentForMasterPropertiesType struct {
	First_Name FirstNameType
	Last_Name LastNameType
	Sig SigType
}

type AgentForOwnerType struct {
	Type string
	Agent_For_Owner_Properties AgentForOwnerPropertiesType
}

type AgentForOwnerPropertiesType struct {
	First_Name FirstNameType
	Last_Name LastNameType
	Sig SigType
	Conditions_For_carriage string
}

type FirstNameType struct {
	Type string
}

type LastNameType struct {
	Type string
}

type SigType struct {
	Type string
}

func readJSON(path string) []byte {
	fmt.Println("\nReading "+path+"\n")
	file, e := ioutil.ReadFile(path)
	if e != nil {
		fmt.Println("File error: %v\n", e)
		os.Exit(1)
	}
	//fmt.Println(string(file))
	return file
}

func compileJSON(file []byte) string {
	return strings.Replace(strings.Replace(string(file), "\n", "", -1), " ", "", -1)
}

func between(value string, a string, b string) string {
    // Get substring between two strings.
    posFirst := strings.Index(value, a)
    if posFirst == -1 {
		return ""
    }
    posLast := strings.Index(value, b)
    if posLast == -1 {
		return ""
    }
    posFirstAdjusted := posFirst + len(a)
    if posFirstAdjusted >= posLast {
		return ""
    }
    	return value[posFirstAdjusted:posLast]
}

func getStructure() []string{
	return []string{
		"{\"type\":\"",
		"\",\"properties\":{\"shipper\":{\"type\":\"",
		"\"},\"BOL_NUM\":{\"type\":",
		"},\"REF_NUM\":{\"type\":",
		"},\"consignee\":{\"type\":",
		"},\"vessel\":{\"type\":",
		"},\"port_of_loading\":{\"type\":",
		"},\"port_of_discharge\":{\"type\":",
		"},\"nofity_address\":{\"type\":\"",
		"\"},\"desc_of_goods\":{\"type\":\"",
		"\"},\"gross_weight\":{\"type\":",
		"},\"freight_payable_amt\":{\"type\":",
		"},\"freight_adv_amt\":{\"type\":",
		"},\"general_instructions\":{\"type\":\"",
		"\"},\"date_shipped\":{\"type\":",
		",\"format\":\"",
		"\"},\"issue_details\":{\"type\":\"",
		"\",\"issue_details_properties\":{\"place_of_issue\":{\"type\":\"",
		"\"},\"date_of_issue\":{\"type\":",
		",\"format\":\"",				//Check! (Duplicated)
		"\"}}},\"num_BOL\":{\"type\":",
		"},\"master_info\":{\"type\":\"",
		"\",\"master_info_properties\":{\"first_name\":{\"type\":\"",
		"\"},\"last_name\":{\"type\":\"",
		"\"},\"sig\":{\"type\":",
		"}}},\"agent_for_master\":{\"type\":\"",
		"\",\"agent_for_master_properties\":{\"first_name\":{\"type\":\"",
		"\"},\"last_name\":{\"type\":\"",
		"\"},\"sig\":{\"type\":",		//Check! (Duplicated)
		"}}},\"agent_for_owner\":{\"type\":\"",
		"\",\"agent_for_owner_properties\":{\"first_name\":{\"type\":\"",
		"\"},\"last_name\":{\"type\":\"",	//Check! (Duplicated)
		"\"},\"sig\":{\"type\":",		//Check! (Duplicated)
		"},\"conditions_for_carriage\":{\"type\":\"",
		"\"}}}}}",
	}
}

func compareJSON(schemeContent string, exampleContent string) string {
	match := false
	//i := -1
	fmt.Printf("\nValidating...\n")
	
	structure := getStructure();
	fmt.Println("Structure Len: "+strconv.Itoa(len(structure)))
	//for i, seg := range structure {
	for i := 0; i < len(structure)-1; i++ {
		//fmt.Println("-: "+seg)
    	//fmt.Println(between(exampleContent, seg, structure[i+1]))
    	fmt.Println(between(exampleContent, structure[i], structure[i+1]))
	}

	/*seg_1 := "{\"type\":\""
	i = strings.Index(exampleContent,seg_1)
	if i != -1 {
		match = true
	}
	fmt.Println("Index: "+strconv.Itoa(i), "Value: "+string(exampleContent[i+9:i+15]))
	
	seg_2 := "\",\"properties\":{\"shipper\":{\"type\":\""
	i = strings.Index(exampleContent,seg_2)
	if i != -1 {
		match = true
	}
	fmt.Println("Index: "+strconv.Itoa(i), "Value: "+string(exampleContent[i+35:i+45]))

	seg_3 := "\"},\"BOL_NUM\":{\"type\":"
	i = strings.Index(exampleContent,seg_3)
	if i != -1 {
		match = true
	}
	fmt.Println("Index: "+strconv.Itoa(i), "Value: "+string(exampleContent[i+21:i+26]))
	
	seg_4 := "},\"REF_NUM\":{\"type\":"
	i = strings.Index(exampleContent,seg_4)
	if i != -1 {
		match = true
	}
	fmt.Println("Index: "+strconv.Itoa(i), "Value: "+string(exampleContent[i+20:i+29]))

    // Test between func.
    fmt.Println(getStructure())*/
	
	//fmt.Println(strings.Index(compactedScheme,"{\"type\":\""))
	/*splittedScheme := strings.Split(compactedScheme, ",")
	splittedExample := strings.Split(compactedExample, ",")
	fmt.Printf("\nScheme:\n\n")
	for _, value := range splittedScheme {
		fmt.Println(value)
	}
	fmt.Printf("\nExample:\n\n")
	for _, value := range splittedExample {
		fmt.Println(value)
	}*/
	if match {
		return "The JSON accomplished the required structure."
	} else {
		return "The JSON did not accomplish the required structure."
	}
}

func printStructure(file JSONObject) {
	//fmt.Printf("Results: %+v\n", file)
	fmt.Println("Type: "+file.Type)
	fmt.Println("Properties:")
	fmt.Println("  Shipper:")
	fmt.Println("    Type: "+file.Properties.Shipper.Type)
	fmt.Println("  Bol_Num:")
	fmt.Println("    Type: "+strconv.Itoa(file.Properties.Bol_Num.Type))
	fmt.Println("  Ref_Num:")
	fmt.Println("    Type: "+strconv.Itoa(file.Properties.Ref_Num.Type))
	fmt.Println("  Consignee:")
	fmt.Println("    Type: "+file.Properties.Consignee.Type)
	fmt.Println("  Vessel:")
	fmt.Println("    Type: "+strconv.Itoa(file.Properties.Vessel.Type))
	fmt.Println("  Port of Loading:")
	fmt.Println("    Type: "+strconv.Itoa(file.Properties.Port_Of_Loading.Type))
	fmt.Println("  Port of Discharge:")
	fmt.Println("    Type: "+strconv.Itoa(file.Properties.Port_Of_Discharge.Type))
	fmt.Println("  Notify Address:")
	fmt.Println("    Type: "+file.Properties.Nofity_Address.Type)
	fmt.Println("  Desc Of Goods:")
	fmt.Println("    Type: "+file.Properties.Desc_Of_Goods.Type)
	fmt.Println("  Gross Weight:")
	fmt.Println("    Type: "+strconv.Itoa(file.Properties.Gross_Weight.Type))
	fmt.Println("  Freight Payable AMT:")
	fmt.Println("    Type: "+strconv.Itoa(file.Properties.Freight_Payable_Amt.Type))
	fmt.Println("  Freight ADV AMT:")
	fmt.Println("    Type: "+strconv.Itoa(file.Properties.Freight_Adv_Amt.Type))
	fmt.Println("  General Instructions:")
	fmt.Println("    Type: "+file.Properties.General_Instructions.Type)
	fmt.Println("  Date Shipped:")
	fmt.Println("    Type: "+strconv.Itoa(file.Properties.Date_Shipped.Type))
	fmt.Println("    Format: "+file.Properties.Date_Shipped.Format)
	fmt.Println("  Issue Details:")
	fmt.Println("    Type: "+file.Properties.Issue_Details.Type)
	fmt.Println("  Num_Bol:")
	fmt.Println("    Type: "+strconv.Itoa(file.Properties.Num_Bol.Type))
	fmt.Println("  Master_Info:")
	fmt.Println("    Type: "+file.Properties.Master_Info.Type)
	fmt.Println("      Properties:")
	fmt.Println("        First Name:")
	fmt.Println("          Type: "+file.Properties.Master_Info.Master_Info_Properties.First_Name.Type)
	fmt.Println("        Last Name:")
	fmt.Println("          Type: "+file.Properties.Master_Info.Master_Info_Properties.Last_Name.Type)
	fmt.Println("        Sig:")
	fmt.Println("          Type: "+file.Properties.Master_Info.Master_Info_Properties.Sig.Type)
	fmt.Println("  Agent_for_Master:")
	fmt.Println("    Type: "+file.Properties.Agent_For_Master.Type)
	fmt.Println("      Properties:")
	fmt.Println("        First Name:")
	fmt.Println("          Type: "+file.Properties.Agent_For_Master.Agent_For_Master_Properties.First_Name.Type)
	fmt.Println("        Last Name:")
	fmt.Println("          Type: "+file.Properties.Agent_For_Master.Agent_For_Master_Properties.Last_Name.Type)
	fmt.Println("        Sig:")
	fmt.Println("          Type: "+file.Properties.Agent_For_Master.Agent_For_Master_Properties.Sig.Type)
	fmt.Println("  Agent_for_Owner:")
	fmt.Println("    Type: "+file.Properties.Agent_For_Owner.Type)
	fmt.Println("      Properties:")
	fmt.Println("        First Name:")
	fmt.Println("          Type: "+file.Properties.Agent_For_Owner.Agent_For_Owner_Properties.First_Name.Type)
	fmt.Println("        Last Name:")
	fmt.Println("          Type: "+file.Properties.Agent_For_Owner.Agent_For_Owner_Properties.Last_Name.Type)
	fmt.Println("        Sig:")
	fmt.Println("          Type: "+file.Properties.Agent_For_Owner.Agent_For_Owner_Properties.Sig.Type)
}

func main() {
	fmt.Println("\nHello Blockfreight™ world!")
	
	//Set Working directory
	wd,_ := os.Getwd()

	//Set paths and JSON of scheme and example
	schemePath := wd+"/go_lang/src/github.com/julian-nunezm/blockfreight/blockfreight_app/files/bf_tx_schema_pub_var_rfc2.json"
	schemeJSON := readJSON(schemePath)
	compactedScheme := compileJSON(schemeJSON)
	fmt.Println(compactedScheme)
	examplePath := wd+"/go_lang/src/github.com/julian-nunezm/blockfreight/blockfreight_app/files/bf_tx_example.json"
	exampleJSON := readJSON(examplePath)
	compactedExample := compileJSON(exampleJSON)
	fmt.Println(compactedExample)

	//Validate structure
	fmt.Println(compareJSON(compactedScheme, compactedExample)+"\n")

	//Save JSON data
	var jsontype JSONObject
    json.Unmarshal(exampleJSON, &jsontype)

    //Print JSON structure
    //printStructure(jsontype)
}