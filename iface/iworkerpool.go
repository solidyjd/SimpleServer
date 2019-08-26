package iface

type IWorkerPool interface {
	// 设置常驻工作协程数量
	SetWorkerCount(n int)
	GetWorkerCount() int

	// 设置最大工作协程数量
	SetMaxWorkerCount(n int)
	GetMaxWorkerCount() int

	// 设置临时工作协程最长空闲时间秒
	SetTemporaryWorkerLive(seconds int)
	GetTemporaryWorkerLive() int

	// 创建工作协程池
	CreatePool()

	// 设置请求处理路由
	SetServer(server IServer)

	// 请求处理
	PostRequest(req IRequest)
}