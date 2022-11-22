package shared

import (
	"context"

	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

type CommunicatePlugin struct {
	plugin.Plugin
	Impl Communicate
}

func (p *CommunicatePlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	RegisterCommunicateServer(s, &GrpcServer{Impl: p.Impl})
	return nil
}

func (p *CommunicatePlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GrpcClient{client: NewCommunicateClient(c)}, nil
}
