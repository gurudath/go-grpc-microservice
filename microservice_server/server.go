package main

import (
	"context"
	"fmt"
	"go-grpc-microservice/microservice/microservicepb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Microservice(ctx context.Context, req *microservicepb.MicroserviceRequest) (*microservicepb.MicroserviceResponse, error) {
	firstName := req.GetMicroservice().GetFirstName()
	result := "Hellow " + firstName
	res := &microservicepb.MicroserviceResponse{
		Result: result,
	}
	log.Printf("Request of Microservice: %v", req)
	return res, nil
}

func main() {
	fmt.Println("Hello Guru")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	microservicepb.RegisterMicroserviceServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
