package ziface

type IMsgHandler interface {
	// DoMsgHandler 调度对应的router
	DoMsgHandler(request IRequest)
	// AddRouter 为消息添加具体的router
	AddRouter(msgId uint32, router IRouter)
	// StartWorkerPool 启动Worker工作池
	StartWorkerPool()
	// AddRequest 往对应的worker线程中添加req
	AddRequest(Idx uint32, req IRequest)
}
