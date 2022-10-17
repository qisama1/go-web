package test

import (
	"fmt"
	"net"
	"testing"
	"time"
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
		_, err := conn.Write([]byte("Hello Zinx V0.1"))
		if err != nil {
			fmt.Println("write err", err)
			return
		}
		buf := make([]byte, 512)
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf err")
			return
		}
		fmt.Printf("server call back %s, cnt := %d \n", buf[:len], len)
		time.Sleep(500 * time.Millisecond)
	}
}
