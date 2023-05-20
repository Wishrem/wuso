package handler

import (
	"log"

	"github.com/Wishrem/wuso/pkg/errno"
	chat "github.com/Wishrem/wuso/server/chat/service"
	"github.com/Wishrem/wuso/server/consts"

	"github.com/Wishrem/wuso/server/types"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func ChatWs(c *gin.Context) {
	req := new(types.ChatSendMsgReq)
	if err := c.ShouldBind(req); err != nil {
		log.Println(err)
		SendFailureResp(c, errno.ParamError)
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		SendFailureResp(c, errno.New(errno.UpgradeFailedCode, err.Error()))
		return
	}

	chat.RegisterClient(c.GetInt64(consts.KeyUserId), conn)
}
