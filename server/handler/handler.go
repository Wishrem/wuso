package handler

import (
	"github.com/Wishrem/wuso/pkg/errno"
	"github.com/Wishrem/wuso/server/types"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func SendResp(c *gin.Context, base types.BaseResp, data interface{}) {
	c.JSON(200, gin.H{
		"code": base.Code,
		"msg":  base.Msg,
		"data": data,
	})
}

func SendFailureResp(c *gin.Context, err error) {
	errno := errno.Convert(err)
	SendResp(c, types.BaseResp{
		Code: int64(errno.Code),
		Msg:  errno.Msg,
	}, nil)
}

func SendSuccessResp(c *gin.Context, data interface{}) {
	SendResp(c, types.BaseResp{
		Code: int64(errno.SuccessCode),
		Msg:  errno.SuccessMsg,
	}, data)
}
