package main

import (
	"context"
	"fmt"
	"log"
	"io"
	_ "sync"
	"time"

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
	//doServerStreaming(c)
	//doClientStreaming(c)
	doBidirectionalStreaming(c)
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

func doClientStreaming(c greetpb.GreetServiceClient){
	fmt.Println("Starting to do a client streaming rpc")

	requests := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest {
			Greeting: &greetpb.Greeting{
				FirstName: "Stefan",
			},
		},
		&greetpb.LongGreetRequest {
			Greeting: &greetpb.Greeting{
				FirstName: "John",
			},
		},
		&greetpb.LongGreetRequest {
			Greeting: &greetpb.Greeting{
				FirstName: "Joe",
			},
		},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling long greet ", err)
	}

	for _, msg := range requests {
		stream.Send(msg)
	}

	res, err := stream.CloseAndRecv()
	if err != nil{
		log.Fatalf("Error while receiving long greet ", err)
	}
	fmt.Printf("LongGreet Response: %v\n", res)
}

func doBidirectionalStreaming(c greetpb.GreetServiceClient){
	fmt.Println("Bidirectional streaming")

	requests := []*greetpb.GreetEveryoneRequest{
		&greetpb.GreetEveryoneRequest {
			Greeting: &greetpb.Greeting{
				FirstName: "Stefan",
			},
		},
		&greetpb.GreetEveryoneRequest {
			Greeting: &greetpb.Greeting{
				FirstName: "John",
			},
		},
		&greetpb.GreetEveryoneRequest {
			Greeting: &greetpb.Greeting{
				FirstName: "Joe",
			},
		},
	}

	//Initialize WaitGroup
	waitch := make(chan struct{})

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while calling long greet ", err)
	}

	go func(){
		for _, msg := range requests {
			fmt.Printf("Sending Request %v \n", msg.GetGreeting().GetFirstName())
			stream.Send(msg)
			time.Sleep(500 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	go func(){
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				close(waitch)
				break
			}
			if err != nil{
				log.Fatalf("Error receiving stream from server: ", err)
				close(waitch)
				break
			}
			fmt.Printf("LongGreet Response: %v\n", res)
		}
	}()

	<- waitch
}
