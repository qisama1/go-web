package znet

import (
	"errors"
	"fmt"
	"net"
	"zinx/utils"
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
	// MsgHandler,其中集成了不同msg的对应的router
	MsgHandler ziface.IMsgHandler
	// 连接管理器
	ConnectionManager ziface.IConnectionManager
	// initMethod
	InitMethod func(conn ziface.IConnection)
	// destroyMethod
	DestroyMethod func(conn ziface.IConnection)
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
	fmt.Printf("[Start] Server%s listen at IP:%s, Port:%d\n", server.Name, server.IP, server.Port)
	go func() {
		// 0. 开启消息队列及Worker工作池
		server.MsgHandler.StartWorkerPool()

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
			// 新建连接的时候，考虑是不是超过了最大个数，如果超出了最大连接数，关闭这次连接
			fmt.Println(server.ConnectionManager.Len(), " 有多少个连接现在")
			if server.ConnectionManager.Len() >= utils.GlobalConfig.MaxConn {
				// TODO 告知用户
				fmt.Println("conn too much")
				conn.Close()
				continue
			}
			// 对客户端的连接进行处理，做一些业务，做一个回写的业务
			// 最大512字节的长度
			dealConn := NewConnection(server, conn, cid, server.MsgHandler)
			go dealConn.Start()
			cid++
		}
	}()

}

func (server *Server) Stop() {
	// TODO 将服务器的资源关闭
	fmt.Println("[STOP]")
	server.ConnectionManager.Clear()
}

func (server *Server) Serve() {
	// 启动server的服务功能
	server.Start()

	// TODO：做一些启动服务器之后的额外业务

	// 阻塞状态
	select {}
}

func (server *Server) AddRouter(msgID uint32, router ziface.IRouter) {
	server.MsgHandler.AddRouter(msgID, router)
}

func (server *Server) Init() {

}

func (server *Server) GetConnectionManager() ziface.IConnectionManager {
	return server.ConnectionManager
}

func NewServer() ziface.IServer {
	return &Server{
		Name:              utils.GlobalConfig.Name,
		IPVersion:         "tcp4",
		IP:                utils.GlobalConfig.Host,
		Port:              utils.GlobalConfig.TcpPort,
		MsgHandler:        NewMsgHandler(),
		ConnectionManager: NewConnectionManager(),
	}
}

func (server *Server) RegistryInitMethod(initMethod func(conn ziface.IConnection)) {
	server.InitMethod = initMethod
}

func (server *Server) RegistryDestroyMethod(destroyMethod func(conn ziface.IConnection)) {
	server.DestroyMethod = destroyMethod
}

func (server *Server) CallInitMethod(conn ziface.IConnection) {
	if server.InitMethod != nil {
		server.InitMethod(conn)
	}
}

func (server *Server) CallDestroyMethod(conn ziface.IConnection) {
	if server.DestroyMethod != nil {
		server.DestroyMethod(conn)
	}
}
