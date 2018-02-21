package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http" // Provides HTTP client and server implementations.
	"strconv"

	"github.com/blockfreight/go-bftx/api/graphqlObj"
	apiHandler "github.com/blockfreight/go-bftx/api/handlers"
	"github.com/blockfreight/go-bftx/lib/app/bf_tx" // Provides some useful functions to work with LevelDB.
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
			"getTransaction": &graphql.Field{
				Type: graphqlObj.TransactionType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					bftxID, isOK := p.Args["id"].(string)
					if !isOK {
						return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
					}

					return apiHandler.GetTransaction(bftxID)
				},
			},
			"queryTransaction": &graphql.Field{
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

					return apiHandler.QueryTransaction(bftxID)
				},
			},
			"getInfo": &graphql.Field{
				Type: graphqlObj.InfoType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return apiHandler.GetInfo()
				},
			},
			"getTotal": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return apiHandler.GetTotal()
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
						return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
					}

					return apiHandler.ConstructBfTx(bftx)
				},
			},
			"encryptBFTX": &graphql.Field{
				Type: graphqlObj.TransactionType,
				Args: graphql.FieldConfigArgument{
					"Id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					bftxID, isOK := p.Args["Id"].(string)
					if !isOK {
						return nil, nil
					}

					return apiHandler.EncryptBfTx(bftxID)
				},
			},
			"decryptBFTX": &graphql.Field{
				Type: graphqlObj.TransactionType,
				Args: graphql.FieldConfigArgument{
					"Id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					bftxID, isOK := p.Args["Id"].(string)
					if !isOK {
						return nil, nil
					}

					return apiHandler.DecryptBfTx(bftxID)
				},
			},
			"signBFTX": &graphql.Field{
				Type: graphqlObj.TransactionType,
				Args: graphql.FieldConfigArgument{
					"Id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					bftxID, isOK := p.Args["Id"].(string)
					if !isOK {
						return nil, nil
					}

					return apiHandler.SignBfTx(bftxID)
				},
			},
			"broadcastBFTX": &graphql.Field{
				Type: graphqlObj.TransactionType,
				Args: graphql.FieldConfigArgument{
					"Id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					bftxID, isOK := p.Args["Id"].(string)
					if !isOK {
						return nil, nil
					}

					return apiHandler.BroadcastBfTx(bftxID)
				},
			},
		},
	},
)

func Start() error {
	http.HandleFunc("/bftx-api", httpHandler(&schema))
	fmt.Println("Now server is running on: http://localhost:12345")
	return http.ListenAndServe(":12345", nil)
}

func httpHandler(schema *graphql.Schema) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		httpStatusResponse := http.StatusOK
		// parse http.Request into handler.RequestOptions
		opts := handler.NewRequestOptions(r)

		// inject context objects http.ResponseWrite and *http.Request into rootValue
		// there is an alternative example of using `net/context` to store context instead of using rootValue
		rootValue := map[string]interface{}{
			"response": rw,
			"request":  r,
			"viewer":   "john_doe",
		}

		// execute graphql query
		// here, we passed in Query, Variables and OperationName extracted from http.Request
		params := graphql.Params{
			Schema:         *schema,
			RequestString:  opts.Query,
			VariableValues: opts.Variables,
			OperationName:  opts.OperationName,
			RootObject:     rootValue,
		}
		result := graphql.Do(params)
		js, err := json.Marshal(result)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)

		}
		if result.HasErrors() {
			httpStatusResponse, err = strconv.Atoi(result.Errors[0].Error())
			if err != nil {
				httpStatusResponse = http.StatusInternalServerError
			}
		}
		rw.WriteHeader(httpStatusResponse)

		rw.Write(js)

	}

}
