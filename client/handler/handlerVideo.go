package handler

import (
	"client/consul"
	"client/proto/video"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"strconv"
)

const (
	userserviceIdVideo = "video"
)

type Video struct {
	Token string `json:"token"` // 用户鉴权token
	Data  string `json:"data"`  // 视频数据
	Title string `json:"title"` // 视频标题
}

type VideoComments struct {
	Token        string `json:"token"`        // 用户鉴权token
	Video_id     int64  `json:"video_id"`     // 视频id
	Action_type  int32  `json:"action_type"`  // 1-发布评论，2-删除评论
	Comment_text string `json:"comment_text"` // 用户填写的评论内容，在action_type=1的时候使用
	Comment_id   int64  `json:"comment_id"`   // 要删除的评论id，在action_type=2的时候使用
}

type VideoLikes struct {
	Token       string `json:"token"`       // 用户鉴权token
	Video_id    int64  `json:"video_id"`    // 视频id
	Action_type int32  `json:"action_type"` // 1-点赞，2-取消点赞
}

type GetComment struct {
	Token    string `json:"token"`    // 用户鉴权token
	Video_id int64  `json:"video_id"` // 视频id
}

func bindJsonGetComment(ctx *gin.Context) (*GetComment, *grpc.ClientConn) {
	discovery, err := consul.Discovery(consul.ConsulIP, "", consul.ConsulPort, userserviceIdVideo)
	if err != nil {
		fmt.Println(err)
	}
	GetCommentc := &GetComment{}
	if len(discovery) == 0 {
		return nil, nil
	}
	ctx.BindJSON(&GetCommentc)
	dial, err := grpc.Dial(
		discovery[0].GetHost()+":"+strconv.Itoa(discovery[0].GetPort()),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		fmt.Println(err)
	}
	return GetCommentc, dial
}

func bindJsonVideo(ctx *gin.Context) (*Video, *grpc.ClientConn) {
	discovery, err := consul.Discovery(consul.ConsulIP, "", consul.ConsulPort, userserviceIdVideo)
	if err != nil {
		fmt.Println(err)
	}
	videoc := &Video{}
	if len(discovery) == 0 {
		return nil, nil
	}
	ctx.BindJSON(&videoc)
	dial, err := grpc.Dial(
		discovery[0].GetHost()+":"+strconv.Itoa(discovery[0].GetPort()),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		fmt.Println(err)
	}
	return videoc, dial
}

func bindJsonComments(ctx *gin.Context) (*grpc.ClientConn, *VideoComments) {
	discovery, err := consul.Discovery(consul.ConsulIP, "", consul.ConsulPort, userserviceIdVideo)
	if err != nil {
		fmt.Println(err)
	}
	Comment := &VideoComments{}
	if len(discovery) == 0 {
		return nil, nil
	}
	ctx.BindJSON(&Comment)
	dial, err := grpc.Dial(
		discovery[0].GetHost()+":"+strconv.Itoa(discovery[0].GetPort()),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		fmt.Println(err)
	}
	return dial, Comment
}

func bindJsonLike(ctx *gin.Context) (*grpc.ClientConn, *VideoLikes) {
	discovery, err := consul.Discovery(consul.ConsulIP, "", consul.ConsulPort, userserviceIdVideo)
	if err != nil {
		fmt.Println(err)
	}
	Like := &VideoLikes{}
	if len(discovery) == 0 {
		return nil, nil
	}
	ctx.BindJSON(&Like)
	dial, err := grpc.Dial(
		discovery[0].GetHost()+":"+strconv.Itoa(discovery[0].GetPort()),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		fmt.Println(err)
	}
	return dial, Like
}

func VideoPublish(ctx *gin.Context) {
	videoc, dial := bindJsonVideo(ctx)
	defer dial.Close()
	client := video.NewVideoClient(dial)
	action, err := client.Action(
		ctx, &video.DouyinPublishActionRequest{
			Token: videoc.Token,
			Data:  []byte(videoc.Data),
			Title: videoc.Title,
		},
	)
	if err != nil {
		ctx.JSON(200, err)
	}
	ctx.JSON(200, action)
}

func VideoGet(ctx *gin.Context) {

}

func VideoComment(ctx *gin.Context) {
	dial, videoComments := bindJsonComments(ctx)
	client := video.NewVideoClient(dial)
	comment, err := client.Comment(
		ctx, &video.DouyinCommentActionRequest{
			Token:       videoComments.Token,
			VideoId:     videoComments.Video_id,
			ActionType:  videoComments.Action_type,
			CommentText: videoComments.Comment_text,
			CommentId:   videoComments.Comment_id,
		},
	)
	if err != nil {
		ctx.JSON(200, err)
	}
	ctx.JSON(200, comment)

}

func VideoLike(ctx *gin.Context) {

	dial, likes := bindJsonLike(ctx)
	client := video.NewVideoClient(dial)
	action, err := client.Videolike(
		ctx, &video.DouyinFavoriteActionRequest{
			Token:      likes.Token,
			VideoId:    likes.Video_id,
			ActionType: likes.Action_type,
		},
	)
	if err != nil {
		ctx.JSON(200, err)
	}
	ctx.JSON(200, action)

}

func VideoGetLike(ctx *gin.Context) {

}

func VideoGetComment(ctx *gin.Context) {
	res, dial := bindJsonGetComment(ctx)
	client := video.NewVideoClient(dial)
	action, err := client.Getcomment(
		ctx, &video.DouyinCommentListRequest{
			Token:   res.Token,
			VideoId: res.Video_id,
		},
	)
	if err != nil {
		ctx.JSON(200, err)
	}
	ctx.JSON(200, action)
}
