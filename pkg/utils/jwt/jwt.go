package jwt

import (
	"time"

	"github.com/Wishrem/wuso/config"
	"github.com/Wishrem/wuso/pkg/errno"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	UserId int64 `json:"user_id"`
	jwt.StandardClaims
}

func (c Claims) Valid() error {
	if !c.VerifyExpiresAt(time.Now().Unix(), true) {
		return jwt.NewValidationError("", jwt.ValidationErrorExpired)
	}
	if !c.VerifyIssuer(config.JWT.Issuer, true) {
		return jwt.NewValidationError("", jwt.ValidationErrorIssuer)
	}
	return nil
}

func Generate(userID int64, secret []byte) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(config.JWT.ExpireTime)
	claims := Claims{
		UserId: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    config.JWT.Issuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(secret)
	return token, err
}

func Parse(token string, secret []byte) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if tokenClaims != nil && tokenClaims.Valid {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}

	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors^jwt.ValidationErrorExpired == 0 {
			return nil, errno.AuthorizationExpired
		}
		if ve.Errors^jwt.ValidationErrorIssuer == 0 {
			return nil, errno.AuthorizationFailed
		}
	}

	return nil, errno.New(errno.AuthorizationFailedCode, err.Error())
}
