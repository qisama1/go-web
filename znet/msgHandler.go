package znet

import "zinx/ziface"

type MsgHandler struct {
	// 存放msgId，所对应的处理方法
	Apis map[uint32]ziface.IRouter
}

func NewMsgHandler() ziface.IMsgHandler {
	return &MsgHandler{
		Apis: make(map[uint32]ziface.IRouter),
	}
}

func (m MsgHandler) DoMsgHandler(request ziface.IRequest) {
	if router, ok := m.Apis[request.GetMsgID()]; ok {
		router.PreHandle(request)
		router.Handler(request)
		router.PostHandle(request)
	} else {
		router := &BaseRouter{}
		m.Apis[request.GetMsgID()] = router
		router.PreHandle(request)
		router.Handler(request)
		router.PostHandle(request)
	}
}

func (m MsgHandler) AddRouter(msgId uint32, router ziface.IRouter) {
	m.Apis[msgId] = router
}
