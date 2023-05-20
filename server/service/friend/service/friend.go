package service

import (
	"context"
	"errors"

	"github.com/Wishrem/wuso/pkg/errno"
	"github.com/Wishrem/wuso/server/service/friend/dal/db"
	"github.com/Wishrem/wuso/server/service/friend/model"
	user "github.com/Wishrem/wuso/server/service/user/service"
	"github.com/Wishrem/wuso/server/types"
	"gorm.io/gorm"
)

func GetFriends(ctx context.Context, userId int64, page int) (*types.GetFriendsResp, error) {
	ids, err := db.GetFriendsIds(ctx, userId, page, 10)
	if err != nil {
		return nil, err
	}

	users, err := user.GetUsers(ctx, ids)
	if err != nil {
		return nil, err
	}

	return &types.GetFriendsResp{
		Users: users,
		Page:  page,
	}, nil
}

func ApplyFriendship(ctx context.Context, senderId, receiverId int64) error {
	has, err := db.HasFriendReq(senderId, receiverId)
	if err != nil {
		return err
	}

	if has {
		return errno.AlreadyAppliedFriendship
	}

	if err := db.CreateFriendReq(ctx, senderId, receiverId); err != nil {
		return err
	}

	return nil
}

func ReplyFriendshipApplication(ctx context.Context, senderId, receiverId int64, accept bool) error {
	if accept {
		return db.CreateFriendship(ctx, senderId, receiverId)
	}

	return db.RefuseFriendReq(ctx, senderId)
}

func GetFriendshipApplications(ctx context.Context, receiverId int64, page int) (*types.GetFriendshipApplicationsResp, error) {
	reqs, err := db.GetFriendshipApplications(ctx, receiverId, page, 10)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.ApplicationsNotFound
		}
		return nil, err
	}
	return convertToGetFriendshipApplicationsResp(reqs, page), nil
}

func convertToGetFriendshipApplicationsResp(reqs []model.FriendReq, page int) *types.GetFriendshipApplicationsResp {
	apps := make([]types.Application, len(reqs))
	for i, v := range reqs {
		apps[i] = types.Application{
			SenderId: v.SenderId,
			Status:   v.Status,
		}
	}
	return &types.GetFriendshipApplicationsResp{
		Applications: apps,
		Page:         page,
	}
}
