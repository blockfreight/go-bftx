package graphqlObj

import "github.com/graphql-go/graphql"

// PropertiesType object for GraphQL integration
var PropertiesType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Properties",
		Fields: graphql.Fields{
			"Shipper": &graphql.Field{
				Type: graphql.String,
			},
			"EncryptionMetaData": &graphql.Field{
				Type: graphql.String,
			},
			"BolNum": &graphql.Field{
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
			"HouseBill": &graphql.Field{
				Type: graphql.String,
			},
			"Packages": &graphql.Field{
				Type: graphql.String,
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
				Type: graphql.String,
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
				Type: IssueDetails,
			},
			"NumBol": &graphql.Field{
				Type: graphql.String,
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

// PropertiesInput object for GraphQL integration
var PropertiesInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "Properties",
		Fields: graphql.InputObjectConfigFieldMap{
			"Shipper": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"EncryptionMetaData": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"BolNum": &graphql.InputObjectFieldConfig{
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
			"HouseBill": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"Packages": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
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
				Type: graphql.String,
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
				Type: IssueDetailsInput,
			},
			"NumBol": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
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
