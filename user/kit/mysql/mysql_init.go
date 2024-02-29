package mysqlDb

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"user/kit/viper"
)

var Db *gorm.DB

func Db_Init() *gorm.DB {
	conf := viper.Viper_init()
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.Get("mysql.username"),
		conf.Get("mysql.password"),
		conf.Get("mysql.host"),
		conf.Get("mysql.port"),
		conf.Get("mysql.dbname"),
	)
	res, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database, error=" + err.Error())
	}
	Db = res
	return res
}
