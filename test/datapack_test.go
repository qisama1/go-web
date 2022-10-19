package test

import (
	"fmt"
	"io"
	"net"
	"testing"
	"zinx/znet"
)

// 只是负责测试datapack拆包 封包的单元测试
func TestDataPack(t *testing.T) {
	/**
	模拟的服务器
	*/
	listener, err := net.Listen("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("server listen fail err", err)
		return
	}

	// 创建一个go负责处理业务
	// 2. 从客户端读取数据，拆包处理
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("server accept error", err)
			}
			go func(conn net.Conn) {
				// 处理客户端的请求
				// ---> 拆包的过程 <---
				dp := znet.NewDataPack()
				// 1. 第一次从conn读，根据head中的dataLen再读取data的内容
				for {
					headData := make([]byte, dp.GetHeadLen())
					_, err := io.ReadFull(conn, headData)
					if err != nil {
						fmt.Println("read head err")
						return
					}
					msgHead, err := dp.Unpack(headData)
					if err != nil {
						fmt.Println("server unpack err", err)
						return
					}
					if msgHead.GetMessageLength() > 0 {
						// 如果有数据再去读
						msg := msgHead.(*znet.Message)
						msg.Data = make([]byte, msg.GetMessageLength())

						// 根据datalen的长度从io流中读取
						_, err := io.ReadFull(conn, msg.Data)
						if err != nil {
							fmt.Println("server unpack data err", err)
							return
						}

						// 完整的消息处理流程
						fmt.Println("-> msgId", msg.Id, ", dataLen:", msg.DataLen, "data=", string(msg.Data))
					}
				}
			}(conn)
		}
	}()

	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client dial err", err)
		return
	}

	// 创建一个封包对象 dp
	dp := znet.NewDataPack()

	// 模拟粘包过程
	msg1 := &znet.Message{
		Id:      1,
		DataLen: 5,
		Data:    []byte{'h', 'e', 'l', 'l', 'o'},
	}
	sendData1, err := dp.Pack(msg1)
	if err != nil {
		fmt.Println("pack data err", err)
		return
	}
	msg2 := &znet.Message{
		Id:      1,
		DataLen: 6,
		Data:    []byte{'e', 'n', 'l', 'n', 'o', 'b'},
	}
	sendData2, err := dp.Pack(msg2)
	if err != nil {
		fmt.Println("pack data err", err)
		return
	}
	sendData := append(sendData1, sendData2...)
	conn.Write(sendData)

	select {}
}
