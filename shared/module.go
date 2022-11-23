package shared

import "github.com/hashicorp/go-plugin"

func ModuleServe(moduleName string, impl Communicate) {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: Handshake,
		Plugins: map[string]plugin.Plugin{
			moduleName: &CommunicatePlugin{Impl: impl},
		},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
