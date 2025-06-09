package main

import (
	"context"
	"log"
	"net"

	pb "gateway/example.com/project/gen/go/userpb"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	log.Printf("Received GetUser request for ID=%s", req.Id)
	log.Printf("Request %v", req.GetId())
	return &pb.UserResponse{
		Id:    req.Id,
		Name:  "John Doe",
		Email: "johndoe@gmail.com",
	}, nil
}

func (s *server) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	log.Printf("Current page: %v Limit: %v", req.Page, req.Limit)
	return &pb.ListUsersResponse{
		Users: []*pb.UserResponse{
			{Id: "1", Name: "Alice", Email: "alice@example.com"},
			{Id: "2", Name: "Bob", Email: "bob@example.com"},
		},
		Total: 128,
	}, nil
}

func (s *server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	log.Printf("Username: %v Email: %v", req.Name, req.Email)
	return &pb.UserResponse{
		Id:    "1223",
		Name:  req.Name,
		Email: req.Email,
	}, nil
}

// TODO: update
// TODO: delete

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	log.Println("gRPC server listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
