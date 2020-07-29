package queries

import (
	"Mamadou9527/Simple-CRUD-API-with-MUX-MongoDB-with-golang/types"
    "Mamadou9527/Simple-CRUD-API-with-MUX-MongoDB-with-golang/database"
    "go.mongodb.org/mongo-driver/bson"
	"github.com/graphql-go/graphql"
	"log"
    "context"
)

// GetUserQuery returns the queries available against user type.
func GetDepartementQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(types.DepartementType),
        Args: graphql.FieldConfigArgument{
            "name": &graphql.ArgumentConfig{
                Type:graphql.String,
            },
        },
		Resolve: func (params graphql.ResolveParams) (interface{}, error) {
            var departements []types.Departement
            if(params.Args["name"] != nil){
                var departement types.Departement
                param := params.Args["name"].(string)
                        
                err := database.Collection.FindOne(context.TODO(), bson.M{"name": param}).Decode(&departement)

                if err != nil {
                    log.Fatal(err)
                }
                departements = append(departements,departement) 
                return departements, nil
            }
            
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

                    departements = append(departements, departement)
                }
                        
            return departements, nil
        },
	}
}
