package tcp

import (
	"fmt"
	"net"
	"simpleserver/implement"
	"simpleserver/utils"
)

type TcpServer struct {
	implement.Server
}

func (server *TcpServer) Start() {
	server.CreateWorkPool()
	go server.ConnManage.RunManage()

	go func() {
		// 获取一个TCP的Addr
		addr, err := net.ResolveTCPAddr(server.IpVersion, fmt.Sprintf("%s:%d", server.Ip, server.Port))
		if err != nil {
			utils.Log.Info("ResolveTCPAddr addr err: ", err)
			panic(err)
		}

		// 监听服务器地址
		listen, err:= net.ListenTCP(server.IpVersion, addr)
		if err != nil {
			utils.Log.Info("ListenTCP err:", err)
			panic(err)
		}

		for {
			// 等待客户端建立连接请求
			conn, err := listen.AcceptTCP()
			if err != nil {
				utils.Log.Info("Accept err ", err)
				continue
			}
			utils.Log.Info("conn remote addr = ", conn.RemoteAddr().String())
			client := server.ConnManage.NewConnection(conn)
			go client.HandleClient()
		}
	}()
	utils.Log.Infof("listening at IP: %s, Port %d", server.Ip, server.Port)
}
