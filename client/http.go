package main

import (
	router "client/router"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	router.UserRouter(engine)
	engine.Run(":8081")
}
