package iface

type IConnmanage interface {
	// 添加
	NewConnection(conn interface{}) IConnection
	// 删除
	Remove(conn IConnection)
	// 获取当前连接数
	ConnCount() int

	// 保持Server指针
	SetServer(server IServer)

	// 发送处理请求
	PostRequest(req IRequest)

	// 启动管理协程
	RunManage()
}
