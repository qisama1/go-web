package utils

import (
	"encoding/json"
	"os"
	"zinx/ziface"
)

// GlobalObj 存储全局参数, 一些参数是由zinx.json进行配置的
type GlobalObj struct {
	// Server
	TcpServer ziface.IServer // 当前Zinx全局的Server对象
	Host      string         // 当前服务器监听的IP
	TcpPort   int            // 当前服务器监听的端口号
	Name      string         // 当前服务器的名称

	// Zinx
	Version        string // 当前zinx的版本号
	MaxConn        int    // 当前服务器主机允许的最大连接数
	MaxPackageSize uint32 // 当前一次传输允许的最大值
}

var GlobalConfig *GlobalObj

// 提供一个init方法，初始化当前的GlobalConfig
func init() {
	// 如果配置文件没有的话，这是默认配置
	GlobalConfig = &GlobalObj{
		Name:           "ZinxServerApp",
		Version:        "V0.4",
		TcpPort:        8999,
		Host:           "0.0.0.0",
		MaxConn:        1000,
		MaxPackageSize: 4096,
	}
	// 从配置文件中加载
	GlobalConfig.Reload()
}

// Reload 读取配置
func (g *GlobalObj) Reload() {
	data, err := os.ReadFile("../conf/zinx.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &g)
	if err != nil {
		panic(err)
	}
}
