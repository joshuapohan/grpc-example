package main

import(
  "fmt"
  "log"
  "net"
  "time"

  "google.golang.org/grpc"

  "github.com/joshuapohan/grpc-example/prime_decomposer/primepb"
)

type server struct{}

func (*server) 	Decompose(req *primepb.PrimeRequest, stream primepb.PrimeDecomposer_DecomposeServer) error {
    fmt.Println("Prime number decomposer called")
    input := req.GetInput()
    var factor int32 = 2
    for input > 1 {
      if input % factor == 0{
        result := &primepb.PrimeResponse{
          Result: int32(factor),
        }
        input = input / factor
        stream.Send(result)
        time.Sleep(1000 * time.Millisecond)
      } else{
        factor++
      }
    }
    return nil
}

func main(){
  lis, err := net.Listen("tcp", "0.0.0.0:50051")
  if err != nil{
    log.Fatalf("Failed to listen: %v", err)
  }

  s := grpc.NewServer()
  primepb.RegisterPrimeDecomposerServer(s, &server{})
  if err := s.Serve(lis); err != nil{
    log.Fatalf("Failed to server: %v", err)
  }
}
