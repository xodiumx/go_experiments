package main

import (
	"context"
	"log"
	"net/http"

	"gateway/example.com/project/gen/go/userpb"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := userpb.RegisterUserServiceHandlerFromEndpoint(ctx, mux, "localhost:50051", opts)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("gRPC-Gateway listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
