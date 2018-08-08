package graphqlObj

import "github.com/graphql-go/graphql"

// ConsolInput object for GraphQL integration
var ConsolInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "Consol",
		Fields: graphql.InputObjectConfigFieldMap{
			"Masterbill": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"ContainerMode": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"PaymentMethod": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"PortOfDischarge": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"PortOfLoading": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"ShipmentType": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"TransportMode": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"VesselName": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"VoyageFlightNo": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"EstimatedTimeOfDeparture": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"EstimatedTimeOfArrival": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"Carrier": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"ContainerNumber": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"ContainerType": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"DeliveryMode": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"Seal": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},
	},
)

// Consol object for GraphQL integration
var Consol = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Consol",
		Fields: graphql.Fields{
			"Masterbill": &graphql.Field{
				Type: graphql.String,
			},
			"ContainerMode": &graphql.Field{
				Type: graphql.String,
			},
			"PaymentMethod": &graphql.Field{
				Type: graphql.String,
			},
			"PortOfDischarge": &graphql.Field{
				Type: graphql.String,
			},
			"PortOfLoading": &graphql.Field{
				Type: graphql.String,
			},
			"ShipmentType": &graphql.Field{
				Type: graphql.String,
			},
			"TransportMode": &graphql.Field{
				Type: graphql.String,
			},
			"VesselName": &graphql.Field{
				Type: graphql.String,
			},
			"VoyageFlightNo": &graphql.Field{
				Type: graphql.String,
			},
			"EstimatedTimeOfDeparture": &graphql.Field{
				Type: graphql.String,
			},
			"EstimatedTimeOfArrival": &graphql.Field{
				Type: graphql.String,
			},
			"Carrier": &graphql.Field{
				Type: graphql.String,
			},
			"ContainerNumber": &graphql.Field{
				Type: graphql.String,
			},
			"ContainerType": &graphql.Field{
				Type: graphql.String,
			},
			"DeliveryMode": &graphql.Field{
				Type: graphql.String,
			},
			"Seal": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
