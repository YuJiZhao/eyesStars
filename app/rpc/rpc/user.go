package rpc

import (
	"eyesStars/app/rpc/generate/userThrift"
	"eyesStars/global"
	"github.com/apache/thrift/lib/go/thrift"
	"net"
	"strconv"
)

/**
 * @author eyesYeager
 * @date 2023/1/29 16:33
 */

func User() (err error, client *userThrift.UserServiceClient, transport thrift.TTransport) {
	// 与服务端建立连接
	socket, err := thrift.NewTSocket(net.JoinHostPort(
		global.Config.Program.ThriftIp,
		strconv.Itoa(global.Config.Program.ThriftPort),
	))
	if err != nil {
		global.Log.Error("userThrift地址连接出错！" + err.Error())
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
		global.Log.Error("userThrift open错误！连接地址：" + global.Config.Program.ThriftIp + ":" + strconv.Itoa(global.Config.Program.ThriftPort) + " ;" + err.Error())
		return err, client, transport
	}

	// 使用User服务
	userProtocol := thrift.NewTMultiplexedProtocol(protocol, "UserService")
	c := thrift.NewTStandardClient(userProtocol, userProtocol)
	client = userThrift.NewUserServiceClient(c)
	return err, client, transport
}
