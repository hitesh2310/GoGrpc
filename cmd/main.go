package main

import (
	"context"
	"fmt"
	"log"
	pb "main/pkg/proto"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedMyServiceServer
}

func (s *server) SendData(ctx context.Context, in *pb.Person) (*pb.Response, error) {

	fmt.Println("Code ctrl reaches here..")
	fmt.Println(fmt.Sprintf("incoming request [%v]", in))

	return &pb.Response{Message: "status:\\\"Request accepted\\\""}, nil
}

func main() {
	fmt.Println("Hello World...")
	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	s := grpc.NewServer()
	pb.RegisterMyServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
