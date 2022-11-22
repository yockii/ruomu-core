package shared

import (
	"context"
)

type GrpcClient struct {
	client CommunicateClient
}

////////////////////////////////////////////////////

type GrpcServer struct {
	Impl Communicate
}

func (g *GrpcServer) Initial(ctx context.Context, request *InitialRequest) (*Empty, error) {
	return &Empty{}, g.Impl.Initial(request.Params)
}

func (g *GrpcServer) InjectCall(ctx context.Context, request *InjectCallRequest) (*InjectCallResponse, error) {
	v, err := g.Impl.InjectCall(request.Code, request.Value)
	return &InjectCallResponse{Result: v}, err
}
