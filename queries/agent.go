package queries

import (
	"Mamadou9527/Simple-CRUD-API-with-MUX-MongoDB-with-golang/types"
    "Mamadou9527/Simple-CRUD-API-with-MUX-MongoDB-with-golang/database"
    "go.mongodb.org/mongo-driver/bson"
	"github.com/graphql-go/graphql"
	"log"
    "context"
)


func GetAgentByUsername() *graphql.Field {
	return &graphql.Field{
		Type: types.AgentType,
        Args: graphql.FieldConfigArgument{
            "username": &graphql.ArgumentConfig{
                Type: graphql.NewNonNull(graphql.String),
            },
        },
        Resolve: func (params graphql.ResolveParams) (interface{}, error) {
                        
            cur, err := database.Collection.Find(context.TODO(), bson.D{{}})

            if err != nil {
                log.Fatal(err)
            }
            defer cur.Close(context.TODO())
            for cur.Next(context.TODO()){
                    var departement types.Departement
                    err := cur.Decode(&departement)
                    if err != nil{
                        log.Fatal(err)
                    }
                var item types.Agent
                for _,item = range(departement.Agents){
                    if (item.Username == params.Args["username"].(string)){
                     return item, nil
                        }
                    
                    }
                }
              return nil,nil   
            },
	}
}