package main

import(
  "context"
  "fmt"
  "io"
  "log"

  "google.golang.org/grpc"

  "github.com/joshuapohan/grpc-example/prime_decomposer/primepb"
)

func main(){
  cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
  if err != nil {
    log.Fatalf("Could not connect : %v", err)
  }
  defer cc.Close()

  c := primepb.NewPrimeDecomposerClient(cc)
  doServerStreaming(c)
}

func doServerStreaming(c primepb.PrimeDecomposerClient){
  fmt.Println("Server streaming grpc")
  req := &primepb.PrimeRequest{
    Input: 120,
  }
  resStream, err := c.Decompose(context.Background(), req)
  if err != nil{
    log.Fatalf("Error while calling server streaming decompose: %v", err)
  }
  for {
    msg, err := resStream.Recv()
    if err == io.EOF{
      break
    }
    if err != nil{
      log.Fatalf("Error while reading stream: %v", err)
    }
    fmt.Println("Response from decompose prime:", msg.GetResult())
  }
}
