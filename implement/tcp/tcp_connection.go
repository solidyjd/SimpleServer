package tcp

import (
	"errors"
	"io"
	"net"
	"simpleserver/iface"
	"simpleserver/utils"
	"time"
)

type TcpConnection struct {
	ConId string
	Mgr iface.IConnmanage
	Conn *net.TCPConn
	//Request iface.IRequest
}

func (tcp *TcpConnection) PostRequest(req iface.IRequest) {
	tcp.Mgr.PostRequest(req)
}

func (tcp *TcpConnection) GetId() string {
	//u1,_ := uuid.NewV4()
	//tcp.ConId = u1.String()
	//utils.Log.Infof("UUIDv4: %s\n", tcp.ConId)
	return tcp.ConId
}



func (tcp *TcpConnection) HandleClient() {
	defer tcp.Mgr.Remove(tcp)
	raw := make([]byte, 4096)
	dataLen, err := tcp.Read(raw)
	if err != nil {
		tcp.Close()
	}

	request := tcp.NewRequest(raw[:dataLen])
	utils.Log.Debug("recv:", string(raw[:dataLen]))
	tcp.PostRequest(request)
}

func (tcp *TcpConnection) NewRequest(rawData []byte) iface.IRequest{
	request := utils.Create("Request").(iface.IRequest)
	request.SetRawData(rawData)
	return request
}

func (tcp *TcpConnection) Read(b []byte) (int, error) {
	tcp.Conn.SetReadDeadline(time.Now().Add(20 * time.Second)) // 设置20s超时
	header := make([]byte, 2) // 报文开头2字节长度字段
	for {
		readLen, err := tcp.Conn.Read(header)
		if err != nil {
			utils.Log.Info(err)
			return 0, err
		}

		if readLen == 0 {
			utils.Log.Info("客户端已主动关闭连接！")
			return 0, errors.New("EOF")
		}

		dataLen := int(uint8(header[0])*255) + int(header[1])
		utils.Log.Debug("Request data len:", dataLen)
		rawData := make([]byte, dataLen)
		if _, err := io.ReadFull(tcp.Conn, rawData); err != nil{
			utils.Log.Error("接收请求报文数据出错")
			return 0, err
		}
		copy(b, rawData)
		return dataLen, nil
	}
}

func (tcp *TcpConnection) Write(b []byte) (int, error) {
	panic("implement me")
}

func (tcp *TcpConnection) Close() error {
	//panic("implement me")
	tcp.Conn.Close()
	tcp.Mgr.Remove(tcp)
	return nil
}

func init(){
	utils.Log.Info("Register TcpConnection")
	utils.Register("Connection", func() interface{} {
		connMgr := &TcpConnection{
			ConId: utils.GenUid(),
			Conn: nil,
		}
		//connMgr.run()
		return connMgr
	})
}
