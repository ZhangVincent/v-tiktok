// @description 常量
// @author zkp15
// @date 2023/8/9 10:51
// @version 1.0

package constants

import (
	"errors"
	"v-tiktok/model"
	"v-tiktok/model/user"
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
	RelationErrorCode         = 6004
	RelationFollowErrorCode   = 6002
	RelationFollowerErrorCode = 6003
	RelationFriendErrorCode   = 6004
	MessageErrorCode          = 7001
)

var (
	DefaultUser = user.User{
		ID:             0,
		Name:           "default_user",
		Signature:      "nothing is everything",
		FavoriteCount:  0,
		FollowCount:    0,
		FollowerCount:  0,
		IsFollow:       false,
		TotalFavorited: "",
		WorkCount:      0,
	}

	DefaultUserModel = model.User{
		Name:           "default_user",
		Signature:      "nothing is everything",
		FavoriteCount:  0,
		FollowCount:    0,
		FollowerCount:  0,
		TotalFavorited: "",
		WorkCount:      0,
	}
)

var (
	ForbiddenWords = []string{"操", "fuck", "日", "你妈", "淦", "逼", "杀", "死"}
)
