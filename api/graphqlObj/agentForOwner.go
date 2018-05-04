package graphqlObj

import "github.com/graphql-go/graphql"

// AgentForOwnerInput object for GraphQL integration
var AgentForOwnerInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "AgentForOwner",
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
			"ConditionsForCarriage": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},
	},
)

// AgentForOwner object for GraphQL integration
var AgentForOwner = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "agentForOwner",
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
			"ConditionsForCarriage": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
