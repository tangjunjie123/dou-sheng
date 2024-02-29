package viper

import (
	"fmt"
	"github.com/spf13/viper"
)

func Viper_init() *viper.Viper {
	res := viper.GetViper()
	res.SetConfigFile("./kit/etc/config.yaml") // 指定配置文件路径
	//res.SetConfigName("config")                                             // 配置文件名称(无扩展名)
	res.SetConfigType("yaml") // 如果配置文件的名称中没有扩展名，则需要配置此项
	err := res.ReadInConfig() // 查找并读取配置文件
	if err != nil {
		fmt.Println(err)
	}
	return res
}
