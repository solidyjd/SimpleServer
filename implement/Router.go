package implement

import (
	"simpleserver/iface"
	"simpleserver/utils"
)

type Router struct {
	trees map[string]iface.IRequestHandler
}

func (this *Router) GetHandler(path string) iface.IRequestHandler {
	return this.trees[path]
}

func (this *Router) AddRouter(path string, handler iface.IRequestHandler) {
	this.trees[path] = handler
}


func init(){
	utils.Log.Info("Register Router")
	utils.Register("Router", func() interface{} {
		return &Router{
			trees:make(map[string]iface.IRequestHandler),
		}
	})
}