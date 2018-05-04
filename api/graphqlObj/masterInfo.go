package graphqlObj

import "github.com/graphql-go/graphql"

// MasterInfoInput object for GraphQL integration
var MasterInfoInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "MasterInfo",
		Fields: graphql.InputObjectConfigFieldMap{
			"FirstName": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"LastName": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"Sig": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},
	},
)

// MasterInfo object for GraphQL integration
var MasterInfo = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "MasterInfo",
		Fields: graphql.Fields{
			"FirstName": &graphql.Field{
				Type: graphql.String,
			},
			"LastName": &graphql.Field{
				Type: graphql.String,
			},
			"Sig": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
