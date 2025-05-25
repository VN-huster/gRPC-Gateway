package main

import (
    "context"
    "flag"
    "net/http"
    "fmt"
    "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
    "google.golang.org/grpc/grpclog"

    gw "github.com/VN-huster/gRPC-Gateway/protogen/golang/proto" // Update
)

var (
    // command-line options:
    // gRPC server endpoint
    grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:8080", "gRPC server endpoint")
)

func run() error {
    ctx := context.Background()
    ctx, cancel := context.WithCancel(ctx)
    defer cancel()

    // Register gRPC server endpoint
    // Note: Make sure the gRPC server is running properly and accessible
    mux := runtime.NewServeMux()
    opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
    err := gw.RegisterGreeterHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts) // Updated function name
    if err != nil {
        return err
    }

    // Start HTTP server (and proxy calls to gRPC server endpoint)
    fmt.Println("Listening on port 8081...\n")
    return http.ListenAndServe(":8081", mux)
}

func main() {
    flag.Parse()

    if err := run(); err != nil {
        grpclog.Fatal(err)
    }
}