package db

import (
	"context"
	"log"

	"github.com/Wishrem/wuso/pkg/errno"
	"github.com/Wishrem/wuso/server/types"
	"github.com/Wishrem/wuso/server/user/model"
	"github.com/yitter/idgenerator-go/idgen"
)

func CreateUser(ctx context.Context, req *types.UserRegisterReq) error {
	if ctx.Err() != nil {
		return errno.ExecuteTimeout
	}
	return DB.Table("user").Create(&model.User{
		ID:       idgen.NextId(),
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}).Error
}

func GetUserByID(ctx context.Context, id int64) (*model.User, error) {
	if ctx.Err() != nil {
		return nil, errno.ExecuteTimeout
	}
	user := new(model.User)
	err := DB.Table("user").Where("id = ?", id).First(user).Error
	return user, err
}

func GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	if ctx.Err() != nil {
		return nil, errno.ExecuteTimeout
	}
	user := new(model.User)
	log.Println(email)
	err := DB.Table("user").Where("email = ?", email).First(user).Error
	return user, err
}

func GetUsersByIds(ctx context.Context, userIds []int64) ([]model.User, error) {
	if ctx.Err() != nil {
		return nil, errno.ExecuteTimeout
	}
	var users []model.User
	err := DB.Table("user").Find(&users, userIds).Error
	return users, err
}
