package middleware

import (
	"github.com/Wishrem/wuso/config"
	"github.com/Wishrem/wuso/pkg/errno"
	"github.com/Wishrem/wuso/pkg/utils/jwt"
	"github.com/Wishrem/wuso/server/consts"
	"github.com/Wishrem/wuso/server/handler"
	"github.com/gin-gonic/gin"
)

func JWT(c *gin.Context) {

	token := c.GetHeader(consts.AuthHeader)
	if token == "" {
		handler.SendFailureResp(c, errno.AuthorizationFailed)
		c.Abort()
		return
	}
	claims, err := jwt.Parse(token, config.JWT.Secret)
	if err != nil {
		handler.SendFailureResp(c, errno.AuthorizationFailed)
		c.Abort()
		return
	}

	token, err = jwt.Generate(claims.UserId, config.JWT.Secret)

	if err != nil {
		handler.SendFailureResp(c, errno.New(errno.AuthorizationFailedCode, err.Error()))
		c.Abort()
		return
	}
	c.Header(consts.AuthHeader, token)
	c.Set(consts.KeyUserId, claims.UserId)

	c.Next()
}
