package dao

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
}

func (this *User) IDIsEmpty() bool {
	if this.ID == 0 {
		return false
	}
	return true
}
func (this *User) TableName() string {
	return "user"
}
