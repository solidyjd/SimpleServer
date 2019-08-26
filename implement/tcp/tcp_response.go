package tcp

import "simpleserver/utils"

type TcpResponse struct {
	Fields  map[string][]byte
}

func (resp *TcpResponse) SetField(name string, value []byte) {
	panic("implement me")
}

func (resp *TcpResponse) Package() []byte {
	panic("implement me")
}


func init()  {
	utils.Log.Info("Register TcpResponse")
	utils.Register("Response", func() interface{} {
		resp := &TcpResponse{
			Fields:make(map[string][]byte),
		}
		return resp
	})
}

