package shared

import (
	"context"
)

type GrpcClient struct {
	client CommunicateClient
}

func (m *GrpcClient) Initial(params map[string]string) error {
	_, err := m.client.Initial(context.Background(), &InitialRequest{
		Params: params,
	})
	return err
}

func (m *GrpcClient) InjectCall(code string, value []byte) ([]byte, error) {
	resp, err := m.client.InjectCall(context.Background(), &InjectCallRequest{
		Code:  code,
		Value: value,
	})
	if err != nil {
		return nil, err
	}
	return resp.Result, nil
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
