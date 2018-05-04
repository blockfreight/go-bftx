package graphqlObj

import "github.com/graphql-go/graphql"

// IssueDetailsInput object for GraphQL integration
var IssueDetailsInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "IssueDetails",
		Fields: graphql.InputObjectConfigFieldMap{
			"PlaceOfIssue": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"DateOfIssue": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},
	},
)

// IssueDetails object for GraphQL integration
var IssueDetails = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "IssueDetails",
		Fields: graphql.Fields{
			"PlaceOfIssue": &graphql.Field{
				Type: graphql.String,
			},
			"DateOfIssue": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
