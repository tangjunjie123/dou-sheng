package dao

import "gorm.io/gorm"

type VideoInfo struct {
	gorm.Model
	VideoName string
	UserId    int64
}
