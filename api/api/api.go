package api

import (
	"net/http" // Provides HTTP client and server implementations.
	//"github.com/gorilla/mux" //Implements a request router and dispatcher for matching incoming requests to their respective handler
	//"github.com/blockfreight/go-bftx/api/transaction"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/blockfreight/go-bftx/lib/pkg/leveldb" // Provides some useful functions to work with LevelDB.
	"github.com/graphql-go/graphql"
)

/*
   Create User object type with fields "id" and "name" by using GraphQLObjectTypeConfig:
       - Name: name of object type
       - Fields: a map of fields by using GraphQLFields
   Setup type of field use GraphQLFieldConfig
*/
var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var transactionType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Transaction",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"portOfDischarge": &graphql.Field{
				Type: graphql.String,
			},
			/*"PortOfLoading": &graphql.Field{
				Type: graphql.String,
			},
			"PortOfDischarge": &graphql.Field{
				Type: graphql.String,
			},*/
		},
	},
)

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: queryType,
	},
)

type user struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type transaction struct {
	ID                  string `json:"id"`
	Shipper             string
	BolNum              int
	RefNum              int
	Consignee           string
	Vessel              int
	PortOfLoading       int
	PortOfDischarge     int `json:"portOfDischarge"`
	NotifyAddress       string
	DescOfGoods         string
	GrossWeight         int
	FreightPayableAmt   int
	FreightAdvAmt       int
	GeneralInstructions string
	DateShipped         string
	NumBol              int
	MasterInfo          string
	AgentForMaster      string
	AgentForOwner       string
	Type                string `json:"type"`
}

var data map[string]user

func Start() error {
	//configuration, _ := config.LoadConfiguration()
	/*router := mux.NewRouter()
	router.HandleFunc("/fulltransaction", transaction.FullTransactionBfTx).Methods("POST")
	router.HandleFunc("/transaction/construct", transaction.ConstructBfTx).Methods("POST")
	router.HandleFunc("/transaction/sign/{id}", transaction.SignBfTx).Methods("PUT")
	router.HandleFunc("/transaction/broadcast/{id}", transaction.BroadcastBfTx).Methods("PUT")
	router.HandleFunc("/transaction/{id}", transaction.GetTransaction).Methods("GET")
	router.HandleFunc("/transaction", transaction.GetTransaction).Methods("GET")*/

	_ = importJSONDataFromFile("data.json", &data)

	http.HandleFunc("/graphql", graphRoute)
	fmt.Println("Now server is running on port 12345")
	fmt.Println("Test with Get      : curl -g 'http://localhost:12345/graphql?query={user(id:\"1\"){name}}'")
	return http.ListenAndServe(":12345", nil)

}

func graphRoute(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")

	result := executeQuery(query, schema)
	json.NewEncoder(w).Encode(result)
}

/*
   Create Query object type with fields "user" has type [userType] by using GraphQLObjectTypeConfig:
       - Name: name of object type
       - Fields: a map of fields by using GraphQLFields
   Setup type of field use GraphQLFieldConfig to define:
       - Type: type of field
       - Args: arguments to query with current field
       - Resolve: function to query data using params from [Args] and return value with current type
*/
var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, isOK := p.Args["id"].(string)
					if isOK {
						return data[idQuery], nil
					}
					return nil, nil
				},
			},
			"transaction": &graphql.Field{
				Type: transactionType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					bftxId, isOK := p.Args["id"].(string)
					if !isOK {
						return nil, nil
					}

					bftx, err := leveldb.GetBfTx(bftxId)
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
	fmt.Printf("%+v\n", result)
	return result
}

//Helper function to import json from file to map
func importJSONDataFromFile(fileName string, result interface{}) (isOK bool) {
	isOK = true
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Print("Error:", err)
		isOK = false
	}
	err = json.Unmarshal(content, result)
	if err != nil {
		isOK = false
		fmt.Print("Error:", err)
	}
	return
}
