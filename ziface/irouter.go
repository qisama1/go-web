package ziface

type IRouter interface {
	// PreHandle 处理业务之前的处理的钩子方法
	PreHandle(request IRequest)
	// Handler 处理conn业务的主方法
	Handler(request IRequest)
	// PostHandle 处理业务之后的钩子方法
	PostHandle(request IRequest)
}
