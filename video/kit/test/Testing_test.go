package test

import (
	"fmt"
	"testing"
	"video/dao"
	mysqlDb "video/kit/mysql"
)

func TestGorm(t *testing.T) {
	db := mysqlDb.Db_Init()
	err := db.AutoMigrate(&dao.CommentInfo{})
	if err != nil {
		fmt.Println(err)
	}
}
