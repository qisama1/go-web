package znet

import (
	"errors"
	"fmt"
	"sync"
	"zinx/ziface"
)

type ConnectionManager struct {
	connections map[uint32]ziface.IConnection // 管理的连接集合
	connLock    sync.RWMutex                  // 保护连接集合的读写锁
}

func NewConnectionManager() *ConnectionManager {
	return &ConnectionManager{
		connections: make(map[uint32]ziface.IConnection),
		connLock:    sync.RWMutex{},
	}
}

func (c ConnectionManager) Add(connection ziface.IConnection) {
	c.connLock.Lock()
	// 类似于对finally的处理
	defer c.connLock.Unlock()
	// 将conn加入到ConnManager中
	c.connections[connection.GetConnId()] = connection
	fmt.Println("connection add to ConnManager suc")
}

func (c ConnectionManager) GetConnection(connId uint32) (ziface.IConnection, error) {
	c.connLock.Lock()
	defer c.connLock.Unlock()
	if conn, ok := c.connections[connId]; ok {
		return conn, nil
	} else {
		return nil, errors.New("没有这个")
	}
}

func (c ConnectionManager) Remove(conn ziface.IConnection) {
	// 保护共享资源Map
	c.connLock.Lock()
	defer c.connLock.Unlock()
	delete(c.connections, conn.GetConnId())
}

func (c ConnectionManager) Len() int {
	c.connLock.RLock()
	defer c.connLock.RUnlock()
	return len(c.connections)
}

func (c ConnectionManager) Clear() {
	c.connLock.Lock()
	defer c.connLock.Unlock()
	for connId, conn := range c.connections {
		conn.Stop()
		delete(c.connections, connId)
	}
	fmt.Println("ClearAll")
}
