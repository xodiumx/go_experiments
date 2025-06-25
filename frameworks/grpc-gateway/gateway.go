package main

import (
	"context"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"

	"gateway/example.com/project/gen/go/userpb"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	handler := h2c.NewHandler(mux, &http2.Server{})

	err := userpb.RegisterUserServiceHandlerFromEndpoint(ctx, mux, "localhost:50051", opts)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("gRPC-Gateway listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
