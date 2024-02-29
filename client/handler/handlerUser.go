package handler

import (
	"client/consul"
	user2 "client/proto/user"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"strconv"
)

type User struct {
	UserId   int64  `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

const (
	userserviceIdUser = "user"
)

func BindJson(ctx *gin.Context) (*User, *grpc.ClientConn) {
	user := User{}
	discovery, err := consul.Discovery(consul.ConsulIP, "", consul.ConsulPort, userserviceIdUser)
	if err != nil {
		fmt.Println(err)
	}
	if len(discovery) == 0 {
		return nil, nil
	}
	ctx.BindJSON(&user)
	dial, err := grpc.Dial(
		discovery[0].GetHost()+":"+strconv.Itoa(discovery[0].GetPort()),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		fmt.Println(err)
	}
	return &user, dial
}

func Register(ctx *gin.Context) {
	userc, dial := BindJson(ctx)
	defer dial.Close()
	client := user2.NewUserClient(dial)
	res, err := client.Register(
		context.Background(), &user2.DouyinUserRegisterRequest{
			Username: userc.UserName,
			Password: userc.Password,
		},
	)
	if err != nil {
		ctx.JSON(200, err)
	}
	ctx.JSON(200, res)
}

func Login(ctx *gin.Context) {
	userc, dial := BindJson(ctx)
	defer dial.Close()
	client := user2.NewUserClient(dial)
	res, err := client.Login(
		context.Background(), &user2.DouyinUserLoginRequest{
			Username: userc.UserName,
			Password: userc.Password,
		},
	)
	if err != nil {
		ctx.JSON(200, err)
	}
	ctx.JSON(200, res)
}

func UserInfo(ctx *gin.Context) {
	userc, dial := BindJson(ctx)
	defer dial.Close()
	client := user2.NewUserClient(dial)
	res, err := client.UserInfo(
		context.Background(), &user2.DouyinUserRequest{
			UserId: userc.UserId,
		},
	)
	if err != nil {
		ctx.JSON(200, err)
	}
	ctx.JSON(200, res)
}
