package routes

import (
	"github.com/Wishrem/wuso/server/handler"
	"github.com/Wishrem/wuso/server/middleware"
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

	r.Group("/chat").Use(middleware.JWT).GET("", handler.ChatWs)

	friend := r.Group("/friend").Use(middleware.JWT)
	{
		friend.POST("/apply", handler.ApplyFriendship)
		friend.POST("/reply", handler.ReplyFriendshipApplication)
		friend.GET("/application", handler.GetFriendshipApplications)
		friend.GET("", handler.GetFriends)
	}

	return r
}
