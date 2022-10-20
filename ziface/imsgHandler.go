package ziface

type IMsgHandler interface {
	// DoMsgHandler 调度对应的router
	DoMsgHandler(request IRequest)
	// AddRouter 为消息添加具体的router
	AddRouter(msgId uint32, router IRouter)
}
