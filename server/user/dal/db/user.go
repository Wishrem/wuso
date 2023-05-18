package db

import (
	"context"

	"github.com/Wishrem/wuso/server/types"
	"github.com/Wishrem/wuso/server/user/model"
	"github.com/yitter/idgenerator-go/idgen"
)

func CreateUser(ctx context.Context, req *types.UserRegisterReq) error {
	return DB.Create(&model.User{
		ID:       idgen.NextId(),
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}).Error
}

func GetUserByID(ctx context.Context, id int64) (*model.User, error) {
	user := new(model.User)
	err := DB.Where("id = ?", id).First(user).Error
	return user, err
}

func GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	user := new(model.User)
	err := DB.Where("email = ?", email).First(user).Error
	return user, err
}
