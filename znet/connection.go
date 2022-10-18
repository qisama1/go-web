package znet

import (
	"fmt"
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
	// 告知连接已经停止的channel
	ExitChan chan bool
	// 该链接处理的方法
	Router ziface.IRouter
}

// StartReader 连接的读业务方法
func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running ...")
	defer fmt.Println("connID = ", c.ConnID, "reader is exit")
	defer c.Stop()

	for {
		// 读取客户端的数据到buf中，目前最大是512字节
		buf := make([]byte, 512)
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("read err", err)
			continue
		}
		req := &Request{
			conn: c,
			data: buf[:cnt],
		}
		// 开个协程去处理
		go func() {
			c.Router.PreHandle(req)
			c.Router.Handler(req)
			c.Router.PostHandle(req)
		}()
	}
}

func (c *Connection) Start() {
	fmt.Println("conn start", c.Conn.RemoteAddr())
	// 启动当前连接的读数据业务
	go c.StartReader()
	// TODO 启动从当前连接写数据的业务
}

func (c *Connection) Stop() {
	if c.isClosed {
		return
	}
	// 关闭连接
	c.Conn.Close()
	c.isClosed = true
	// 回收资源
	close(c.ExitChan)
}

func (c *Connection) GetTcpConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnId() uint32 {
	return c.ConnID
}

func (c *Connection) GetRemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) Send(data []byte) error {
	//TODO implement me
	panic("implement me")
}

func NewConnection(conn *net.TCPConn, connId uint32, router ziface.IRouter) *Connection {
	return &Connection{
		Conn:     conn,
		ConnID:   connId,
		isClosed: false,
		ExitChan: make(chan bool, 1),
		Router:   router,
	}
}
