package ziface

import "net"

// IConnection 定义连接模块的接口
type IConnection interface {
	// Start 启动连接，让当前的连接准备开始工作
	Start()
	// Stop 停止连接，结束当前连接的工作
	Stop()
	// GetTcpConnection 获取当前连接所绑定的socket conn
	GetTcpConnection() *net.TCPConn
	// GetConnId 获取当前连接模块的连接ID
	GetConnId() uint32
	// GetRemoteAddr 获取远程客户端的TCP状态 IP,PORT
	GetRemoteAddr() net.Addr
	// SendMsg 发送数据给远程的客户端
	SendMsg(msgId uint32, data []byte) error
}

// HandleFunc 定义一个处理连接业务的方法
type HandleFunc func(*net.TCPConn, []byte, int) error
