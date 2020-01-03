package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	"github.com/joshuapohan/grpc-example/greet/greetpb"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect : %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Joshua",
			LastName:  "Pohan",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to call greet method %v", err)
	}
	fmt.Printf("Response : %v", res)
}
