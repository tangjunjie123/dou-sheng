package service

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"strconv"
	"user/common"
	"user/dao"
	pro "user/proto"
)

type MyGrpcServer struct {
	pro.UnimplementedUserServer
}

func (myServer *MyGrpcServer) Register(
	ctx context.Context,
	request *pro.DouyinUserRegisterRequest,
) (*pro.DouyinUserRegisterResponse, error) {
	user := dao.User{
		Username: request.Username,
		Password: request.Password,
	}
	isEmpty := dao.IsEmpty(ctx, user)
	if isEmpty {
		return &pro.DouyinUserRegisterResponse{
			StatusCode: 200,
			StatusMsg:  common.RegisterError,
			UserId:     0,
			Token:      "",
		}, nil
	}
	// 添加md5加密
	user.Password = common.Md5(user.Password)
	row, err := dao.Insert(ctx, user)
	if row == 0 && err != nil {
		return &pro.DouyinUserRegisterResponse{
			StatusCode: 200,
			StatusMsg:  common.RegisterError,
			UserId:     0,
			Token:      "",
		}, err
	}
	dao.FindByUsername(ctx, &user)
	return &pro.DouyinUserRegisterResponse{
		StatusCode: 200,
		StatusMsg:  "",
		UserId:     int64(user.ID),
		Token:      common.GenerateTokenUsingHs256(common.User(user)),
	}, nil
}

func (myServer *MyGrpcServer) Login(
	ctx context.Context,
	request *pro.DouyinUserLoginRequest,
) (*pro.DouyinUserLoginResponse, error) {
	user := dao.User{
		Username: request.Username,
		Password: request.Password,
	}
	res, err := dao.FindByUsername(ctx, &user)
	if res.IDIsEmpty() == false || common.Md5(user.Password) != res.Password {
		StatusMsg := new(string)
		var errString = common.LoginError
		StatusMsg = &errString
		return &pro.DouyinUserLoginResponse{
			StatusCode: 200,
			StatusMsg:  StatusMsg, // ??
			UserId:     0,
			Token:      "",
		}, err
	}
	hs256 := common.GenerateTokenUsingHs256(common.User(user))
	StatusMsg := new(string)
	var errString = common.LoginOk
	StatusMsg = &errString
	return &pro.DouyinUserLoginResponse{
		StatusCode: 200,
		StatusMsg:  StatusMsg,
		UserId:     0,
		Token:      hs256,
	}, nil
}

func SecviceInit(Ip string, port int) {

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
