package iface


type IServer interface {
	// 初始化配置
	Init()
	// 启动服务
	Start()
	// 停止服务
	Stop()
	// 创建工作goroutine池
	CreateWorkPool()

	// 设置路由
	SetRouter(router IRouter)
	GetRouter() IRouter

	// 发送处理请求
	PostRequest(req IRequest)
}