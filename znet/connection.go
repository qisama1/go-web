package znet

import (
	"errors"
	"fmt"
	"io"
	"net"
	"zinx/utils"
	"zinx/ziface"
)

// Connection /**
type Connection struct {
	// 当前连接的socket TCP套接字
	Conn *net.TCPConn
	// 连接的ID
	ConnID uint32
	// 当前连接的状态
	isClosed bool
	// 告知连接已经停止的channel
	ExitChan chan bool
	// 该链接处理的方法
	MsgHandler ziface.IMsgHandler
	// 进行通信的channel
	MsgChannel chan []byte
}

// StartReader 连接的读业务方法
func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running ...")
	defer fmt.Println("connID = ", c.ConnID, "reader is exit")
	defer c.Stop()

	for {
		// 读取客户端的数据到buf中，目前最大是512字节
		//buf := make([]byte, utils.GlobalConfig.MaxPackageSize)
		//cnt, err := c.Conn.Read(buf)
		//if err != nil {
		//	fmt.Println("read err", err)
		//	continue
		//}
		dp := NewDataPack()
		msgHead := make([]byte, dp.GetHeadLen())

		// 读取客户端的Msg Head -二进制流 8个字节
		_, err := io.ReadFull(c.GetTcpConnection(), msgHead)
		if err != nil {
			fmt.Println("read msg head error", err)
			break
		}
		// 拆包，得到msgID和msgDataLen放在msg对象消息中
		msg, err := dp.Unpack(msgHead)
		if err != nil {
			fmt.Println("unpack err", err)
			break
		}
		var data []byte
		if msg.GetMessageLength() > 0 {
			data = make([]byte, msg.GetMessageLength())
			if _, err := io.ReadFull(c.GetTcpConnection(), data); err != nil {
				fmt.Println("read msg data error", err)
				break
			}
		}

		msg.SetMessage(data)
		// 得到真正的数据
		req := &Request{
			conn: c,
			data: msg,
		}
		// 开个协程去处理
		// 把消息交给工作池
		if utils.GlobalConfig.WorkPoolSize > 0 {
			go c.MsgHandler.AddRequest(req.conn.GetConnId(), req)
		} else {
			// 如果没开启线程工作池，就单独去处理咯
			c.MsgHandler.DoMsgHandler(req)
		}

	}
}

// SendMsg 提供SendMsg方法，将要给客户端发送的数据，先进行封包，再发送
func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	if c.isClosed {
		return errors.New("connection closed when send message")
	}
	// 将Data封包 len/id/data
	dp := DataPack{}
	msg, err := dp.Pack(&Message{
		Id:      msgId,
		DataLen: uint32(len(data)),
		Data:    data,
	})
	if err != nil {
		fmt.Println("pack err", err)
		return err
	}
	// 将数据发送给客户端
	c.MsgChannel <- msg
	return nil
}

func (c *Connection) StartWriter() {
	fmt.Println("[writer...]")
	defer fmt.Println(c.GetRemoteAddr().String(), "[conn Writer exit!]")

	// 等待消息进来
	for {
		select {
		case data := <-c.MsgChannel:
			// 有数据来了
			if _, err := c.Conn.Write(data); err != nil {
				fmt.Println("send data err", err)
				return
			}
		case <-c.ExitChan:
			return
		}
	}
}

func (c *Connection) Start() {
	fmt.Println("conn start", c.Conn.RemoteAddr())
	// 启动当前连接的读数据业务
	go c.StartReader()
	// 启动当前连接去进行写业务
	go c.StartWriter()
	// TODO 启动从当前连接写数据的业务

}

func (c *Connection) Stop() {
	if c.isClosed {
		return
	}
	// 关闭连接
	c.Conn.Close()
	c.isClosed = true
	// 回收资源
	c.ExitChan <- true
	close(c.ExitChan)
	close(c.MsgChannel)
}

func (c *Connection) GetTcpConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnId() uint32 {
	return c.ConnID
}

func (c *Connection) GetRemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func NewConnection(conn *net.TCPConn, connId uint32, router ziface.IMsgHandler) *Connection {
	return &Connection{
		Conn:       conn,
		ConnID:     connId,
		isClosed:   false,
		ExitChan:   make(chan bool, 1),
		MsgHandler: router,
		MsgChannel: make(chan []byte),
	}
}
