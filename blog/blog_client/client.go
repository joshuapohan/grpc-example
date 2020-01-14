package main

import (
	"context"
	"fmt"
	"log"
    "io"

	"google.golang.org/grpc"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/grpc/codes"

	"github.com/joshuapohan/grpc-example/blog/blogpb"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect : %v", err)
	}
	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)
	createBlog(c)
}


func createBlog(c blogpb.BlogServiceClient) {
    fmt.Println("Creating blog")
    req := &blogpb.CreateBlogRequest{
        Blog: &blogpb.Blog{
            AuthorId: "Stephane",
            Title: "My First Blog",
            Content: "Content of the first blog",
        },
    }

	res, err := c.	CreateBlog(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to create blog %v", err)
	}
	fmt.Printf("Blog has been created : %v", res)
    blogId := res.GetBlog().GetId()

    //Read blog
    fmt.Println("Reading blog")
	_, err2 := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{BlogId:"12324"})
    if err2 != nil {
		log.Println("Failed to read blog %v", err2)
	}

    res2, err2 := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{BlogId:blogId})
    if err2 != nil {
		log.Println("Failed to read blog %v", err2)
	}
    fmt.Println("Received blog :", res2.GetBlog())

    //Update blog
    fmt.Println("Updating blog")
    newBlog := &blogpb.Blog{
        Id: blogId,
        Title: "updated title",
        Content: "updated content",
        AuthorId: "New author",
    }
    updatedRes, updatedErr := c.UpdateBlog(context.Background(), &blogpb.UpdateBlogRequest{Blog: newBlog})
    if updatedErr != nil {
        log.Println("Failed to read blog %v", updatedErr)
    }
    fmt.Println("Updated blog :", updatedRes.GetBlog())

    //Delete Blog
    fmt.Println("Deleting blog")
    deleteRes, deleteErr := c.	DeleteBlog(context.Background(), &blogpb.DeleteBlogRequest{Id: blogId})
    if deleteErr != nil {
        log.Println("Failed to delete blog %v", deleteErr)
    }
    fmt.Println("Deleted blog :", deleteRes.GetId())

    //List BlogId
    fmt.Println("Listing blogs")
    listStream, listErr := c.ListBlog(context.Background(),&blogpb.ListBlogRequest{})
    if listErr != nil{
        log.Println("Failed to list blog: ", listErr)
    }
    for {
        ListBlogResponse, err := listStream.Recv()
        if err == io.EOF {
            break
        }
        if err != nil{
            log.Println("error reading stream ", err)
            break
        }
        fmt.Printf("Received blog: %v", ListBlogResponse.GetBlog())
    }
}
