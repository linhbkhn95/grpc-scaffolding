package handler

import (
	"context"

	examplev1 "github.com/linhbkhn95/rpc-service/go/example/v1"
)

type ExampleServer struct {
	examplev1.UnimplementedExampleServiceServer
}

func NewExampleServer() examplev1.ExampleServiceServer {
	return &ExampleServer{}
}

//TODO: implement methods of this service.

// SayHello will send hello term to server.
func (s ExampleServer) SayHello(ctx context.Context, req *examplev1.SayHelloRequest) (*examplev1.SayHelloResponse, error) {
	return &examplev1.SayHelloResponse{}, nil
}

// SayGoodbye will send goodbye term to server.
func (s ExampleServer) SayGoodbye(ctx context.Context, req *examplev1.SayGoodbyeRequest) (*examplev1.SayGoodbyeResponse, error) {
	return &examplev1.SayGoodbyeResponse{}, nil
}
