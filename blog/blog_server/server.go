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
	"github.com/mongodb/mongo-go-driver/bson"
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

func (*server) ReadBlog(ctx context.Context, req *blogpb.ReadBlogRequest) (*blogpb.ReadBlogResponse, error) {
	fmt.Println("Getting blog")
	id := req.GetBlogId()
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Can't parse oid"),
		)
	}
	data := &blogItem{}
	filter := bson.M{"_id": objectId}
	docRes := collection.FindOne(context.Background(), filter)
	if err := docRes.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find blog with ID: %v", err),
		)
	}

	return &blogpb.ReadBlogResponse{
		Blog: &blogpb.Blog{
			Id: data.ID.Hex(),
			AuthorId: data.AuthorID,
			Title: data.Title,
			Content: data.Content,
		},
	}, nil
}

func (*server) 	UpdateBlog(ctx context.Context, req *blogpb.UpdateBlogRequest) (*blogpb.UpdateBlogResponse, error) {
	fmt.Println("Updating Blog")
	blog := req.GetBlog()
	blogId := blog.GetId()

	objectId, err := primitive.ObjectIDFromHex(blogId)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Can't parse oid"),
		)
	}
	data := &blogItem{}
	filter := bson.M{"_id": objectId}
	docRes := collection.FindOne(context.Background(), filter)
	if err := docRes.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find blog with ID: %v", err),
		)
	}

	data.Title = blog.GetTitle()
	data.Content = blog.GetContent()
	data.AuthorID = blog.GetAuthorId()
	_, updateErr := collection.ReplaceOne(context.Background(), filter, data)
	if updateErr != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Can't update object in MongoDB %v", updateErr),
		)
	}
	return &blogpb.UpdateBlogResponse{
		Blog: dataToBlogPB(data),
	}, nil
}

func (*server) DeleteBlog(ctx context.Context, req *blogpb.DeleteBlogRequest) (*blogpb.DeleteBlogRespose, error) {
	fmt.Println("Deleting Blog")

	blogId := req.GetId()

	objectId, err := primitive.ObjectIDFromHex(blogId)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Can't parse oid"),
		)
	}
	filter := bson.M{"_id": objectId}

	res, err := collection.DeleteOne(context.Background(), filter)
	if err != nil{
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot delete blog %v", err),
		)
	}

	if res.DeletedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find blog with id %v", objectId),
		)
	}

	return &blogpb.DeleteBlogRespose{
		Id: objectId.Hex(),
	}, nil
}

func (*server) ListBlog(req *blogpb.ListBlogRequest, stream blogpb.BlogService_ListBlogServer) error{
	fmt.Println("Listing blogs")
	cursor, err := collection.Find(context.Background(), primitive.D{{}})
	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Can't list blogs: %v", err),
		)
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		data := &blogItem{}
		err := cursor.Decode(data)
		if err != nil{
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Can't decpde blog: %v", err),
			)
		}
		stream.Send(&blogpb.ListBlogResponse{Blog: dataToBlogPB(data)})
	}
	if err := cursor.Err(); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}
	return nil
}

func dataToBlogPB(data *blogItem) *blogpb.Blog{
	return &blogpb.Blog{
		Id: data.ID.Hex(),
		AuthorId: data.AuthorID,
		Title: data.Title,
		Content: data.Content,
	}
}

func (*server) CreateBlog(ctx context.Context, req *blogpb.CreateBlogRequest) (*blogpb.CreateBlogResponse, error){
	fmt.Println("Creating Blog")
	blog := req.GetBlog()

	data := blogItem{
		ID: primitive.NewObjectID(),
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
