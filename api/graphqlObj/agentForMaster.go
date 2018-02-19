package graphqlObj

import "github.com/graphql-go/graphql"

// AgentForMasterInput object for GraphQL integration
var AgentForMasterInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "AgentForMaster",
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

// AgentForMaster object for GraphQL integration
var AgentForMaster = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "AgentForMaster",
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
