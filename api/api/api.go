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

var issueDetailsInput = graphql.NewInputObject(
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

var issueDetails = graphql.NewObject(
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

var masterInfoInput = graphql.NewInputObject(
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

var masterInfo = graphql.NewObject(
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

var agentForMasterInput = graphql.NewInputObject(
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

var agentForMaster = graphql.NewObject(
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

var agentForOwnerInput = graphql.NewInputObject(
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

var agentForOwner = graphql.NewObject(
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

var propertiesInput = graphql.NewInputObject(
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
				Type: issueDetailsInput,
			},
			"NumBol": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			"MasterInfo": &graphql.InputObjectFieldConfig{
				Type: masterInfoInput,
			},
			"AgentForMaster": &graphql.InputObjectFieldConfig{
				Type: agentForMasterInput,
			},
			"AgentForOwner": &graphql.InputObjectFieldConfig{
				Type: agentForOwnerInput,
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
	query := r.URL.Query().Get("query")
	if query == "" {
		params, _ := ioutil.ReadAll(r.Body)
		query = string(params)
	}

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

					fmt.Printf("%+v\n", bftx)

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
						Type:        propertiesInput,
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
