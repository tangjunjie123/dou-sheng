package dbinit

import (
	mysqlDb "relation/kit/mysql"
	"relation/kit/redis"
	"relation/kit/viper"
)

func Init() {
	viper.Viper_init()
	mysqlDb.Db_Init()
	redis.Redis_init()
}
