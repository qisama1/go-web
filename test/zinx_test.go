package test

import (
	"testing"
	"zinx/znet"
)

func TestZinx(t *testing.T) {
	// 1. 创建一个server
	s := znet.NewServer("zinx")
	// 2. 启动server
	s.Serve()
}
