package main

import (
	"context"
	"fmt"
	"go-grpc-microservice/microservice/microservicepb"
	"log"
	"time"

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
	doClientStreamc(c)

}

func doClientStreamc(c microservicepb.MicroserviceServiceClient) {
	fmt.Println("Starting doClientStreamc")

	request := []*microservicepb.MicroserviceRequest{
		&microservicepb.MicroserviceRequest{
			Microservice: &microservicepb.Microservice{
				FirstName: "Gurudath1",
			},
		},
		&microservicepb.MicroserviceRequest{
			Microservice: &microservicepb.Microservice{
				FirstName: "Gurudath2",
			},
		},
		&microservicepb.MicroserviceRequest{
			Microservice: &microservicepb.Microservice{
				FirstName: "Gurudath3",
			},
		},
		&microservicepb.MicroserviceRequest{
			Microservice: &microservicepb.Microservice{
				FirstName: "Gurudath4",
			},
		},
		&microservicepb.MicroserviceRequest{
			Microservice: &microservicepb.Microservice{
				FirstName: "Gurudath5",
			},
		},
		&microservicepb.MicroserviceRequest{
			Microservice: &microservicepb.Microservice{
				FirstName: "Gurudath6",
			},
		},
		&microservicepb.MicroserviceRequest{
			Microservice: &microservicepb.Microservice{
				FirstName: "Gurudath7",
			},
		},
	}
	stream, err := c.Microservice(context.Background())
	if err != nil {
		log.Fatalf("Error Client %v", err)
	}

	for _, req := range request {
		fmt.Printf(" Sending req %v\n", req)
		stream.Send(req)
		time.Sleep(100 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error Client CloseAndRecv %v", err)
	}
	log.Printf("Response of Microservice: %v", res)
}
