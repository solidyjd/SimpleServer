package tcp

import (
	"net"
	"simpleserver/iface"
	"simpleserver/utils"
)

type TcpConnManage struct {
	Server iface.IServer
	connections map[string]iface.IConnection //管理的连接信息
	addChain chan iface.IConnection		// 添加新连接通道
	removeChain chan iface.IConnection		// 删除连接通道
	//cleanChain chan bool          // 清空
}

func (connMgr *TcpConnManage) SetServer(server iface.IServer) {
	connMgr.Server = server
}

func (connMgr *TcpConnManage) PostRequest(req iface.IRequest) {
	connMgr.Server.PostRequest(req)
}

func (connMgr *TcpConnManage) NewConnection(conn interface{}) iface.IConnection {
	utils.Log.Debug("NewConnectin..........")
	newConn := utils.Create("Connection")
	if newConn != nil {
		tcp := newConn.(*TcpConnection)
		tcp.Conn = conn.(*net.TCPConn)
		tcp.Mgr = connMgr
		//tcp.ConId = utils.GenUid()
		connMgr.addChain <- newConn.(iface.IConnection)
		return newConn.(iface.IConnection)
	}
	return nil
}

func (connMgr *TcpConnManage) Remove(conn iface.IConnection) {
	connMgr.removeChain <- conn
}

func (connMgr *TcpConnManage) ConnCount() int {
	return len(connMgr.connections)
}

func (connMgr *TcpConnManage) RunManage() {
	for {
		select {
		case conn := <- connMgr.addChain:
			//将conn连接添加到ConnMananger中
			connMgr.connections[conn.GetId()] = conn
			utils.Log.Info("connection add to ConnManager: connId = ", conn.GetId())
		case conn := <- connMgr.removeChain:
			//删除连接信息
			delete(connMgr.connections, conn.GetId())
			utils.Log.Info("connection Remove ConnID=",conn.GetId())
		//case <- connMgr.cleanChain:
		//	//停止并删除全部的连接信息
		//	for connID, conn := range connMgr.connections {
		//		//停止
		//		conn.Stop()
		//		//删除
		//		delete(connMgr.connections,connID)
		//	}
		//	utils.Log.Info("Delete All Connections ", connMgr.ConnCount())
		}
	}
}

func init(){
	utils.Log.Info("Register ConnManage")
	utils.Register("ConnManage", func() interface{} {
		connMgr := &TcpConnManage{
			connections: make(map[string]iface.IConnection),
			addChain:    make(chan iface.IConnection, 100),
			removeChain: make(chan iface.IConnection, 100),
		}
		//go connMgr.RunManage()
		return connMgr
	})
}

