package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http" // Provides HTTP client and server implementations.

	"github.com/blockfreight/go-bftx/api/graphqlObj"
	"github.com/blockfreight/go-bftx/api/handlers"
	"github.com/blockfreight/go-bftx/lib/app/bf_tx"
	"github.com/blockfreight/go-bftx/lib/pkg/leveldb" // Provides some useful functions to work with LevelDB.
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	},
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"transaction": &graphql.Field{
				Type: graphqlObj.TransactionType,
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
				Type: graphqlObj.TransactionType,
				Args: graphql.FieldConfigArgument{
					"Type": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"Properties": &graphql.ArgumentConfig{
						Description: "Transaction properties.",
						Type:        graphqlObj.PropertiesInput,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					bftx := bf_tx.BF_TX{}
					jsonProperties, err := json.Marshal(p.Args)
					if err = json.Unmarshal([]byte(jsonProperties), &bftx); err != nil {
						fmt.Printf("err")
						fmt.Print(err)
					}

					bftx, err = handlers.ConstructBfTx(bftx)
					if err != nil {
						return nil, err
					}

					return bftx, err
				},
			},
		},
	},
)

func Start() error {
	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	http.Handle("/graphql", h)
	fmt.Println("Now server is running on: http://localhost:12345")
	return http.ListenAndServe(":12345", nil)
}

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

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	return result
}
