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
	// AddRouter 添加路由
	AddRouter(msgId uint32, router IRouter)
	// GetConnectionManager 获取连接管理器
	GetConnectionManager() IConnectionManager
	RegistryInitMethod(initMethod func(conn IConnection))
	RegistryDestroyMethod(initMethod func(conn IConnection))
	CallInitMethod(conn IConnection)
	CallDestroyMethod(conn IConnection)
}
