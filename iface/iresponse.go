package iface

type IResponse interface {
	SetField(name string, value []byte)
	Package() []byte
}
