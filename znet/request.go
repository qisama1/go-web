package znet

import "zinx/ziface"

type Request struct {
	// 和客户端建立好的Connection
	conn ziface.IConnection
	// 客户端请求的数据
	data ziface.IMessage
}

func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.data.GetMessage()
}

func (r *Request) GetMsgID() uint32 {
	return r.data.GetMessageId()
}
