package types

type ApplyFriendshipReq struct {
	ReceiverId int64 `json:"receiver_id" form:"receiver_id" binding:"required"`
}

type ReplyFriendshipApplicationReq struct {
	Token    string `json:"token" form:"token" binding:"required"`
	SenderId int64  `json:"sender_id" form:"sender_id" binding:"required"`
	Accept   bool   `json:"accept" form:"accept" binding:"required"`
}

type GetFriendsReq struct {
	Page int `json:"page" form:"page" binding:"required"`
}

type GetFriendsResp struct {
	Users []User `json:"users"`
	Page  int    `json:"page"`
}

type GetFriendshipApplicationsReq struct {
	Page int `json:"page" form:"page" binding:"required"`
}

type Application struct {
	SenderId int64  `json:"sender_id"`
	Status   string `json:"status"`
}

type GetFriendshipApplicationsResp struct {
	Applications []Application `json:"applications"`
	Page         int           `json:"page"`
}
