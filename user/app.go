package main

import (
	"user/kit/consul"
	"user/kit/dbinit"
	"user/service"
)

func main() {
	dbinit.Init()
	consulIp, consulPort, registerIp, registerPort := consul.InitAddressConfig()
	go service.SecviceInit(registerIp, registerPort)
	consul.Register("user", consulIp, registerIp, registerPort, consulPort)
}
