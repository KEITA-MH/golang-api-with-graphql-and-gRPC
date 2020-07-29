package database


import (
    "log"
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var Collection *mongo.Collection

func init(){
    
    clientOptions := options.Client().ApplyURI(URL)
    client, err := mongo.Connect(context.TODO(), clientOptions)
    
    if err != nil {
        log.Fatal(err)
    }
    
    err = client.Ping(context.TODO(), nil)
    
    if err != nil {
        log.Fatal(err)
    }
    
    Collection = client.Database("conFusion").Collection("departemants")
}