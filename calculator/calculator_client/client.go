package main

import (
	"context"
	"fmt"
	"log"

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
	doUnary(c)
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
