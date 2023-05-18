package handler

import (
	"context"
	"time"

	"github.com/Wishrem/wuso/pkg/errno"
	"github.com/Wishrem/wuso/server/types"
	user "github.com/Wishrem/wuso/server/user/service"
	"github.com/gin-gonic/gin"
)

func UserRegisterReq(c *gin.Context) {
	req := new(types.UserRegisterReq)
	if err := c.ShouldBind(req); err != nil {
		SendFailureResp(c, err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := user.CreateUser(ctx, req); err != nil {
		SendFailureResp(c, err)
		return
	}

	SendSuccessResp(c, nil)
}

func UserLoginReq(c *gin.Context) {
	req := new(types.UserLoginReq)
	if err := c.ShouldBind(req); err != nil {
		SendFailureResp(c, errno.ParamError)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	resp, err := user.LoginUser(ctx, req)
	if err != nil {
		SendFailureResp(c, err)
		return
	}

	SendSuccessResp(c, resp)
}
