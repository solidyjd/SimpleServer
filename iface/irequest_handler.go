package iface

type IRequestHandler interface {
	Process(req IRequest, resp IResponse)
}
