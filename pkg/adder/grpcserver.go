package adder

import (
	context "context"
	"github.com/ayupov-ayaz/grpc/pkg/api"
)

// GRPCServer implementation
type GRPCServer struct{}

func NewGRPCServer() *GRPCServer {
	return &GRPCServer{}
}

func (G GRPCServer) Add(ctx context.Context, request *api.AddRequest) (*api.AddResponse, error) {
	return &api.AddResponse{Result: request.GetX() + request.GetY()}, nil
}
