package service

import (
	"context"
	"errors"
	"regexp"

	"github.com/Wishrem/wuso/config"
	"github.com/Wishrem/wuso/pkg/errno"
	"github.com/Wishrem/wuso/pkg/utils/jwt"
	"github.com/Wishrem/wuso/server/types"
	"github.com/Wishrem/wuso/server/user/dal/db"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var regex = regexp.MustCompile(`^[a-zA-Z0-9_]+@[a-zA-Z0-9]+\.[a-zA-Z]{2,}$`)

func encryptedPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func checkPassword(password, encryptedPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(password)) == nil
}

func CreateUser(ctx context.Context, req *types.UserRegisterReq) error {
	if !regex.MatchString(req.Email) {
		return errno.InvalidEmailFormat
	}

	pwd, err := encryptedPassword(req.Password)
	if err != nil {
		return err
	}
	req.Password = pwd

	if ctx.Err() != nil {
		return errno.ExecuteTimeout
	}

	if err := db.CreateUser(ctx, req); err != nil {
		mysqlErr := &mysql.MySQLError{}
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return errno.DuplicatedEmail
		}
		return err
	}

	return nil
}

func LoginUser(ctx context.Context, req *types.UserLoginReq) (*types.UserResp, error) {
	user, err := db.GetUserByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.UserNotFound
		}
		return nil, err
	}

	if !checkPassword(req.Password, user.Password) {
		return nil, errno.WrongPassword
	}

	token, err := jwt.Generate(user.ID, config.JWT.Secret)
	if err != nil {
		return nil, err
	}

	return &types.UserResp{
		User: types.User{
			ID:    user.ID,
			Email: user.Email,
			Name:  user.Name,
		},
		Token: token,
	}, nil
}
