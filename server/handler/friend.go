package handler

import (
	"context"
	"time"

	"github.com/Wishrem/wuso/pkg/errno"
	"github.com/Wishrem/wuso/server/consts"
	friend "github.com/Wishrem/wuso/server/friend/service"
	"github.com/Wishrem/wuso/server/types"
	"github.com/gin-gonic/gin"
)

func ApplyFriendship(c *gin.Context) {
	req := new(types.ApplyFriendshipReq)
	if err := c.ShouldBind(req); err != nil {
		SendFailureResp(c, errno.ParamError)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := friend.ApplyFriendship(ctx, c.GetInt64(consts.KeyUserId), req.ReceiverId); err != nil {
		SendFailureResp(c, err)
		return
	}

	SendSuccessResp(c, nil)
}

func ReplyFriendshipApplication(c *gin.Context) {
	req := new(types.ReplyFriendshipApplicationReq)
	if err := c.ShouldBind(req); err != nil {
		SendFailureResp(c, errno.ParamError)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := friend.ReplyFriendshipApplication(ctx, req.SenderId, c.GetInt64(consts.KeyUserId), req.Accept); err != nil {
		SendFailureResp(c, err)
		return
	}

	SendSuccessResp(c, nil)
}

func GetFriendshipApplications(c *gin.Context) {
	req := new(types.GetFriendshipApplicationsReq)
	if err := c.ShouldBind(req); err != nil {
		SendFailureResp(c, errno.ParamError)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	resp, err := friend.GetFriendshipApplications(ctx, c.GetInt64(consts.KeyUserId), req.Page)
	if err != nil {
		SendFailureResp(c, err)
		return
	}

	SendSuccessResp(c, resp)
}

func GetFriends(c *gin.Context) {
	req := new(types.GetFriendsReq)
	if err := c.ShouldBind(req); err != nil {
		SendFailureResp(c, errno.ParamError)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	resp, err := friend.GetFriends(ctx, c.GetInt64(consts.KeyUserId), req.Page)
	if err != nil {
		SendFailureResp(c, err)
		return
	}

	SendSuccessResp(c, resp)
}
