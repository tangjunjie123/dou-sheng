package service

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"relation/common"
	"relation/dao"
	"relation/kit/dbinit"
	pro "relation/proto"
	"strconv"
)

type MyGrpcServer struct {
	pro.UnimplementedUserServer
}
type user struct {
	ID       float64
	CreateAt string
	UpdateAt string
	DeleteAt string
	Username string
	Password string
}

func (myServer *MyGrpcServer) Concern(ctx context.Context, Request *pro.DouyinRelationActionRequest) (
	*pro.DouyinRelationActionResponse,
	error,
) {
	hs256, err2 := common.ParseTokenHs256(Request.Token)
	if err2 != nil {
		return nil, err2
	}
	targetUser, err := dao.FindByID(ctx, Request.ToUserId)
	if err != nil {
		return &pro.DouyinRelationActionResponse{
			StatusCode: 0,
			StatusMsg:  err.Error(),
		}, err
	}
	err = dao.AddConcern(ctx, hs256.User.Username, targetUser.Username)
	if err != nil {
		return &pro.DouyinRelationActionResponse{
			StatusCode: 0,
			StatusMsg:  err.Error(),
		}, err
	}

	return &pro.DouyinRelationActionResponse{
		StatusCode: 200,
		StatusMsg:  "ok",
	}, nil
}

func (myServer *MyGrpcServer) FansList(
	ctx context.Context, request *pro.DouyinRelationFollowerListRequest,
) (*pro.DouyinRelationFollowerListResponse, error) {
	user, err := dao.FindByID(ctx, request.UserId)
	if err != nil {
		return nil, err
	}

	fans, err := dao.GetFans(ctx, user.Username)
	if err != nil {
		return nil, err
	}
	res := make([]*pro.User, 0)
	for _, v := range fans {
		res = append(
			res, &pro.User{
				Id:              0,
				Name:            v,
				FollowCount:     0,
				FollowerCount:   0,
				IsFollow:        true,
				Avatar:          "",
				BackgroundImage: "",
				Signature:       "",
				TotalFavorited:  0,
				WorkCount:       0,
				FavoriteCount:   0,
			},
		)
	}
	return &pro.DouyinRelationFollowerListResponse{
		StatusCode: 200,
		StatusMsg:  "ok",
		UserList:   res,
	}, nil
}

func (myServer *MyGrpcServer) ConcernList(ctx context.Context, request *pro.DouyinRelationFollowListRequest) (
	*pro.DouyinRelationFollowListResponse, error,
) {

	user, err := dao.FindByID(ctx, request.UserId)
	if err != nil {
		return nil, err
	}

	fans, err := dao.GetConcern(ctx, user.Username)
	if err != nil {
		return nil, err
	}
	res := make([]*pro.User, 0)
	for _, v := range fans {
		res = append(
			res, &pro.User{
				Id:              0,
				Name:            v,
				FollowCount:     0,
				FollowerCount:   0,
				IsFollow:        true,
				Avatar:          "",
				BackgroundImage: "",
				Signature:       "",
				TotalFavorited:  0,
				WorkCount:       0,
				FavoriteCount:   0,
			},
		)
	}
	return &pro.DouyinRelationFollowListResponse{
		StatusCode: 200,
		StatusMsg:  "ok",
		UserList:   res,
	}, nil
}

func SecviceInit(Ip string, port int) {
	dbinit.Init()
	// 创建 Tcp 连接
	listener, err := net.Listen("tcp", Ip+":"+strconv.Itoa(port))
	//listener, err := net.Listen("tcp", "127.0.0.1:8091")
	if err != nil {
		log.Fatalf("监听失败:  %v", err)
	}
	grpcServer := grpc.NewServer()
	pro.RegisterUserServer(grpcServer, &MyGrpcServer{})
	reflection.Register(grpcServer)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
