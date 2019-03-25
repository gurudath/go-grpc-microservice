package main

import (
	"context"
	"fmt"
	"go-grpc-microservice/microservice/microservicepb"
	"io"
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
	doBiDirectionMicroservice(c)

}

func doBiDirectionMicroservice(c microservicepb.MicroserviceServiceClient) {
	fmt.Println("Microservice Bi Direction")

	stream, err := c.Microservice(context.Background())

	if err != nil {
		log.Fatalf("Error while creating stream: %v", err)
		return
	}

	requests := []*microservicepb.MicroserviceRequest{
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
		&microservicepb.MicroserviceRequest{
			Microservice: &microservicepb.Microservice{
				FirstName: "Gurudath8",
			},
		},
		&microservicepb.MicroserviceRequest{
			Microservice: &microservicepb.Microservice{
				FirstName: "Gurudath9",
			},
		},
		&microservicepb.MicroserviceRequest{
			Microservice: &microservicepb.Microservice{
				FirstName: "Gurudath10",
			},
		},
	}
	waitc := make(chan struct{})

	go func() {
		for _, req := range requests {
			fmt.Printf("Sending Message %v\n", req)
			stream.Send(req)
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while reciving: %v\n", err)
				break
			}
			fmt.Printf("Recived: %v\n", res.GetResult())
		}
		close(waitc)
	}()

	<-waitc
}
