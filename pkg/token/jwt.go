// @description 用户token令牌
// @author zkp15
// @date 2023/8/9 11:23
// @version 1.0

package token

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
	"v-tiktok/model/config"
	"v-tiktok/model/constants"
)

type UserClaims struct {
	*jwt.RegisteredClaims

	Name   string `json:"name"`
	UserId int64  `json:"user_id"`
}

func CreateToken(name string, userId int64) (string, error) {
	var (
		jwtConf   = config.Instance.Jwt
		expiredAt = time.Now().Add(time.Duration(jwtConf.ExpireDays) * 24 * 60 * 60 * time.Second)
	)
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, &UserClaims{
		RegisteredClaims: &jwt.RegisteredClaims{
			Issuer:    jwtConf.Issuer,
			ExpiresAt: jwt.NewNumericDate(expiredAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		Name:   name,
		UserId: userId,
	})
	return claims.SignedString([]byte(jwtConf.SignKey))
}

func ParseToken(tokenString string) (*UserClaims, error) {
	jwtConf := config.Instance.Jwt
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(jwtConf.SignKey), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, constants.MalformedErr
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, constants.ExpiredErr
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, constants.NotValidYetErr
			} else {
				return nil, constants.InvalidErr
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, constants.InvalidErr
	} else {
		return nil, constants.InvalidErr
	}
}
