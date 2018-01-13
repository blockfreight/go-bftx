package graphqlObj

import "github.com/graphql-go/graphql"

var PropertiesType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Properties",
		Fields: graphql.Fields{
			"Shipper": &graphql.Field{
				Type: graphql.String,
			},
			"BolNum": &graphql.Field{
				Type: graphql.Int,
			},
			"RefNum": &graphql.Field{
				Type: graphql.Int,
			},
			"Consignee": &graphql.Field{
				Type: graphql.String,
			},
			"Vessel": &graphql.Field{
				Type: graphql.Int,
			},
			"PortOfLoading": &graphql.Field{
				Type: graphql.String,
			},
			"PortOfDischarge": &graphql.Field{
				Type: graphql.String,
			},
			"HouseBill": &graphql.Field{
				Type: graphql.String,
			},
			"Packages": &graphql.Field{
				Type: graphql.Int,
			},
			"PackType": &graphql.Field{
				Type: graphql.String,
			},
			"INCOTerms": &graphql.Field{
				Type: graphql.String,
			},
			"Destination": &graphql.Field{
				Type: graphql.String,
			},
			"MarksAndNumbers": &graphql.Field{
				Type: graphql.String,
			},
			"UnitOfWeight": &graphql.Field{
				Type: graphql.String,
			},
			"DeliverAgent": &graphql.Field{
				Type: graphql.String,
			},
			"ReceiveAgent": &graphql.Field{
				Type: graphql.String,
			},
			"Container": &graphql.Field{
				Type: graphql.String,
			},
			"ContainerSeal": &graphql.Field{
				Type: graphql.String,
			},
			"ContainerMode": &graphql.Field{
				Type: graphql.String,
			},
			"ContainerType": &graphql.Field{
				Type: graphql.String,
			},
			"Volume": &graphql.Field{
				Type: graphql.Float,
			},
			"UnitOfVolume": &graphql.Field{
				Type: graphql.String,
			},
			"NotifyAddress": &graphql.Field{
				Type: graphql.String,
			},
			"DescOfGoods": &graphql.Field{
				Type: graphql.String,
			},
			"GrossWeight": &graphql.Field{
				Type: graphql.Float,
			},
			"FreightPayableAmt": &graphql.Field{
				Type: graphql.Int,
			},
			"FreightAdvAmt": &graphql.Field{
				Type: graphql.Int,
			},
			"GeneralInstructions": &graphql.Field{
				Type: graphql.String,
			},
			"DateShipped": &graphql.Field{
				Type: graphql.String,
			},
			"IssueDetails": &graphql.Field{
				Type: IssueDetails,
			},
			"NumBol": &graphql.Field{
				Type: graphql.Int,
			},
			"MasterInfo": &graphql.Field{
				Type: MasterInfo,
			},
			"AgentForMaster": &graphql.Field{
				Type: AgentForMaster,
			},
			"AgentForOwner": &graphql.Field{
				Type: AgentForOwner,
			},
		},
	},
)

var PropertiesInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "Properties",
		Fields: graphql.InputObjectConfigFieldMap{
			"Shipper": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"BolNum": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			"RefNum": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			"Consignee": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"Vessel": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			"PortOfLoading": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"PortOfDischarge": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"HouseBill": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"Packages": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			"PackType": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"INCOTerms": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"Destination": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"MarksAndNumbers": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"UnitOfWeight": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"DeliverAgent": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"ReceiveAgent": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"Container": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"ContainerSeal": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"ContainerMode": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"ContainerType": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"Volume": &graphql.InputObjectFieldConfig{
				Type: graphql.Float,
			},
			"UnitOfVolume": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"NotifyAddress": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"DescOfGoods": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"GrossWeight": &graphql.InputObjectFieldConfig{
				Type: graphql.Float,
			},
			"FreightPayableAmt": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			"FreightAdvAmt": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			"GeneralInstructions": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"DateShipped": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"IssueDetails": &graphql.InputObjectFieldConfig{
				Type: IssueDetailsInput,
			},
			"NumBol": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			"MasterInfo": &graphql.InputObjectFieldConfig{
				Type: MasterInfoInput,
			},
			"AgentForMaster": &graphql.InputObjectFieldConfig{
				Type: AgentForMasterInput,
			},
			"AgentForOwner": &graphql.InputObjectFieldConfig{
				Type: AgentForOwnerInput,
			},
		},
	},
)
