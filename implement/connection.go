package implement

import (
	"github.com/satori/go.uuid"
	"simpleserver/iface"
	"simpleserver/utils"
)

type Connection struct {
	uid string
}

func (conn *Connection) NewRequest(rawData []byte) iface.IRequest {
	panic("implement me")
}

func (conn *Connection) Read(b []byte) (int, error) {
	panic("implement me")
}

func (conn *Connection) Write(b []byte) (int, error) {
	panic("implement me")
}

func (conn *Connection) Close() error {
	panic("implement me")
}

func (conn *Connection) GetId() {
	//panic("implement me")
	u1,_ := uuid.NewV4()
	utils.Log.Infof("UUIDv4: %s\n", u1)
}

func (conn *Connection)HandleClient(){

}

