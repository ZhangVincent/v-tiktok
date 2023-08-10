// @description 常量
// @author zkp15
// @date 2023/8/9 10:51
// @version 1.0

package constants

import (
	"errors"
)

var (
	TokenNotFoundErr = errors.New("token not found")
	ExpiredErr       = errors.New("token is expired")
	NotValidYetErr   = errors.New("token not active yet")
	MalformedErr     = errors.New("that's not even a token")
	InvalidErr       = errors.New("couldn't handle this token")
)

const (
	UserTokenHeader = "Authorization"
	UserTokenParam  = "_user_token"
	UserIdKey       = "__user_id"
	UserKey         = "__user"
)

const (
	FeedErrorCode             = 1001
	UserRegisterErrorCode     = 2001
	UserLoginErrorCode        = 2002
	UserInfoErrorCode         = 2003
	PublishErrorCode          = 3001
	FavoriteErrorCode         = 4001
	CommentErrorCode          = 5001
	RelationFollowErrorCode   = 6001
	RelationFollowerErrorCode = 6002
	RelationErrorCode         = 6003
	MessageErrorCode          = 7001
)
