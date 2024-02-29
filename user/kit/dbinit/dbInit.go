package dbinit

import (
	mysqlDb "user/kit/mysql"
	"user/kit/redis"
	"user/kit/viper"
)

func Init() {
	viper.Viper_init()
	mysqlDb.Db_Init()
	redis.Redis_init()

}
