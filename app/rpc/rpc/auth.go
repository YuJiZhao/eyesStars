package rpc

import (
	"eyesStars/app/rpc/generate/authThrift"
	"eyesStars/global"
	"github.com/apache/thrift/lib/go/thrift"
	"net"
	"strconv"
)

/**
 * @author eyesYeager
 * @date 2023/1/29 16:33
 */

// Auth 鉴权接口调用
func Auth() (err error, client *authThrift.AuthServiceClient, transport thrift.TTransport) {
	// 与服务端建立连接
	socket, err := thrift.NewTSocket(net.JoinHostPort(
		global.Config.Program.ThriftIp,
		strconv.Itoa(global.Config.Program.ThriftPort),
	))
	if err != nil {
		global.Log.Error("authThrift地址连接出错！" + err.Error())
		return err, client, transport
	}

	// 创建协议工厂
	protocolFactory := thrift.NewTCompactProtocolFactory()
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())

	// 获取传输方式与传输协议
	transport, _ = transportFactory.GetTransport(socket)
	protocol := protocolFactory.GetProtocol(transport)

	// 为通信开启传输
	if err = transport.Open(); err != nil {
		global.Log.Error("authThrift open错误！连接地址：" + global.Config.Program.ThriftIp + ":" + strconv.Itoa(global.Config.Program.ThriftPort) + " ;" + err.Error())
		return err, client, transport
	}

	// 使用Auth服务
	authProtocol := thrift.NewTMultiplexedProtocol(protocol, "AuthService")
	c := thrift.NewTStandardClient(authProtocol, authProtocol)
	client = authThrift.NewAuthServiceClient(c)
	return err, client, transport
}
