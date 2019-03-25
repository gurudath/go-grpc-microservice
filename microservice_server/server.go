package main

import (
	"fmt"
	"go-grpc-microservice/microservice/microservicepb"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Microservice(stream microservicepb.MicroserviceService_MicroserviceServer) error {
	fmt.Println("Microservice Server Reqest")
	result := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&microservicepb.MicroserviceResponse{
				Result: result,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream %v", err)
		}
		firstName := req.GetMicroservice().GetFirstName()
		result += "Hello " + firstName + "! "
		fmt.Printf("Microservice Server Reqest %v\n", result)
	}
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
