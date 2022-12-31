package main

import (
	"blog_go/dao"
	"blog_go/proto"
	"blog_go/router"
	"blog_go/service"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

func main() {
	var err error

	// 连接数据库
	err = dao.InitMySQL()
	if err != nil {
		return
	}

	// 用goroutine启动grpc服务端
	go useGrpc()

	r := router.SetUpRouter()

	err = r.Run(":9091")
	if err != nil {
		return
	}
}

type server struct {
	proto.UnimplementedBlogServiceServer
}

func (s *server) GetAllBlogs(ctx context.Context, request *proto.BlogsRequest) (*proto.BlogsResponse, error) {
	blogs, err := service.GetAllBlogs()
	fmt.Println("request to all blogs")
	if err != nil {
		return nil, err
	}
	var result []*proto.BlogResponse
	for _, blog := range blogs {
		result = append(result,
			&proto.BlogResponse{
				Id:         int32(blog.Id),
				Category:   blog.Category,
				CreateTime: blog.CreateTime,
				Context:    blog.Content,
				Title:      blog.Title,
			})
	}
	return &proto.BlogsResponse{Blogs: result}, nil
}

func (s *server) GetBlogById(ctx context.Context, in *proto.BlogRequest) (*proto.BlogResponse, error) {
	fmt.Println("get client request ID : ", in.Id)
	id := in.GetId()
	blog, err := service.GetBlogById(int(id))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &proto.BlogResponse{
		Id:         int32(blog.Id),
		Category:   blog.Category,
		CreateTime: blog.CreateTime,
		Context:    blog.Content,
		Title:      blog.Title,
	}, nil
}

func useGrpc() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return
	}
	s := grpc.NewServer()
	proto.RegisterBlogServiceServer(s, &server{})

	err = s.Serve(lis)

	if err != nil {
		fmt.Printf("failed to serve: %v\n", err)
		return
	}
}
