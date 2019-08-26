package iface

type IRequest interface {
	//// 设置请求报文原始数据
	//SetRawData(raw []byte)
	// 设置请求报文处理器
	//SetHandler(handler IRequestHandler)

	// 连接
	SetConnection(conn IConnection)
	GetConnection() IConnection

	// 设置请求包解析器
	SetPackageParser(parser IParser)

	// 解析请求包
	ParsePackage(raw []byte)

	// 获取请求路径
	GetPath() string

	// 取解析后的字段
	Get(name string) []byte
}