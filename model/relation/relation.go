// @description TODO
// @author zkp15
// @date 2023/8/10 21:27
// @version 1.0

package relation

import "v-tiktok/model/user"

type ActionRequest struct {
	ActionType string `json:"action_type" form:"action_type" binding:"required"` // 1-关注，2-取消关注
	ToUserID   int64  `json:"to_user_id" form:"to_user_id" binding:"required"`   // 对方用户id
	Token      string `json:"token" form:"token" binding:"required"`             // 用户鉴权token
}

type ActionResponse struct {
	Response user.Response `json:"response"` // 状态码
}

type FollowListRequest struct {
	Token  string `json:"token" form:"token" binding:"required"`     // 用户鉴权token
	UserID int64  `json:"user_id" form:"user_id" binding:"required"` // 用户id
}

type FollowListResponse struct {
	Response user.Response `json:"response"`  // 状态码
	UserList []user.User   `json:"user_list"` // 用户信息列表
}

type FollowerListRequest struct {
	Token  string `json:"token" form:"token" binding:"required"`     // 用户鉴权token
	UserID int64  `json:"user_id" form:"user_id" binding:"required"` // 用户id
}

type FollowerListResponse struct {
	Response user.Response `json:"response"`  // 状态码
	UserList []user.User   `json:"user_list"` // 用户列表
}

type FriendListRequest struct {
	Token  string `json:"token" form:"token" binding:"required"`     // 用户鉴权token
	UserID int64  `json:"user_id" form:"user_id" binding:"required"` // 用户id
}

type FriendListResponse struct {
	Response user.Response    `json:"response"`  // 状态码
	UserList []FriendUserList `json:"user_list"` // 用户列表
}

type FriendUserList struct {
	Avatar          string `json:"avatar,omitempty"`           // 用户头像
	BackgroundImage string `json:"background_image,omitempty"` // 用户个人页顶部大图
	FavoriteCount   int64  `json:"favorite_count,omitempty"`   // 喜欢数
	FollowCount     int64  `json:"follow_count,omitempty"`     // 关注总数
	FollowerCount   int64  `json:"follower_count,omitempty"`   // 粉丝总数
	ID              int64  `json:"id,omitempty"`               // 用户id
	IsFollow        bool   `json:"is_follow,omitempty"`        // true-已关注，false-未关注
	Name            string `json:"name,omitempty"`             // 用户名称
	Signature       string `json:"signature,omitempty"`        // 个人简介
	TotalFavorited  string `json:"total_favorited,omitempty"`  // 获赞数量
	WorkCount       int64  `json:"work_count,omitempty"`       // 作品数
	Message         string `json:"message,omitempty"`          // 和该好友的最新聊天消息
	MsgType         int64  `json:"msg_type"`                   // message消息的类型，0 => 当前请求用户接收的消息， 1 => 当前请求用户发送的消息
}

//type FriendUserList struct {
//	User    user.User `json:"user"`
//	Message string    `json:"message,omitempty"` // 和该好友的最新聊天消息
//	MsgType int64     `json:"msg_type"`          // message消息的类型，0 => 当前请求用户接收的消息， 1 => 当前请求用户发送的消息
//}
