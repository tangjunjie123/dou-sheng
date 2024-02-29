package main

import (
	"relation/kit/consul"
	"relation/kit/dbinit"
	"relation/service"
)

func main() {
	dbinit.Init()
	consulIp, consulPort, registerIp, registerPort := consul.InitAddressConfig()
	go service.SecviceInit(registerIp, registerPort)
	consul.Register("relation", consulIp, registerIp, registerPort, consulPort)
	// bug 健康检测接口端口与服务接口端口相同 ， 健康检测接口导致被占用，导致注册失败
}
