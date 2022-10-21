package handler

import (
	"context"

	healthv1 "github.com/linhbkhn95/rpc-service/go/health/v1"
)

type HealthServer struct {
	healthv1.UnimplementedHealthServiceServer
}

func NewHealthServer() healthv1.HealthServiceServer {
	return &HealthServer{}
}

//TODO: implement methods of this service.

// Ready ...
func (s HealthServer) Ready(ctx context.Context, req *healthv1.ReadyRequest) (*healthv1.ReadyResponse, error) {
	return &healthv1.ReadyResponse{}, nil
}
