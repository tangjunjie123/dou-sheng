package service

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"strconv"
	"video/common"
	"video/dao"
	"video/kit/dbinit"
	pro "video/proto"
)

type MyGrpcServer struct {
	pro.UnimplementedVideoServer
}

func (*MyGrpcServer) Action(
	ctx context.Context,
	request *pro.DouyinPublishActionRequest,
) (*pro.DouyinPublishActionResponse, error) {
	hs256, err := common.ParseTokenHs256(request.Token)
	if err != nil {
		return nil, err
	}
	username, err := dao.FindByUsername(ctx, &dao.User{Username: hs256.User.Username})
	// todo 事务
	uid := username.ID
	err = dao.VideoInsertinfo(ctx, request.Title, int64(uid))
	if err != nil {
		return nil, err
	}
	info, err := dao.VideoFindinfo(ctx, request.Title, int64(uid))
	if err != nil {
		return nil, err
	}
	err = dao.Publish(ctx, request.Data, int64(info.ID))
	if err != nil {
		return nil, err
	}
	return &pro.DouyinPublishActionResponse{
		StatusCode: 200,
		StatusMsg:  "ok",
	}, nil
}
func (*MyGrpcServer) Videolike(
	ctx context.Context,
	request *pro.DouyinFavoriteActionRequest,
) (*pro.DouyinFavoriteActionResponse, error) {
	hs256, err := common.ParseTokenHs256(request.Token)
	if err != nil {
		return nil, err
	}
	username, err := dao.FindByUsername(ctx, &dao.User{Username: hs256.User.Username})
	if err != nil {
		return nil, err
	}
	isLike := dao.VideoIsLike(ctx, request.VideoId, int64(username.ID))
	if isLike {
		err := dao.VideoDisLike(ctx, request.VideoId, int64(username.ID))
		if err != nil {
			return nil, err
		}
		return &pro.DouyinFavoriteActionResponse{
			StatusCode: 200,
			StatusMsg:  "ok",
		}, nil

	}

	err = dao.VideoLike(ctx, request.VideoId, int64(username.ID))
	if err != nil {
		return nil, err
	}
	return &pro.DouyinFavoriteActionResponse{
		StatusCode: 200,
		StatusMsg:  "ok",
	}, nil
}

func (*MyGrpcServer) Comment(
	ctx context.Context,
	request *pro.DouyinCommentActionRequest,
) (*pro.DouyinCommentActionResponse, error) {
	hs256, err := common.ParseTokenHs256(request.Token)
	if err != nil {
		return nil, err
	}
	username, err := dao.FindByUsername(ctx, &dao.User{Username: hs256.User.Username})
	dao.VideoCommentinfo(ctx, int64(username.ID), request)
	info, err := dao.VideoGetCommentinfo(ctx, request, int64(username.ID))
	if err != nil {
		return nil, err
	}
	err = dao.VideoComments(ctx, request.VideoId, int64(username.ID), info)
	if err != nil {
		return nil, err
	}
	return &pro.DouyinCommentActionResponse{
		StatusCode: 200,
		StatusMsg:  "ok",
	}, nil
}
func (*MyGrpcServer) Getcomment(
	ctx context.Context,
	request *pro.DouyinCommentListRequest,
) (*pro.DouyinCommentListResponse, error) {

	comments := dao.VideoGetComments(ctx, request.VideoId)
	List := make([]*pro.Comment, 0)
	for _, v := range comments {
		List = append(
			List, &pro.Comment{
				Id: int64(v.ID),
				User: &pro.User{
					Id:              v.UserId,
					Name:            "v.UserId",
					FollowCount:     0,
					FollowerCount:   0,
					IsFollow:        false,
					Avatar:          "",
					BackgroundImage: "",
					Signature:       "",
					TotalFavorited:  0,
					WorkCount:       0,
					FavoriteCount:   0,
				},
				Content:    v.Comment,
				CreateDate: v.CreatedAt.String(),
			},
		)
	}
	return &pro.DouyinCommentListResponse{
		StatusCode:  200,
		StatusMsg:   "ok",
		CommentList: List,
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
	pro.RegisterVideoServer(grpcServer, &MyGrpcServer{})
	reflection.Register(grpcServer)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
