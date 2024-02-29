package router

import (
	"client/handler"
	"github.com/gin-gonic/gin"
)

func UserRouter(engine *gin.Engine) {
	group := engine.Group("/douyin")
	{
		group.GET("/user/register", handler.Register).
			POST("/user/login", handler.Login).
			GET("/user")
	}
	group = engine.Group("/douyin/relation")
	{
		group.GET("/follow/list", handler.ConcernList).
			GET("/follower/list", handler.FansList).
			GET("/action", handler.Actions)
	}
	group = engine.Group("/douyin")
	{
		group.GET("/publish/action", handler.VideoPublish).
			POST("/publish/list", handler.VideoGet)
	}
	group = engine.Group("/douyin")
	{
		group.GET("/comment/action", handler.VideoComment).
			POST("/comment/list", handler.VideoGetComment)
	}
	group = engine.Group("/douyin")
	{
		group.GET("/favorite/action", handler.VideoLike).
			POST("/favorite/list", handler.VideoGetLike)
	}

}
