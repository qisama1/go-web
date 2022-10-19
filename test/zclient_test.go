package test

import (
	"fmt"
	"io"
	"net"
	"testing"
	"time"
	"zinx/znet"
)

/*
模拟客户端
*/
func TestClient(t *testing.T) {
	fmt.Println("client start...")
	time.Sleep(1 * time.Second)
	// 1.直接连接远程服务器得到一个conn
	// TODO 配置文件写活
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("conn err", err)
		return
	}
	// 2.连接调用write去写数据
	for {
		fmt.Println("client...")
		dp := znet.NewDataPack()
		binaryMsg, err := dp.Pack(&znet.Message{
			Id:      0,
			DataLen: uint32(len("ZinxV0.5 come")),
			Data:    []byte("ZinxV0.5 come"),
		})
		if err != nil {
			fmt.Println("pack err", err)
			return
		}
		_, err = conn.Write(binaryMsg)
		if err != nil {
			return
		}
		// 服务器给我们回复一个message， msgid为1
		msgHead := make([]byte, dp.GetHeadLen())
		// 读取客户端的Msg Head -二进制流 8个字节
		_, err = io.ReadFull(conn, msgHead)
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
			if _, err := io.ReadFull(conn, data); err != nil {
				fmt.Println("read msg data error", err)
				break
			}
		}
		fmt.Println("rev message ", string(data))

		time.Sleep(500 * time.Millisecond)
	}
}
