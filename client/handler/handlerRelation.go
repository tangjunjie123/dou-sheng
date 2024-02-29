package handler

import (
	"client/consul"
	pro "client/proto/relation"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"strconv"
)

type Action struct {
	Token      string `json:"token"`
	ToUserId   int64  `json:"to_user_id"`
	ActionType int64  `json:"action_type"`
}
type Relation struct {
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

const (
	serviceIdRelation = "relation"
)

func BindJsonAction(ctx *gin.Context) (*Action, *grpc.ClientConn) {
	actionc := Action{}
	discovery, err := consul.Discovery(consul.ConsulIP, "", consul.ConsulPort, serviceIdRelation)
	if err != nil {
		fmt.Println(err)
	}
	if len(discovery) == 0 {
		return nil, nil
	}
	ctx.BindJSON(&actionc)
	dial, err := grpc.Dial(
		discovery[0].GetHost()+":"+strconv.Itoa(discovery[0].GetPort()),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		fmt.Println(err)
	}
	return &actionc, dial
}

func BindJsonRelation(ctx *gin.Context) (*Relation, *grpc.ClientConn) {
	discovery, err := consul.Discovery("127.0.0.1", "", 8500, serviceIdRelation)
	if err != nil {
		ctx.Error(err)
		return nil, nil
	}
	if len(discovery) == 0 {
		return nil, nil
	}
	relationc := Relation{}
	ctx.BindJSON(&relationc)
	dial, err := grpc.Dial(
		discovery[0].GetHost()+":"+strconv.Itoa(discovery[0].GetPort()),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		fmt.Println(err)
	}
	return &relationc, dial

}

func Actions(ctx *gin.Context) {
	action, conn := BindJsonAction(ctx)
	defer conn.Close()
	client := pro.NewUserClient(conn)
	res, err := client.Concern(
		ctx, &pro.DouyinRelationActionRequest{
			Token:      action.Token,
			ToUserId:   action.ToUserId,
			ActionType: int32(action.ActionType),
		},
	)
	if err != nil {
		ctx.JSON(200, err)
	}
	ctx.JSON(200, res)
}

func FansList(ctx *gin.Context) {
	relation, conn := BindJsonRelation(ctx)
	defer conn.Close()
	client := pro.NewUserClient(conn)
	res, err := client.FansList(
		context.Background(), &pro.DouyinRelationFollowerListRequest{
			UserId: relation.UserId,
			Token:  relation.Token,
		},
	)
	if err != nil {
		ctx.JSON(200, err)
	}
	ctx.JSON(200, res)
}

func ConcernList(ctx *gin.Context) {
	relation, conn := BindJsonRelation(ctx)
	defer conn.Close()
	client := pro.NewUserClient(conn)
	res, err := client.ConcernList(
		context.Background(), &pro.DouyinRelationFollowListRequest{
			UserId: relation.UserId,
			Token:  relation.Token,
		},
	)
	if err != nil {
		ctx.JSON(200, err)
	}
	ctx.JSON(200, res)
}
