package main

import (
	"context"
	"fmt"
	"go-grpc-microservice/microservice/microservicepb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hellow i am client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not comment: %v\n", err)
	}
	defer cc.Close()

	c := microservicepb.NewMicroserviceServiceClient(cc)
	// fmt.Printf("Creating client: %f\n", c)
	doUnary(c)

}

func doUnary(c microservicepb.MicroserviceServiceClient) {
	fmt.Println("Starting Unary")
	req := &microservicepb.MicroserviceRequest{
		Microservice: &microservicepb.Microservice{
			FirstName: "Gurudath",
			LastName:  "BN",
		},
	}
	res, err := c.Microservice(context.Background(), req)
	if err != nil {
		log.Fatalf("Error Client %v", err)
	}
	log.Printf("Response of Microservice: %v", res.Result)
}
