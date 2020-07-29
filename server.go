package main 

import (
    "log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
    
    "Mamadou9527/Simple-CRUD-API-with-MUX-MongoDB-with-golang/queries"
)

func main(){
    schema, err := graphql.NewSchema(graphql.SchemaConfig{
        Query: graphql.NewObject(graphql.ObjectConfig{
            Name: "RootQuery",
            Fields: queries.GetRootFields(),
        }),
    })
    if err != nil {
        log.Fatal(err)
    }
    httpHandler := handler.New(&handler.Config{
        Schema: &schema,
        Pretty: true,
        GraphiQL: true,
    })
    
    http.Handle("/graphql", httpHandler)
    log.Print("ready: listening...\n")
    http.ListenAndServe(":4000", nil)
}