package znet

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

}

func (server *Server) Stop() {

}

func (server *Server) Serve() {

}
