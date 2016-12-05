package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"encoding/json"
	"strconv"
)

//Declaration of structure
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
    Notify_Address NotifyAddressType
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
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	//fmt.Println(string(file))
	return file
}

func compileJSON(file []byte) string {
	return strings.Replace(strings.Replace(string(file), "\n", "", -1), " ", "", -1)
}

func between(value string, a string, b string, n int) string {
    // Get substring between two strings.
    var posLast int
    if n < 1 {
    	n = 1
    }
    for i := 1; i <= n; i++ {
    	posFirst := strings.Index(value, a)
    	if posFirst == -1 {
			return ""
    	}
    	posLast = strings.Index(value, b)
    	if posLast == -1 {
			return ""
    	}
    	posFirstAdjusted := posFirst + len(a)
    	if posFirstAdjusted >= posLast {
			return ""
    	}
    	//return value[posFirstAdjusted:posLast]
    	if n > 1 {
    		value = value[posFirstAdjusted:]
    	} else {
    		value = value[posFirstAdjusted:posLast]
    		//fmt.Println(value, len(value))
    	}
    }
    return value
}

func getStructure() []string {
	wd,_ := os.Getwd()		//Check!
	schemePath := wd+"/go_lang/src/github.com/julian-nunezm/blockfreight/blockfreight-alpha/files/bf_tx_schema_pub_var_rfc2.json"
	schemeJSON := readJSON(schemePath)
	cj := compileJSON(schemeJSON)
	return []string{
		cj[0:9],
		cj[15:50],		
		cj[56:77],		//BOL_NUM
		cj[85:105],		//REF_NUM
		cj[113:135],	//consignee
		cj[141:160],	//vessel
		cj[168:196],	//port_of_loading
		cj[204:234],	//port_of_discharge
		cj[242:270],	//notify_address
		cj[276:304],	//desc_of_goods
		cj[310:336],	//gross_weight
		cj[344:376],	//freight_payable_amt
		cj[384:412],	//freight_adv_amt
		cj[420:454],	//general_instructions
		cj[460:486],	//date_shipped
		cj[494:505],	//format
		cj[514:542],	//issue_details
		cj[548:604],	//place_of_issue
		cj[610:637],	//date_of_issue
		//cj[645:656],	//format 			//Check n-th element
		cj[665:688],	//num_BOL
		cj[696:721],	//master_info
		cj[727:777],	//first_name
		cj[783:807],	//last_name
		cj[813:830],	//sig
		cj[836:868],	//agent_for_master
		cj[874:929],	//first_name
		//cj[935:959],	//last_name 		//Check n-th element
		//cj[965:982],	//sig 				//Check n-th element
		cj[988:1019],	//agent_for_owner
		cj[1025:1079],	//first_name
		//cj[1085:1109],	//last_name 	//Check n-th element
		//cj[1115:1132],	//sig 			//Check n-th element
		cj[1138:1175],	//conditions_for_carriage
		cj[1181:1187],
	}
}

func compareJSON(schemeContent string, exampleContent string) string {
	//match := true
	fmt.Printf("\nValidating...\n")
	structure := getStructure();
	times := 1
	//fmt.Println("Structure Len: "+strconv.Itoa(len(structure)))
	for i := 0; i < len(structure)-1; i++ {
		/*if i == 18 {
			times = 2
			fmt.Println(structure[i], structure[i+1])
		}*/
		if between(exampleContent, structure[i], structure[i+1], times) == "" {
			//match = false
			return "The JSON did not accomplish the required structure."
		}
	}
	return "The JSON accomplished the required structure."
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
	fmt.Println("    Type: "+file.Properties.Notify_Address.Type)
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
	fmt.Println("    Properties:")
	fmt.Println("      Place of Issue:")
	fmt.Println("        Type: "+file.Properties.Issue_Details.Issue_Details_Properties.Place_Of_Issue.Type)
	fmt.Println("      Date of Issue:")
	fmt.Println("        Type: "+strconv.Itoa(file.Properties.Issue_Details.Issue_Details_Properties.Date_Of_Issue.Type))
	fmt.Println("        Format: "+file.Properties.Issue_Details.Issue_Details_Properties.Date_Of_Issue.Format)
	fmt.Println("  Num_Bol:")
	fmt.Println("    Type: "+strconv.Itoa(file.Properties.Num_Bol.Type))
	fmt.Println("  Master_Info:")
	fmt.Println("    Type: "+file.Properties.Master_Info.Type)
	fmt.Println("    Properties:")
	fmt.Println("      First Name:")
	fmt.Println("        Type: "+file.Properties.Master_Info.Master_Info_Properties.First_Name.Type)
	fmt.Println("      Last Name:")
	fmt.Println("        Type: "+file.Properties.Master_Info.Master_Info_Properties.Last_Name.Type)
	fmt.Println("      Sig:")
	fmt.Println("        Type: "+file.Properties.Master_Info.Master_Info_Properties.Sig.Type)
	fmt.Println("  Agent_for_Master:")
	fmt.Println("    Type: "+file.Properties.Agent_For_Master.Type)
	fmt.Println("    Properties:")
	fmt.Println("      First Name:")
	fmt.Println("        Type: "+file.Properties.Agent_For_Master.Agent_For_Master_Properties.First_Name.Type)
	fmt.Println("      Last Name:")
	fmt.Println("        Type: "+file.Properties.Agent_For_Master.Agent_For_Master_Properties.Last_Name.Type)
	fmt.Println("      Sig:")
	fmt.Println("        Type: "+file.Properties.Agent_For_Master.Agent_For_Master_Properties.Sig.Type)
	fmt.Println("  Agent_for_Owner:")
	fmt.Println("    Type: "+file.Properties.Agent_For_Owner.Type)
	fmt.Println("    Properties:")
	fmt.Println("      First Name:")
	fmt.Println("        Type: "+file.Properties.Agent_For_Owner.Agent_For_Owner_Properties.First_Name.Type)
	fmt.Println("      Last Name:")
	fmt.Println("        Type: "+file.Properties.Agent_For_Owner.Agent_For_Owner_Properties.Last_Name.Type)
	fmt.Println("      Sig:")
	fmt.Println("        Type: "+file.Properties.Agent_For_Owner.Agent_For_Owner_Properties.Sig.Type)
}

func main() {
	fmt.Println("\nHello Blockfreight™ world!")
	
	//Set Working directory
	wd,_ := os.Getwd()

	//Set paths and JSON of scheme and example
	schemePath := wd+"/go_lang/src/github.com/julian-nunezm/blockfreight/blockfreight-alpha/files/bf_tx_schema_pub_var_rfc2.json"
	schemeJSON := readJSON(schemePath)
	compactedScheme := compileJSON(schemeJSON)
	//fmt.Println(compactedScheme)
	examplePath := wd+"/go_lang/src/github.com/julian-nunezm/blockfreight/blockfreight-alpha/files/bf_tx_example.json"
	exampleJSON := readJSON(examplePath)
	compactedExample := compileJSON(exampleJSON)
	//fmt.Println(compactedExample)

	//Validate structure
	fmt.Println(compareJSON(compactedScheme, compactedExample)+"\n")

	//Save JSON data
	var jsontype JSONObject
    json.Unmarshal(exampleJSON, &jsontype)

    //Print JSON structure
    //printStructure(jsontype)
}