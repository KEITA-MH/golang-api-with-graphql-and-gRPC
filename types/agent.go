package types


import (
    "github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Agent struct{
    ID          primitive.ObjectID   `json:"_id" bson:"_id"`
    Username    string               `json:"username" bson:"username"`
    Email       string               `json:"email" bson:"email"`
}

var AgentType = graphql.NewObject(graphql.ObjectConfig{
    Name: "Agent",
    Fields: graphql.Fields{
        "_id": &graphql.Field{
            Type: graphql.ID,
        },
        "username": &graphql.Field{
            Type: graphql.String,
        },
        "email": &graphql.Field{
            Type: graphql.String,
        },
    },
})