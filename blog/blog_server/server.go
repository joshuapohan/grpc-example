package main

import (
	"fmt"
	"log"
	"net"
    "os"
    "os/signal"
    "context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"github.com/mongodb/mongo-go-driver/mongo"
    "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"


	"github.com/joshuapohan/grpc-example/blog/blogpb"
)

var collection *mongo.Collection

type blogItem struct {
    ID primitive.ObjectID `bson:"_id, omitempty"`
    AuthorID string `bson:"author_id"`
    Content string `bson:"content"`
    Title string `bson:"title"`
}

type server struct{
}

func (*server) CreateBlog(ctx context.Context, req *blogpb.CreateBlogRequest) (*blogpb.CreateBlogResponse, error){
	blog := req.GetBlog()

	data := blogItem{
		AuthorID: blog.GetAuthorId(),
		Title: blog.GetTitle(),
		Content: blog.GetContent(),
	}

	res, err := collection.InsertOne(context.Background(), data)
	if err !=nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal Error: %v", err),
		)
	}
	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		if err !=nil {
			return nil, status.Errorf(
				codes.Internal,
				fmt.Sprintf("Can't convert to oid"),
			)
		}
	}
	return &blogpb.CreateBlogResponse{
		Blog: &blogpb.Blog{
			Id: oid.Hex(),
			AuthorId: blog.GetAuthorId(),
			Title: blog.GetTitle(),
			Content: blog.GetContent(),
		},
	}, nil
}

func main() {
    //detailed error if crash
    log.SetFlags(log.LstdFlags | log.Lshortfile)

    //connect to mongodb
    fmt.Println("Connecting to mongodb")
    client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil {
        log.Fatal(err)
    }
    err = client.Connect(context.TODO())
    if err != nil {
        log.Fatal(err)
    }

    collection = client.Database("mydb").Collection("blog")


    fmt.Println("Blog service started")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	blogpb.RegisterBlogServiceServer(s, &server{})

    go func(){
        fmt.Println("Starting server")
        if err := s.Serve(lis); err != nil {
        	log.Fatalf("Failed to listen: %v", err)
        }
    }()

    // wait for control c to exit
    ch := make(chan os.Signal, 1)
    signal.Notify(ch, os.Interrupt)

    //block until signal is Received
    <-ch
    fmt.Println("Stopping the server")
    s.Stop()
    fmt.Println("Stopping listener")
    lis.Close()
    fmt.Println("Closing mongodb connection")
    client.Disconnect(context.TODO())
    fmt.Println("Closing application")
}
