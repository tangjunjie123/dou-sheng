package dbinit

import (
	"video/kit/mongodb"
	mysqlDb "video/kit/mysql"
	"video/kit/redis"
	"video/kit/viper"
)

func Init() {
	viper.Viper_init()
	mysqlDb.Db_Init()
	redis.Redis_init()
	mongodb.Mongodb_init()
}
