package graphqlObj

import (
	"github.com/graphql-go/graphql"
)

var InfoType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Info",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
