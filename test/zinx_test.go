package test

import (
	"fmt"
	"testing"
	"zinx/ziface"
	"zinx/znet"
)

type MyRouter struct {
	znet.BaseRouter
}

//// PreHandle 处理业务之前的处理的钩子方法
//func (myRouter *MyRouter) PreHandle(request ziface.IRequest) {
//	fmt.Println("Call PreHandler")
//	_, err := request.GetConnection().GetTcpConnection().Write([]byte("before ping...\n"))
//	if err != nil {
//		fmt.Println("Call PreHandle fail")
//	}
//}

// Handler 处理conn业务的主方法
func (myRouter *MyRouter) Handler(request ziface.IRequest) {
	fmt.Println("Handler xx 0.6")
	err := request.GetConnection().SendMsg(1, []byte("ping...\n"))
	//_, err := request.GetConnection().GetTcpConnection().Write([]byte("ping...\n"))
	if err != nil {
		fmt.Println("Call Handle fail")
	}
}

//// PostHandle 处理业务之后的钩子方法
//func (myRouter *MyRouter) PostHandle(request ziface.IRequest) {
//	fmt.Println("Call PostHandler")
//	_, err := request.GetConnection().GetTcpConnection().Write([]byte("after ping...\n"))
//	if err != nil {
//		fmt.Println("Call PostHandle fail")
//	}
//}

func TestZinx(t *testing.T) {
	// 1. 创建一个server
	s := znet.NewServer()
	// 2. 启动server
	s.AddRouter(0, &MyRouter{})
	s.Serve()
}
