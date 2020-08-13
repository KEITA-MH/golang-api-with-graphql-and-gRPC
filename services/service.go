package main

import(
	"log"
	"context"
	"errors"
	"net"

	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc"
	"Mamadou9527/Simple-CRUD-API-with-MUX-MongoDB-with-golang/database"
	"Mamadou9527/Simple-CRUD-API-with-MUX-MongoDB-with-golang/types"
	"google.golang.org/grpc/reflection"
	"Mamadou9527/Simple-CRUD-API-with-MUX-MongoDB-with-golang/services/proto"
)
type departemantService struct{}

func (departemantService) GetDepartements(_ context.Context,  in *proto.Empty) (*proto.Departemants, error) {
	cur, err := database.Collection.Find(context.TODO(), bson.D{{}})

	if err != nil {
		return &proto.Departemants{}, errors.New("Something went wrong!")
	}
	defer cur.Close(context.TODO())


	 for cur.Next(context.TODO()){
        var departement types.Departement
        err := cur.Decode(&departement)
        if err != nil{
            return &proto.Departemants{}, errors.New("Something went wrong!")
        }

        log.Println(departement);
    }

    return &proto.Departemants{}, nil
}


func main(){

	server := grpc.NewServer()
	proto.RegisterDepartemantServiceServer(server, departemantService{})

	listener, err := net.Listen("tcp", ":6000")
	if err != nil {
		log.Fatal("Error creating listener: ", err.Error())
	}
	reflection.Register(server)
	server.Serve(listener)
}