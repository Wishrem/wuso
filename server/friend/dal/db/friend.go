package db

import (
	"context"
	"errors"

	"github.com/Wishrem/wuso/pkg/errno"
	"github.com/Wishrem/wuso/server/friend/model"
	"github.com/yitter/idgenerator-go/idgen"
	"gorm.io/gorm"
)

func GetFriendsIds(ctx context.Context, userId int64, page, items int) ([]int64, error) {
	if ctx.Err() != nil {
		return nil, errno.ExecuteTimeout
	}
	var friendships []model.Friendship
	if err := DB.Table("friendship").Where("user_id1 = ? OR user_id2 = ?", userId).Find(&friendships).Limit(items).Offset(page * items).Error; err != nil {
		return nil, err
	}

	ids := make([]int64, 1)
	for _, v := range friendships {
		if v.UserId1 != userId {
			ids = append(ids, v.UserId1)
		} else {
			ids = append(ids, v.UserId2)
		}
	}

	return ids, nil
}

func HasFriendReq(senderId, receiverId int64) (bool, error) {
	req := new(model.FriendReq)
	err := DB.Table("friend_req").Where("sender_id = ? AND receiver_id = ?", senderId, receiverId).First(req).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		} else {
			return true, err
		}
	} else {
		return true, nil
	}
}

func CreateFriendReq(ctx context.Context, senderId, receiverId int64) error {
	if ctx.Err() != nil {
		return errno.ExecuteTimeout
	}
	return DB.Table("friend_req").Create(&model.FriendReq{
		Id:         idgen.NextId(),
		SenderId:   senderId,
		ReceiverId: receiverId,
		Status:     "0",
	}).Error
}

func RefuseFriendReq(ctx context.Context, senderId int64) error {
	if ctx.Err() != nil {
		return errno.ExecuteTimeout
	}
	return DB.Table("friend_req").Where("sender_id = ?", senderId).Update("status", "2").Error
}

func CreateFriendship(ctx context.Context, senderId, userId int64) error {
	if ctx.Err() != nil {
		return errno.ExecuteTimeout
	}

	tx := DB.Begin()

	if err := tx.Table("friendship").Create(&model.Friendship{
		Id:      idgen.NextId(),
		UserId1: senderId,
		UserId2: userId,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Table("friend_req").Where("sender_id = ?", senderId).Update("status", "1").Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func DeleteFriendship(ctx context.Context, userId1, userId2 int64) error {
	if ctx.Err() != nil {
		return errno.ExecuteTimeout
	}
	return DB.Table("friendship").Delete("user_id1 = ? AND user_id2 = ?", userId1, userId2).Error
}

func GetFriendshipApplications(ctx context.Context, receiverId int64, page, items int) ([]model.FriendReq, error) {
	if ctx.Err() != nil {
		return nil, errno.ExecuteTimeout
	}
	var reqs []model.FriendReq
	if err := DB.Table("friendship").Where("receiver_id = ?", receiverId).Find(&reqs).Limit(items).Offset(page * items).Error; err != nil {
		return nil, err
	}
	return reqs, nil
}
