package utils

import (
	"net"
	"simpleserver/iface"
)

type funCreateConn func(conn net.Conn) iface.IConnection
var (
	// 保存注册好的工厂信息
	factoryByName = make(map[string]func() interface{})
)
// 注册一个类生成工厂
func Register(name string, factory func() interface{}) {
	factoryByName[name] = factory
}

// 根据名称创建对应的类
func Create(name string) interface{} {
	if f, ok := factoryByName[name]; ok {
		return f()
	} else {
		//panic("name not found")
		return nil
	}
}
