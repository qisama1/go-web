package znet

import (
	"fmt"
	"math/rand"
	"zinx/utils"
	"zinx/ziface"
)

type MsgHandler struct {
	// 存放msgId，所对应的处理方法
	Apis map[uint32]ziface.IRouter
	// 负责Worker取任务的消息队列
	TaskQueue []chan ziface.IRequest
	// 业务工作池的Worker数量
	WorkerPoolSize uint32
}

func NewMsgHandler() ziface.IMsgHandler {
	return &MsgHandler{
		Apis:           make(map[uint32]ziface.IRouter),
		WorkerPoolSize: utils.GlobalConfig.WorkPoolSize,
		TaskQueue:      make([]chan ziface.IRequest, utils.GlobalConfig.WorkPoolSize),
	}
}

func (m *MsgHandler) DoMsgHandler(request ziface.IRequest) {
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

func (m *MsgHandler) AddRouter(msgId uint32, router ziface.IRouter) {
	m.Apis[msgId] = router
}

// StartWorkerPool 启动一个worker工作池
func (m *MsgHandler) StartWorkerPool() {
	// 根据workerPoolSize分别开启worker
	for i := 0; i < int(m.WorkerPoolSize); i++ {
		// 一个worker被启动
		// 给每个队列对应的channel初始化
		m.TaskQueue[i] = make(chan ziface.IRequest, utils.GlobalConfig.MaxWorkTaskLen)
		// 启动当前的worker
		go m.startWorker(i, m.TaskQueue[i])
	}
}

// startWorker 添加一个worker到工作池
func (m *MsgHandler) startWorker(idx int, reqChan chan ziface.IRequest) {
	fmt.Printf("worker %d 在工作\n", idx)
	for {
		select {
		case req := <-reqChan:
			fmt.Println(idx, "处理请求")
			m.DoMsgHandler(req)
		}
	}
}

func (m *MsgHandler) AddRequest(Idx uint32, req ziface.IRequest) {
	m.TaskQueue[m.GetWorkerId(Idx)] <- req
}

func (m *MsgHandler) GetWorkerId(reqId uint32) int {
	return rand.Intn(902423) % int(m.WorkerPoolSize)
}
