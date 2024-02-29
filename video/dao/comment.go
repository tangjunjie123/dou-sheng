package dao

import "gorm.io/gorm"

type CommentInfo struct {
	VideoId int64
	UserId  int64
	Comment string
	gorm.Model
}
