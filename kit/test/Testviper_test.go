package test

import (
	"fmt"
	"gorm.io/gorm"
	"kit/mysql"
	"kit/viper"
	"testing"
)

func TestViper(t *testing.T) {
	init := viper.Viper_init()
	get := init.Get("mysql.host")
	fmt.Println(get)
}
func TestGorm(t *testing.T) {
	db := mysqlDb.DbInit()
	err := db.AutoMigrate(&user{})
	if err != nil {
		fmt.Println(err)
	}

}

type user struct {
	gorm.Model
	Username string
	Password string
}
