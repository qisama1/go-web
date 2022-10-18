package znet

import "zinx/ziface"

type Request struct {
	// 和客户端建立好的Connection
	conn ziface.IConnection
	// 客户端请求的数据
	data []byte
}

func (r Request) GetConnection() ziface.IConnection {
	return r.conn
}

func (r Request) GetData() []byte {
	return r.data
}

// NewRequest
// 这是一个接口
// /*
func NewRequest(conn ziface.IConnection, data []byte) *Request {
	return &Request{
		conn: conn,
		data: data,
	}
}
