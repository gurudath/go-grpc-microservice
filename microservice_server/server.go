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
	fmt.Println("Microservice Bi Direction")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error in server %v", err)
			return err
		}
		firstName := req.GetMicroservice().GetFirstName()
		result := "Hellow " + firstName
		stdErr := stream.Send(&microservicepb.MicroserviceResponse{
			Result: result,
		})
		if stdErr != nil {
			log.Fatalf("Error in data from client: %v", stdErr)
			return stdErr
		}
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
