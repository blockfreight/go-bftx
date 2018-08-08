package graphqlObj

import "github.com/graphql-go/graphql"

// ShipmentInput object for GraphQL integration
var ShipmentInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "MasterInfo",
		Fields: graphql.InputObjectConfigFieldMap{
			"Housebill": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"ContainerMode": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"GoodsDescription": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"MarksAndNumbers": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"HBLAWBChargesDisplay": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"PackQuantity": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"PackType": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"Weight": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"Volume": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"ShippedOnBoard": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"TransportMode": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"EstimatedTimeOfDeparture": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"EstimatedTimeOfArrival": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"Consignee": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"Consignor": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"PackingLineCommodity": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"ContainerNumber": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"INCOTERM": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"ShipmentType": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"ReleaseType": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},
	},
)

// MasterInfo object for GraphQL integration
var Shipment = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "MasterInfo",
		Fields: graphql.Fields{
			"Housebill": &graphql.Field{
				Type: graphql.String,
			},
			"ContainerMode": &graphql.Field{
				Type: graphql.String,
			},
			"GoodsDescription": &graphql.Field{
				Type: graphql.String,
			},
			"MarksAndNumbers": &graphql.Field{
				Type: graphql.String,
			},
			"HBLAWBChargesDisplay": &graphql.Field{
				Type: graphql.String,
			},
			"PackQuantity": &graphql.Field{
				Type: graphql.String,
			},
			"PackType": &graphql.Field{
				Type: graphql.String,
			},
			"Weight": &graphql.Field{
				Type: graphql.String,
			},
			"Volume": &graphql.Field{
				Type: graphql.String,
			},
			"ShippedOnBoard": &graphql.Field{
				Type: graphql.String,
			},
			"TransportMode": &graphql.Field{
				Type: graphql.String,
			},
			"EstimatedTimeOfDeparture": &graphql.Field{
				Type: graphql.String,
			},
			"EstimatedTimeOfArrival": &graphql.Field{
				Type: graphql.String,
			},
			"Consignee": &graphql.Field{
				Type: graphql.String,
			},
			"Consignor": &graphql.Field{
				Type: graphql.String,
			},
			"PackingLineCommodity": &graphql.Field{
				Type: graphql.String,
			},
			"ContainerNumber": &graphql.Field{
				Type: graphql.String,
			},
			"INCOTERM": &graphql.Field{
				Type: graphql.String,
			},
			"ShipmentType": &graphql.Field{
				Type: graphql.String,
			},
			"ReleaseType": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
