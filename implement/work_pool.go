package implement

import (
	"simpleserver/iface"
	"simpleserver/utils"
	"time"
)

type WorkPool struct {
	// 常驻工作协程数量
	FormalWorkerCount int
	// 最大工作协程数量
	MaxWorkerCount int
	// 临时工，最长空闲时间秒
	TemporaryWorkerLive int

	// 当前工作goroutine数量
	CurWorkerCount  int

	// 请求处理路由
	server iface.IServer

	TaskQueue chan iface.IRequest
	AddWorker chan bool
	DelWorker chan bool
}

func (pool *WorkPool) SetServer(server iface.IServer) {
	pool.server = server
}

// 发送处理请求
func (pool *WorkPool) PostRequest(req iface.IRequest) {
	select {
	case pool.TaskQueue <- req:
		utils.Log.Info("request ", req)
		return
	default:
		utils.Log.Warn("队列满了")
		//return errors.New("队列满了")
	}
}

func (pool *WorkPool) SetWorkerCount(n int) {
	//panic("implement me")
	pool.FormalWorkerCount = n
}

func (pool *WorkPool) GetWorkerCount() int {
	//panic("implement me")
	return pool.FormalWorkerCount
}

func (pool *WorkPool) SetMaxWorkerCount(n int) {
	//panic("implement me")
	pool.MaxWorkerCount = n
}

func (pool *WorkPool) GetMaxWorkerCount() int {
	return pool.MaxWorkerCount
}

func (pool *WorkPool) SetTemporaryWorkerLive(seconds int) {
	pool.TemporaryWorkerLive = seconds
}

func (pool *WorkPool) GetTemporaryWorkerLive() int {
	return pool.TemporaryWorkerLive
}

// 获取请求路径，找到处理请求对象，并执行处理任务
func (pool *WorkPool) DoJob(req iface.IRequest) {
	handler := pool.server.GetRouter().GetHandler(req.GetPath())
	resp := utils.Create("Response").(iface.IResponse)
	handler.Process(req, resp)
}

// 启动工作池
func (pool *WorkPool) CreatePool() {
	for i:=0; i<pool.FormalWorkerCount; i++ {
		go pool.doWork()
	}
	pool.CurWorkerCount = pool.FormalWorkerCount
	pool.manageTempWorker()
	utils.Log.Info("CreatePool size:", pool.FormalWorkerCount)
}

// 管理临时工，控制总工作goroutine的数量
func  (pool *WorkPool) manageTempWorker() {
	for {
		select {
		case <- pool.AddWorker:
			pool.CurWorkerCount++
		case <- pool.DelWorker:
			pool.CurWorkerCount--
		}
	}
}
// 添加一个临时工
func (pool *WorkPool) CreateTempWorker(req iface.IRequest) {
	if pool.CurWorkerCount < pool.MaxWorkerCount {
		// 处理工作
		go pool.doTempWork()
		pool.AddWorker <- true
	}
}

// 临时工，空闲时间超过则自动退出
func (pool *WorkPool) doTempWork() {
	utils.Log.Info("创建一个临时工....")
	timer := time.NewTimer(time.Duration(pool.TemporaryWorkerLive) * time.Second)
	for {
		select {
		case req := <- pool.TaskQueue :
			pool.DoJob(req)
			timer.Reset(time.Duration(pool.TemporaryWorkerLive) * time.Second)
		case <- timer.C:
			utils.Log.Info("临时工空闲超时，退出。")
			pool.DelWorker <- true
			return
		}
	}
}

// 正式工，一直不停工作
func (pool *WorkPool) doWork() {
	for {
		select {
		case req := <- pool.TaskQueue :
			pool.DoJob(req)
		}
	}
}

func init(){
	utils.Log.Info("Register WorkerPool")
	utils.Register("WorkerPool", func() interface{} {
		return &WorkPool{
			FormalWorkerCount:10,
			MaxWorkerCount: 10000,
			TemporaryWorkerLive:300,
			TaskQueue: make(chan iface.IRequest, 1024),
			AddWorker: make(chan bool),
			DelWorker: make(chan bool),
		}
	})
}
