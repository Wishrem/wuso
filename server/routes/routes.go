package routes

import (
	"github.com/Wishrem/wuso/server/handler"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", handler.Ping)

	user := r.Group("/user")
	{
		user.POST("/register", handler.UserRegisterReq)
		user.POST("/login", handler.UserLoginReq)
	}

	chat := r.Group("/chat")
	{
		chat.GET("", handler.ChatWs)
	}

	return r
}
