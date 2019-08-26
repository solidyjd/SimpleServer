package iface

type IRouter interface {
	GetHandler(path string) IRequestHandler
	AddRouter(path string, handler IRequestHandler)
}
