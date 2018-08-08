package graphqlObj

import "github.com/graphql-go/graphql"

// ExtensionInput object for GraphQL integration
var ExtensionInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "ExtensionInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"ServiceLevel": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},
	},
)

// Extension object for GraphQL integration
var Extension = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Extension",
		Fields: graphql.Fields{
			"ServiceLevel": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
