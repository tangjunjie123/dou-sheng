package test

import (
	"client/token"
	"fmt"
	"gorm.io/gorm"
	"time"
	"user/dao"
	"user/kit/viper"

	"testing"
	mysqlDb "user/kit/mysql"
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
func TestToken(t *testing.T) {
	tokenString, err := token.GenerateTokenUsingHs256(
		dao.User{
			Model:    gorm.Model{},
			Name:     "123",
			Account:  "123",
			Password: "12312123",
		},
	)
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(2000 * time.Millisecond)
	tokenHs256, err := token.ParseTokenHs256(tokenString)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(tokenHs256)
}

type user struct {
	gorm.Model
	Name     string
	Account  string
	Password string
}
