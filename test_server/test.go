package main

import (
	"fmt"
	"simpleserver/implement"
	"simpleserver/implement/tcp"
)

type Iface interface {
	Do()
}
type base struct{
	id int
}
func (b base)Do(){
	fmt.Println("dododododo")
}

type sub struct{
	base
	name string
}

func create() Iface {
	//ss := p.(*sub)
	//var f sub
	//f = sub{name:"hello"}
	var b interface{}
	b = &sub{name:"nnnn"}

	b1 := &base{id:13343}
	//b1.(Iface) error
	var bt interface{} = b1
	c1 := bt.(Iface)  // 只有Interface才能这样

	fmt.Println(c1)

	bb := b.(*sub)
	bb.id = 3
	c := b.(Iface)
	fmt.Println(b)
	//c,ok := b.(iface)
	return c
}
func main(){
	server := tcp.TcpServer{Server: implement.Server{IpVersion: "tcp4", Ip: "0.0.0.0", Port: 6666}}
	server.Init()
	server.Start()
	select {}
}
