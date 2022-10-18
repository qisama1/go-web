package znet

import (
	"errors"
	"fmt"
	"net"
	"zinx/ziface"
)

// Server 实现了iServer
type Server struct {
	// 服务的名称
	Name string
	// 绑定的IP版本
	IPVersion string
	// 监听的IP
	IP string
	// 监听的port
	Port int
	// Router, 注册的连接对应的处理业务
	Router ziface.IRouter
}

// CallBackToClient 定义当前客户端所绑定的handler api，目前是写死的，后面应该是留有接口，让用户提供
func CallBackToClient(conn *net.TCPConn, data []byte, len int) error {
	// 回写的业务
	fmt.Println("[Conn handler] CallBackToClient")
	if _, err := conn.Write(data[:len]); err != nil {
		fmt.Println("CallBackToClient write err", err)
		return errors.New("CallBackToClient err")
	}
	return nil
}

func (server *Server) Start() {
	fmt.Printf("[Start] Server listen at IP:%s, Port:%d\n", server.IP, server.Port)
	go func() {
		// 1. 获取一个TCP的Addr，获取一个套接字
		addr, err := net.ResolveTCPAddr(server.IPVersion, fmt.Sprintf("%s:%d", server.IP, server.Port))
		if err != nil {
			fmt.Println("resolve tcp addr error : ", err)
			return
		}
		// 2. 监听服务器的地址
		listener, err := net.ListenTCP(server.IPVersion, addr)
		if err != nil {
			fmt.Println("listen ", server.IPVersion, "err", err)
			return
		}
		fmt.Println("start Zinx server success", server.Name, " Listening")
		var cid uint32
		cid = 0
		// 3. 阻塞等待客户端进行连接，处理客户端连接业务
		for {
			// 如果没有客户端，会一直阻塞，有客户端的话就会阻塞返回
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}
			// 对客户端的连接进行处理，做一些业务，做一个回写的业务
			// 最大512字节的长度
			dealConn := NewConnection(conn, cid, server.Router)
			go dealConn.Start()
			cid++
		}
	}()

}

func (server *Server) Stop() {
	// TODO 将服务器的资源关闭
}

func (server *Server) Serve() {
	// 启动server的服务功能
	server.Start()

	// TODO：做一些启动服务器之后的额外业务

	// 阻塞状态
	select {}
}

func (server *Server) Init() {

}

func (server *Server) AddRouter(router ziface.IRouter) {
	server.Router = router
	fmt.Println("Add Router down!")
}

func NewServer(name string) ziface.IServer {
	return &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
		Router:    nil,
	}
}
