package main

import (
	"context"
	"fmt"
	"log"
	"io"
	"time"

	"google.golang.org/grpc"

	"github.com/joshuapohan/grpc-example/calculator/calculatorpb"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect : %v", err)
	}
	defer cc.Close()

	c := calculatorpb.NewCalculateServiceClient(cc)
	//doUnary(c)
	//doAverage(c)
	doMaximum(c)
}

func doUnary(c calculatorpb.CalculateServiceClient) {
	req := &calculatorpb.CalculatorRequest{
		CalculatorInput: &calculatorpb.CalculatorInput{
			A: 4,
			B: 5,
		},
	}
	res, err := c.Add(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to call calculate method %v", err)
	}
	fmt.Printf("Response : %v", res)
}

func doAverage(c calculatorpb.CalculateServiceClient){
	req := []*calculatorpb.AverageRequest{
		&calculatorpb.AverageRequest{
			Input:5,
		},
		&calculatorpb.AverageRequest{
			Input:13,
		},
		&calculatorpb.AverageRequest{
			Input:32,
		},
		&calculatorpb.AverageRequest{
			Input:8,
		},
	}

	stream, err := c.Average(context.Background())
	if err != nil {
		log.Fatalf("Failed executing rpc function")
	}
	for _, msg := range req {
		stream.Send(msg)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving average ")
	}
	fmt.Printf("Average Response: %v ", res)
}

func doMaximum(c calculatorpb.CalculateServiceClient) {
	input := []*calculatorpb.MaximumRequest{
		&calculatorpb.MaximumRequest{
			Input:3,
		},
		&calculatorpb.MaximumRequest{
			Input:2,
		},
		&calculatorpb.MaximumRequest{
			Input:7,
		},
		&calculatorpb.MaximumRequest{
			Input:9,
		},
		&calculatorpb.MaximumRequest{
			Input:1,
		},
		&calculatorpb.MaximumRequest{
			Input:33,
		},
	}

	waitch := make(chan struct{})

	fmt.Println("Calling maximum rpc service")
	stream, err := c.Maximum(context.Background())
	if err != nil {
		log.Fatalf("Error when calling maximum service: %v", err)
	}

	go func(){
		for _, msg := range input {
			fmt.Printf("Sending request: %v \n", msg)
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
			if err != nil {
				log.Fatalf("Error reading server response stream: %v", err)
				break
			}
			fmt.Println("Maximum value: ", res.GetResult())
		}
	}()

	<- waitch
}
