package model

type Friendship struct {
	Id      int64
	UserId1 int64
	UserId2 int64
}

type FriendReq struct {
	Id         int64
	SenderId   int64
	ReceiverId int64
	Status     string
}
