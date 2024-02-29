package consul

import "github.com/spf13/viper"

func InitAddressConfig() (string, int, string, int) {
	consulIp := viper.GetString("consul.hostAddr")
	consulPort := viper.GetInt("consul.hostPort")
	registerIp := viper.GetString("consul.registerAddr")
	registerPort := viper.GetInt("consul.registerPort")
	return consulIp, consulPort, registerIp, registerPort
}
