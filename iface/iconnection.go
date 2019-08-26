package iface

type IConnection interface {
	// 当新的连接建立后，开始处理连接工作
	HandleClient()
	// 用接收到达数据创建一个Request
	NewRequest(rawData []byte) IRequest

	// 发送处理请求
	PostRequest(req IRequest)

	// 接收数据
	Read(b []byte) (int, error)
	// 发送数据
	Write(b []byte) (int, error)
	// 关闭
	Close() error

	// id
	GetId() string
}