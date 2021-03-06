package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"io"
	"math"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	
	"github.com/joshuapohan/grpc-example/calculator/calculatorpb"
)

type server struct{}

func (*server) Add(ctx context.Context, req *calculatorpb.CalculatorRequest) (*calculatorpb.CalculatorResponse, error) {
	fmt.Printf("Add invoked with input %v", req)
	a := req.GetCalculatorInput().GetA()
	b := req.GetCalculatorInput().GetB()
	result := a + b
	res := &calculatorpb.CalculatorResponse{
		Result: result,
	}
	return res, nil
}

func (*server) Average(stream calculatorpb.CalculateService_AverageServer) error {
	fmt.Printf("Average invoked with client streaming")
	count := 0
	var sum int32 = 0
	var res float32 = 0.0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			if count > 0 {
				res = float32(sum) / float32(count)
			} else{
				res = 0
			}
			return stream.SendAndClose(&calculatorpb.AverageResponse{
				Result: res,
			})
		}
		if err != nil {
			log.Fatalf("Error when reading client stream: ", err)
		}
		value := req.GetInput()
		sum += value
		count++
	}
}

func (*server) Maximum(stream calculatorpb.CalculateService_MaximumServer) error {
	fmt.Println("Bidirection maximum function called")
	var curMax int32 = 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error when reading client stream: %v ", err)
			return err
		}
		curValue := req.GetInput()
		if curValue > curMax {
			curMax = curValue
		}
		stream.Send(&calculatorpb.MaximumResponse{
			Result: curMax,
		})
	}
	return nil
}

func (*server) 	SquareRoot(c context.Context, req *calculatorpb.SquareRootRequest) (*calculatorpb.SquareRootResponse, error) {
	fmt.Println("Calculate Square Root rpc called")
	input := req.GetInput()


	if input > 0 {
		res := &calculatorpb.SquareRootResponse{
			Result: math.Sqrt(float64(input)),
		}
		return res, nil
	} else{
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negative number: %v", input),
		)
	}
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	calculatorpb.RegisterCalculateServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
}
