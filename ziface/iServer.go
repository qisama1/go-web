package ziface

// IServer 定义一个服务器接口
type IServer interface {
	// Init 初始化服务器
	Init()
	// Start 启动服务器
	Start()
	// Stop 停止服务器
	Stop()
	// Serve 运行服务器
	Serve()
}
