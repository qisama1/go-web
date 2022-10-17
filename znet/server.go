package znet

import (
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
			go func() {
				for {
					buf := make([]byte, 512)
					len, err := conn.Read(buf)
					if err != nil {
						fmt.Println("Read err", err)
						continue
					}
					if _, err := conn.Write(buf[:len]); err != nil {
						fmt.Println("write back err", err)
						continue
					}
				}
			}()
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

func NewServer(name string) ziface.IServer {
	return &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}
}
