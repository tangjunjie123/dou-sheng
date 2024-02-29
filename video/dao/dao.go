package dao

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"strconv"
	"video/kit/mongodb"
	mysqlDb "video/kit/mysql"
	"video/kit/redis"
	pro "video/proto"
)

type Comments struct {
	_ID     primitive.ObjectID `bson:"_id"`
	comment CommentInfo
}

func Publish(ctx context.Context, data []byte, VideoID int64) error {
	d := bson.D{{"video", data}}
	_, err := mongodb.Mongo.Collection("video"+strconv.Itoa(int(VideoID))).InsertOne(ctx, d)
	if err != nil {
		return err
	}
	return nil
}

func GetLastVideo(ctx context.Context, Ids []int) [][]byte {
	mongodb.Mongo.Collection("video"+strconv.Itoa(int(Ids[0]))).Find(ctx, bson.D{})
	return nil
}

func VideoInsertinfo(ctx context.Context, Name string, uid int64) error {
	tx := mysqlDb.Db.Create(&VideoInfo{VideoName: Name, UserId: uid})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func VideoFindinfo(ctx context.Context, Name string, uid int64) (*VideoInfo, error) {
	res := &VideoInfo{}
	tx := mysqlDb.Db.Where("video_name = ? && user_id = ?", Name, uid).Find(&res)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return res, nil
}

func VideoGetLastNum(ctx *context.Context, limit int) ([]VideoInfo, error) {
	res := make([]VideoInfo, 0)
	tx := mysqlDb.Db.Model(&VideoInfo{}).Order("ID desc").Limit(limit).Find(&res)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return res, nil
}

func VideoComments(ctx context.Context, videoId int64, uid int64, comment *CommentInfo) error {
	_, err := mongodb.Mongo.Collection("video"+strconv.Itoa(int(videoId))).InsertOne(ctx, comment)
	if err != nil {
		return err
	}
	return nil
}

func VideoCommentinfo(ctx context.Context, uid int64, comment *pro.DouyinCommentActionRequest) {
	mysqlDb.Db.Create(
		&CommentInfo{
			VideoId: comment.VideoId,
			UserId:  uid,
			Comment: comment.CommentText,
		},
	)
}

func VideoGetCommentinfo(ctx context.Context, comment *pro.DouyinCommentActionRequest, uid int64) (
	*CommentInfo,
	error,
) {
	res := &CommentInfo{}
	tx := mysqlDb.Db.Where(
		"video_id = ? && user_id = ? && comment = ?",
		comment.VideoId,
		uid,
		comment.CommentText,
	).Find(res)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return res, nil
}

func VideoGetComments(ctx context.Context, videoId int64) []CommentInfo {
	query := bson.M{"video": bson.M{"$exists": false}}
	//query := bson.D{{}}
	cur, _ := mongodb.Mongo.Collection("video"+strconv.Itoa(int(videoId))).Find(ctx, query)
	res := make([]CommentInfo, 0)
	for cur.Next(context.TODO()) {
		// 创建一个值，将单个文档解码为该值
		var elem CommentInfo
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		res = append(res, elem)
	}
	return res
}

func VideoLike(ctx context.Context, videoID int64, userID int64) error {
	client := redis.Redis
	client.SAdd(ctx, strconv.Itoa(int(videoID))+"like", userID)
	client.SAdd(ctx, strconv.Itoa(int(userID))+"liker", videoID)
	return nil
}
func VideoDisLike(ctx context.Context, videoID int64, userID int64) error {
	client := redis.Redis
	client.SRem(ctx, strconv.Itoa(int(videoID))+"like", userID)
	client.SRem(ctx, strconv.Itoa(int(userID))+"liker", videoID)
	return nil
}
func VideoIsLike(ctx context.Context, videoID int64, userID int64) bool {
	client := redis.Redis
	member := client.SIsMember(ctx, strconv.Itoa(int(videoID))+"like", userID)
	result, _ := member.Result()
	return result
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
