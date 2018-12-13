package graphqlObj

import "github.com/graphql-go/graphql"

// PropertiesType object for GraphQL integration
var PropertiesType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Properties",
		Fields: graphql.Fields{
			"Consol": &graphql.Field{
				Type: Consol,
			},
			"Shipment": &graphql.Field{
				Type: Shipment,
			},
			"Extension": &graphql.Field{
				Type: Extension,
			},
			"Shipper": &graphql.Field{
				Type: graphql.String,
			},
			"BolNum": &graphql.Field{
				Type: graphql.String,
			},
			"NumBol": &graphql.Field{
				Type: graphql.String,
			},
			"RefNum": &graphql.Field{
				Type: graphql.String,
			},
			"Consignee": &graphql.Field{
				Type: graphql.String,
			},
			"Vessel": &graphql.Field{
				Type: graphql.String,
			},
			"PortOfLoading": &graphql.Field{
				Type: graphql.String,
			},
			"PortOfDischarge": &graphql.Field{
				Type: graphql.String,
			},
			"NotifyAddress": &graphql.Field{
				Type: graphql.String,
			},
			"DescOfGoods": &graphql.Field{
				Type: graphql.String,
			},
			"GrossWeight": &graphql.Field{
				Type: graphql.String,
			},
			"FreightPayableAmt": &graphql.Field{
				Type: graphql.String,
			},
			"FreightAdvAmt": &graphql.Field{
				Type: graphql.String,
			},
			"GeneralInstructions": &graphql.Field{
				Type: graphql.String,
			},
			"DateShipped": &graphql.Field{
				Type: graphql.String,
			},
			"IssueDetails": &graphql.Field{
				Type: graphql.String,
			},
			"MasterInfo": &graphql.Field{
				Type: graphql.String,
			},
			"AgentForMaster": &graphql.Field{
				Type: graphql.String,
			},
			"AgentForOwner": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

// PropertiesInput object for GraphQL integration
var PropertiesInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "Properties",
		Fields: graphql.InputObjectConfigFieldMap{
			"Consol": &graphql.InputObjectFieldConfig{
				Type: ConsolInput,
			},
			"Shipment": &graphql.InputObjectFieldConfig{
				Type: ShipmentInput,
			},
			"Extension": &graphql.InputObjectFieldConfig{
				Type: ExtensionInput,
			},
			"Shipper": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"BolNum": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"NumBol": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"RefNum": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"Consignee": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"Vessel": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"PortOfLoading": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"PortOfDischarge": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"NotifyAddress": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"DescOfGoods": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"GrossWeight": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"FreightPayableAmt": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"FreightAdvAmt": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"GeneralInstructions": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"DateShipped": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"IssueDetails": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"MasterInfo": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"AgentForMaster": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"AgentForOwner": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},
	},
)
