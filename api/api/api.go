package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http" // Provides HTTP client and server implementations.

	"github.com/blockfreight/go-bftx/lib/app/bf_tx"

	"github.com/blockfreight/go-bftx/api/transaction"
	"github.com/blockfreight/go-bftx/lib/pkg/leveldb" // Provides some useful functions to work with LevelDB.
	"github.com/graphql-go/graphql"
)

func Start() error {
	http.HandleFunc("/graphql", graphRoute)
	fmt.Println("Now server is running on port 12345")
	fmt.Println("Test with Get      : curl -g 'http://localhost:12345/graphql?query={transaction(id:<BFTX-ID>){Id}}'")
	return http.ListenAndServe(":12345", nil)

}

var issueDetails = graphql.NewInputObject(
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

var masterInfo = graphql.NewInputObject(
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

var agentForMaster = graphql.NewInputObject(
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

var agentForOwner = graphql.NewInputObject(
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

var propertiesType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "Properties",
		Fields: graphql.InputObjectConfigFieldMap{
			"Shipper": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"BolNum": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			"RefNum": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			"Consignee": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"Vessel": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			"PortOfLoading": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			"PortOfDischarge": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			"NotifyAddress": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"DescOfGoods": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"GrossWeight": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			"FreightPayableAmt": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			"FreightAdvAmt": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			"GeneralInstructions": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"DateShipped": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"IssueDetails": &graphql.InputObjectFieldConfig{
				Type: issueDetails,
			},
			"NumBol": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			"MasterInfo": &graphql.InputObjectFieldConfig{
				Type: masterInfo,
			},
			"AgentForMaster": &graphql.InputObjectFieldConfig{
				Type: agentForMaster,
			},
			"AgentForOwner": &graphql.InputObjectFieldConfig{
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
		Query:    queryType,
		Mutation: mutationType,
	},
)

func graphRoute(w http.ResponseWriter, r *http.Request) {
	query, _ := ioutil.ReadAll(r.Body)
	result := executeQuery(string(query), schema)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
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

var mutationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"constructBFTX": &graphql.Field{
				Type: transactionType,
				Args: graphql.FieldConfigArgument{
					"Type": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"Properties": &graphql.ArgumentConfig{
						Description: "Transaction properties.",
						Type:        propertiesType,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					bftx := bf_tx.BF_TX{}
					jsonProperties, err := json.Marshal(p.Args)
					if err = json.Unmarshal([]byte(jsonProperties), &bftx); err != nil {
						fmt.Printf("err")
						fmt.Print(err)
					}
					fmt.Printf("%+v\n", bftx)

					bftx, err = transaction.ConstructBfTx(bftx)
					if err != nil {
						return nil, err
					}

					return bftx, err
				},
			},
		},
	},
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {

	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	return result
}
