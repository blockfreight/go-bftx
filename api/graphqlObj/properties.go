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
		},
	},
)
