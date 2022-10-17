package znet

import (
	"net"
	"zinx/ziface"
)

// Connection /**
type Connection struct {
	// 当前连接的socket TCP套接字
	Conn *net.TCPConn
	// 连接的ID
	ConnID uint32
	// 当前连接的状态
	isClosed bool
	// 当前连接所绑定的处理业务的方法
	handleAPI ziface.HandleFunc
	// 告知连接已经停止的channel
	ExitChan chan bool
}

func (c Connection) Start() {
	//TODO implement me
	panic("implement me")
}

func (c Connection) Stop() {
	//TODO implement me
	panic("implement me")
}

func (c Connection) GetTcpConnection() *net.TCPConn {
	//TODO implement me
	panic("implement me")
}

func (c Connection) GetConnId() uint32 {
	//TODO implement me
	panic("implement me")
}

func (c Connection) GetRemoteAddr() net.Addr {
	//TODO implement me
	panic("implement me")
}

func (c Connection) Send(data []byte) error {
	//TODO implement me
	panic("implement me")
}

func NewConnection(conn *net.TCPConn, connId uint32, callback_api ziface.HandleFunc) *Connection {
	return &Connection{
		Conn:      conn,
		ConnID:    connId,
		handleAPI: callback_api,
		isClosed:  false,
		ExitChan:  make(chan bool, 1),
	}
}
