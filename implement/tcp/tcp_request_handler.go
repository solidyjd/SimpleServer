package tcp

import (
	"simpleserver/iface"
	"simpleserver/utils"
)

type TcpRequestHandler struct {

}

func (handler *TcpRequestHandler) Process(req iface.IRequest, resp iface.IResponse) {
	conn := req.GetConnection()
	tcp := conn.(*TcpConnection)
	tcp.Conn.Write([]byte("0x00023132"))
}

func init()  {
	utils.Log.Info("Register TcpRequestHandler")
	utils.Register("RequestHandler", func() interface{} {
		handler := &TcpRequestHandler{}
		return handler
	})
}

