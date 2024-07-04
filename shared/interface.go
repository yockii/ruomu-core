package shared

import (
	"github.com/hashicorp/go-plugin"
)

var Handshake = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "RUOMU_COMM",
	MagicCookieValue: "MagicSword",
}

type Communicate interface {
	// Initial 初始化，模块需要的全部参数
	Initial(map[string]string) error
	// InjectCall 注入点调用，code为模块注册注入点时的代码，value为请求的参数json；返回json
	InjectCall(code string, headers map[string][]string, value []byte) ([]byte, error)

	// 以下为grpc必须实现的接口
	mustEmbedUnimplementedCommunicateServer()
}

//  protoc --go_out=. --go-grpc_out=. shared/*.proto
