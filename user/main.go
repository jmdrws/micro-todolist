package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"user/conf"
	"user/core"
	"user/service"
)


func main() {
	//todo 数据库配置错误
	conf.Init()
	//注册etcd
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)
	//1、得到微服务实例
	microService := micro.NewService(
			micro.Name("rpcUserService"),
			micro.Address("127.0.0.1:8082"),
			micro.Registry(etcdReg),
	)
	//初始化
	microService.Init()
	//服务注册
	//将用户服务注册到etcd中
	_ = service.RegisterUserServiceHandler(microService.Server(), new(core.UserService))
	_ = microService.Run()

}