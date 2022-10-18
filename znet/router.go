package znet

import "zinx/ziface"

// BaseRouter 这只是一个默认的路由，如果用户有自己的需要，可以自己重写
type BaseRouter struct {
}

func (b BaseRouter) PreHandle(request ziface.IRequest) {

}

func (b BaseRouter) Handler(request ziface.IRequest) {

}

func (b BaseRouter) PostHandle(request ziface.IRequest) {

}
