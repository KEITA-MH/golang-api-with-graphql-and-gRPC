package types


import (
    "github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Departement struct{
    ID     primitive.ObjectID  `json:"_id" bson:"_id"`
    Name   string              `json:"name" bson:"name"`
    Agents    []Agent          `json:"agents" bson:"agents"`
}

var DepartementType = graphql.NewObject(graphql.ObjectConfig{
    Name: "Departement",
    Fields: graphql.Fields{
        "_id": &graphql.Field{
            Type: graphql.ID,
        },
        "name": &graphql.Field{
            Type: graphql.String,
        },
         "agents": &graphql.Field{
            Type: graphql.NewList(AgentType),
            Args: graphql.FieldConfigArgument{
                 "username": &graphql.ArgumentConfig{
                     Type: graphql.String,
                 },
                "email": &graphql.ArgumentConfig{
                     Type: graphql.String,
                 },
             },
             Resolve: func (p graphql.ResolveParams) (interface{}, error){
                 dpt := p.Source.(Departement)
                 if (p.Args["username"] != nil || p.Args["email"] != nil){
                     var agent []Agent
                     for _, item := range(dpt.Agents){
                         
                         if( item.Username == p.Args["username"]) {
                             agent = append(agent, item)
                         }
                         
                         if (item.Email == p.Args["email"]){
                              agent = append(agent, item)
                         }
                         
                     }
                     return agent, nil
                 }
                 
                 return dpt.Agents, nil
             },
        },
    },
})