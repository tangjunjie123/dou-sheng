package dao

import (
	"context"
	"fmt"
	mysqlDb "user/kit/mysql"
	"user/kit/redis"
)

func Insert(ctx context.Context, user User) (int64, error) {
	db := mysqlDb.Db
	tx := db.WithContext(ctx).Create(&user)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return tx.RowsAffected, tx.Error
	}
	return tx.RowsAffected, nil
}
func IsEmpty(ctx context.Context, user User) bool {
	db := mysqlDb.Db
	dbw := db.WithContext(ctx)
	res := new(User)
	dbw.Where("username = ?", user.Username).Find(&res)
	return res.IDIsEmpty()
}
func Find(ctx context.Context, user *User) (*User, error) {
	db := mysqlDb.Db
	dbw := db.WithContext(ctx)
	res := new(User)
	tx := dbw.Where("username = ? ", user.Username, user.Password).Find(&res)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}
	return res, nil
}

func FindByUsername(ctx context.Context, user *User) (*User, error) {
	db := mysqlDb.Db
	dbw := db.WithContext(ctx)
	res := new(User)
	tx := dbw.Where("username = ? ", user.Username).Find(&res)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}
	return res, nil
}
func FindByID(ctx context.Context, ID int64) (*User, error) {
	db := mysqlDb.Db
	dbw := db.WithContext(ctx)
	res := new(User)
	tx := dbw.Where("ID = ? ", ID).Find(&res)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}
	return res, nil
}
func Update(ctx context.Context, user User) error {
	db := mysqlDb.Db
	dbw := db.WithContext(ctx)
	tx := dbw.Where("account = ?", user.Username)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func AddConcern(ctx context.Context, user User, targetUser *User) error {
	client := redis.Redis
	client.SAdd(ctx, targetUser.Username+"fans", user.Username)
	client.SAdd(ctx, user.Username+"concern", targetUser.Username)
	return nil
}
