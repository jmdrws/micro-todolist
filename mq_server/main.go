package main

import (
	"mq_server/conf"
	"mq_server/service"
)

func main() {
	conf.Init()

	forever := make(chan bool)
	service.CreateTask()
	<-forever
}
