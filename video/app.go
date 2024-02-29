package main

import (
	"video/kit/consul"
	"video/kit/dbinit"
	"video/service"
)

func main() {
	dbinit.Init()
	consulIp, consulPort, registerIp, registerPort := consul.InitAddressConfig()
	go service.SecviceInit(registerIp, registerPort)
	consul.Register("video", consulIp, registerIp, registerPort, consulPort)
	// bug 健康检测接口端口与服务接口端口相同 ， 健康检测接口导致被占用，导致注册失败
}
