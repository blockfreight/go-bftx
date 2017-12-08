package api

import (
	"encoding/json"
	"fmt"
	"net/http" // Provides HTTP client and server implementations.

	"github.com/blockfreight/go-bftx/lib/pkg/leveldb" // Provides some useful functions to work with LevelDB.
	"github.com/graphql-go/graphql"
)

func Start() error {
	http.HandleFunc("/graphql", graphRoute)
	fmt.Println("Now server is running on port 12345")
	fmt.Println("Test with Get      : curl -g 'http://localhost:12345/graphql?query={transaction(id:<BFTX-ID>){Id}}'")
	return http.ListenAndServe(":12345", nil)

}

var issueDetailsProperties = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "IssueDetailsProperties",
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

var issueDetails = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "IssueDetails",
		Fields: graphql.Fields{
			"Type": &graphql.Field{
				Type: graphql.String,
			},
			"Properties": &graphql.Field{
				Type: issueDetailsProperties,
			},
		},
	},
)

var masterInfoProperties = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "MasterInfoProperties",
		Fields: graphql.Fields{
			"FirstName": &graphql.Field{
				Type: graphql.String,
			},
			"LastName": &graphql.Field{
				Type: graphql.String,
			},
			"sig": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var masterInfo = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "MasterInfo",
		Fields: graphql.Fields{
			"Type": &graphql.Field{
				Type: graphql.String,
			},
			"Properties": &graphql.Field{
				Type: masterInfoProperties,
			},
		},
	},
)

var agentForMasterProperties = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "AgentForMasterProperties",
		Fields: graphql.Fields{
			"FirstName": &graphql.Field{
				Type: graphql.String,
			},
			"LastName": &graphql.Field{
				Type: graphql.String,
			},
			"sig": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var agentForMaster = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "AgentForMaster",
		Fields: graphql.Fields{
			"Type": &graphql.Field{
				Type: graphql.String,
			},
			"Properties": &graphql.Field{
				Type: agentForMasterProperties,
			},
		},
	},
)

var agentForOwnerProperties = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "AgentForOwnerProperties",
		Fields: graphql.Fields{
			"FirstName": &graphql.Field{
				Type: graphql.String,
			},
			"LastName": &graphql.Field{
				Type: graphql.String,
			},
			"sig": &graphql.Field{
				Type: graphql.String,
			},
			"ConditionsForCarriage": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var agentForOwner = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "AgentForOwner",
		Fields: graphql.Fields{
			"Type": &graphql.Field{
				Type: graphql.String,
			},
			"Properties": &graphql.Field{
				Type: agentForOwnerProperties,
			},
		},
	},
)

var propertiesType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Properties",
		Fields: graphql.Fields{
			"Shipper": &graphql.Field{
				Type: graphql.String,
			},
			"BolNum": &graphql.Field{
				Type: graphql.Int,
			},
			"RefNum": &graphql.Field{
				Type: graphql.Int,
			},
			"Consignee": &graphql.Field{
				Type: graphql.String,
			},
			"Vessel": &graphql.Field{
				Type: graphql.Int,
			},
			"PortOfLoading": &graphql.Field{
				Type: graphql.Int,
			},
			"PortOfDischarge": &graphql.Field{
				Type: graphql.Int,
			},
			"NotifyAddress": &graphql.Field{
				Type: graphql.String,
			},
			"DescOfGoods": &graphql.Field{
				Type: graphql.String,
			},
			"GrossWeight": &graphql.Field{
				Type: graphql.Int,
			},
			"FreightPayableAmt": &graphql.Field{
				Type: graphql.Int,
			},
			"FreightAdvAmt": &graphql.Field{
				Type: graphql.Int,
			},
			"GeneralInstructions": &graphql.Field{
				Type: graphql.String,
			},
			"DateShipped": &graphql.Field{
				Type: graphql.Int,
			},
			"IssueDetails": &graphql.Field{
				Type: issueDetails,
			},
			"NumBol": &graphql.Field{
				Type: graphql.Int,
			},
			"MasterInfo": &graphql.Field{
				Type: masterInfo,
			},
			"AgentForMaster": &graphql.Field{
				Type: agentForMaster,
			},
			"AgentForOwner": &graphql.Field{
				Type: agentForOwner,
			},
		},
	},
)

var transactionType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Transaction",
		Fields: graphql.Fields{
			"Id": &graphql.Field{
				Type: graphql.String,
			},
			"Type": &graphql.Field{
				Type: graphql.String,
			},
			"Verified": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Transmitted": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Properties": &graphql.Field{
				Type: propertiesType,
			},
		},
	},
)

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: queryType,
	},
)

func graphRoute(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")

	result := executeQuery(query, schema)
	json.NewEncoder(w).Encode(result)
}

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"transaction": &graphql.Field{
				Type: transactionType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					bftxID, isOK := p.Args["id"].(string)
					if !isOK {
						return nil, nil
					}

					bftx, err := leveldb.GetBfTx(bftxID)
					if err != nil {
						return nil, nil
					}
					return bftx, nil
				},
			},
		},
	})

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}

	return result
}
