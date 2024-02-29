package dao

import (
	"context"
	"fmt"
	mysqlDb "relation/kit/mysql"
	"relation/kit/redis"
)

func AddConcern(ctx context.Context, user string, targetUser string) error {
	client := redis.Redis
	client.SAdd(ctx, targetUser+"fans", user)
	client.SAdd(ctx, user+"concern", targetUser)
	return nil
}

func RemoveConcern(ctx context.Context, user string, targetUser string) error {
	client := redis.Redis
	client.SRem(ctx, targetUser+"fans", user)
	client.SRem(ctx, user+"concern", targetUser)
	return nil
}

func GetFans(ctx context.Context, user string) ([]string, error) {
	client := redis.Redis
	res, err := client.SMembers(ctx, user+"fans").Result()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetConcern(ctx context.Context, user string) ([]string, error) {
	client := redis.Redis
	res, err := client.SMembers(ctx, user+"concern").Result()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func FindByID(ctx context.Context, ID int64) (*User, error) {
	db := mysqlDb.Db
	dbw := db.WithContext(ctx)
	res := new(User)
	tx := dbw.Where("id = ? ", ID).Find(&res)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}
	return res, nil
}
