package types

type ChatSendMsgReq struct {
	Token string `json:"token" form:"token" binding:"required"`
}
