package main

import (
	"toDoListDemo/conf"
	"toDoListDemo/routes"
)

func main() {
	conf.Init() //初始化，主要是（1）连接mysql数据库  （2）读取redis的配置信息
	g := routes.NewRouter()
	g.Run(conf.HttpPort)

}
