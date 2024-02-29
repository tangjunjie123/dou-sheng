package test

import (
	"fmt"
	"user/kit/viper"

	"testing"
)

func TestViper(t *testing.T) {
	init := viper.Viper_init()
	get := init.Get("mysql.host")
	fmt.Println(get)
}
