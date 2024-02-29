package mongodb

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Mongo *mongo.Database

func Mongodb_init() *mongo.Client {
	username := viper.GetString("mongodb.username")
	password := viper.GetString("mongodb.password")
	host := viper.GetString("mongodb.host")
	port := viper.GetInt("mongodb.port")
	dbname := viper.GetString("mongodb.dbname")
	client, err := mongo.Connect(
		context.Background(),
		options.Client().
			// 连接地址
			ApplyURI(fmt.Sprintf("mongodb://%s:%d", host, port)).
			// 设置验证参数
			SetAuth(
				options.Credential{
					// 用户名
					Username: fmt.Sprintf("%s", username),
					// 密码
					Password: fmt.Sprintf("%s", password),
				},
			),
	)

	if err != nil {
		panic(err)
	}
	Mongo = client.Database(fmt.Sprintf("%s", dbname))
	return client
}
