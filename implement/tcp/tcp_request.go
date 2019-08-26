package tcp

import (
	"simpleserver/iface"
	"simpleserver/utils"
)

type TcpRequest struct {
	Conn       iface.IConnection
	path       string
	RawData    []byte
	//Handler    iface.IRequestHandler
	PkgParser  iface.IParser
}

func (request *TcpRequest) SetConnection(conn iface.IConnection) {
	request.Conn = conn
}

func (request *TcpRequest) GetConnection() iface.IConnection {
	return request.Conn
}

func (request *TcpRequest) GetPath() string {
	return request.path
}

// 解包请求数据
func (request *TcpRequest) ParsePackage(raw []byte) {
	request.RawData = raw
	request.path = "*"
}

// 获取解包后的字段值
func (request *TcpRequest) Get(name string) []byte {
	return []byte(name)
}

//func (request *TcpRequest) SetHandler(handler iface.IRequestHandler) {
//	request.Handler = handler
//}

func (request *TcpRequest) SetPackageParser(parser iface.IParser) {
	request.PkgParser = parser
}

func init()  {
	utils.Log.Info("Register TcpRequest")
	utils.Register("Request", func() interface{} {
		req := &TcpRequest{}
		return req
	})
}