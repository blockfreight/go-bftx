package bf_tx

import (
    "crypto/ecdsa"
)

// Define Blockfreight™ Transaction (BF_TX) transaction standard

type BF_TX struct {
    Type string
    Properties Properties
    PrivateKey ecdsa.PrivateKey
    Signhash []uint8
    Signed bool
}

type Properties struct {
    Shipper Shipper
    Bol_Num BolNum
    Ref_Num RefNum
    Consignee Consignee
    Vessel Vessel
    Port_of_Loading PortLoading
    Port_of_Discharge PortDischarge
    Notify_Address NotifyAddress
    Desc_of_Goods DescGoods
    Gross_Weight GrossWeight
    Freight_Payable_Amt FreightPayableAmt
    Freight_Adv_Amt FreightAdvAmt
    General_Instructions GeneralInstructions
    Date_Shipped Date
    Issue_Details IssueDetails
    Num_Bol NumBol								// Is it the same Bol_Num?
    Master_Info MasterInfo
    Agent_for_Master AgentMaster
    Agent_for_Owner AgentOwner
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
    Type string									//Null
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
    Type int									//Should it be float?
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
    Type int
    Format string
}

type IssueDetails struct {
	Type string
	Properties IssueDetailsProperties
}

type IssueDetailsProperties struct {
	Place_of_Issue PlaceIssue
	Date_of_Issue Date
}

type PlaceIssue struct {
	Type string
}

type NumBol struct {
	Type int
}

type MasterInfo struct {
	Type string
	Properties MasterInfoProperties
}

type MasterInfoProperties struct {
	First_Name FirstName
	Last_Name LastName
	Sig Sig
}

type AgentMaster struct {
	Type string
	Properties AgentMasterProperties
}

type AgentMasterProperties struct {
	First_Name FirstName
	Last_Name LastName
	Sig Sig
}

type AgentOwner struct {
	Type string
	Properties AgentOwnerProperties
}

type AgentOwnerProperties struct {
	First_Name FirstName
	Last_Name LastName
	Sig Sig
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