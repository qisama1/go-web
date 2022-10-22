package ziface

type IConnectionManager interface {
	// Add 添加连接
	Add(connection IConnection)
	// GetConnection 获取连接
	GetConnection(connId uint32) IConnection
	// Remove 删除连接
	Remove(conn IConnection)
	// Len 获取当前连接总数
	Len() int
	// Clear 清理所有连接
	Clear()
}
