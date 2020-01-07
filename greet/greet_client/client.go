package main

import (
	"context"
	"fmt"
	"log"
	"io"

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
	//doUnary(c)
	doServerStreaming(c)
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

func doServerStreaming(c greetpb.GreetServiceClient){
	fmt.Println("Server streaming grpc")
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Joshua",
			LastName: "Pohan",
		},
	}
	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling server streaming GreetAnyTimes rpc: %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading stream,: %v", err)
		}
		log.Printf("Response from GreetManyTimes: %v", msg.GetResult())
	}
}
