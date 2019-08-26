package implement

import (
	"simpleserver/iface"
	"simpleserver/utils"
)

type Server struct {
	IpVersion string
	Ip		string
	Port    int
	ConnManage iface.IConnmanage
	WorkerPool iface.IWorkerPool
	Router iface.IRouter
}

func (server *Server) SetRouter(router iface.IRouter) {
	server.Router = router
}

func (server *Server) PostRequest(req iface.IRequest) {
	server.WorkerPool.PostRequest(req)
}

func (server *Server) Init() {
	//panic("implement me")
	utils.Log.Debug("Init server")
	if t := utils.Create("ConnManage"); t != nil {
		server.ConnManage = t.(iface.IConnmanage)
		server.ConnManage.SetServer(server)
		utils.Log.Info("Create ConnManage:", server.ConnManage)
	}

	//if t := utils.Create("RequestHandler"); t != nil {
	//	server.ReqHandler = t.(iface.IRequestHandler)
	//	utils.Log.Info("Create RequestHandler:", server.ReqHandler)
	//}

	if t := utils.Create("WorkerPool"); t != nil {
		server.WorkerPool = t.(iface.IWorkerPool)
		//server.WorkerPool.SetR
		utils.Log.Info("Create WorkerPool:", server.WorkerPool)
	}

}

func (server *Server) Start() {
	panic("implement me")
}

func (server *Server) Stop() {
	panic("implement me")
}

func (server *Server) CreateWorkPool() {
	server.WorkerPool.CreatePool()
}






