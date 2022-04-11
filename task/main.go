package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"task/conf"
	"task/core"
	"task/service"
)


func main() {
	conf.Init()
	//注册etcd
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)
	//1、得到微服务实例
	microService := micro.NewService(
			micro.Name("rpcTaskService"),		//微服务名字
			micro.Address("127.0.0.1:8082"),
			micro.Registry(etcdReg),				//etcd 注册件
	)
	//结构命令行参数，初始化
	microService.Init()
	//服务注册
	//将用户服务注册到etcd中
	_ = service.RegisterTaskServiceHandler(microService.Server(), new(core.TaskService))
	_ = microService.Run()

}