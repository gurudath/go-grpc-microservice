package main

import (
	"fmt"
	"go-grpc-microservice/microservice/microservicepb"
	"log"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Microservice(req *microservicepb.MicroserviceRequest, stream microservicepb.MicroserviceService_MicroserviceServer) error {
	fmt.Println("Triggering Microservice Server")
	firstName := req.GetMicroservice().GetFirstName()
	for i := 0; i < 10; i++ {
		result := "Hellow " + firstName + " Number " + strconv.Itoa(i)
		res := &microservicepb.MicroserviceResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
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
