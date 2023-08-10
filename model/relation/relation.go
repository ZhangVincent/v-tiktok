// @description TODO
// @author zkp15
// @date 2023/8/10 21:27
// @version 1.0

package relation

import "v-tiktok/model/user"

type ActionRequest struct {
	ActionType string `json:"action_type"` // 1-关注，2-取消关注
	ToUserID   string `json:"to_user_id"`  // 对方用户id
	Token      string `json:"token"`       // 用户鉴权token
}

type ActionResponse struct {
	Response user.Response `json:"response"` // 状态码
}

type FollowListRequest struct {
	Token  string `json:"token"`   // 用户鉴权token
	UserID string `json:"user_id"` // 用户id
}

type FollowListResponse struct {
	Response user.Response `json:"response"`  // 状态码
	UserList []user.User   `json:"user_list"` // 用户信息列表
}

type FollowerListRequest struct {
	Token  string `json:"token"`   // 用户鉴权token
	UserID string `json:"user_id"` // 用户id
}

type FollowerListResponse struct {
	Response user.Response `json:"response"`  // 状态码
	UserList []user.User   `json:"user_list"` // 用户列表
}

type FriendListRequest struct {
	Token  string `json:"token"`   // 用户鉴权token
	UserID string `json:"user_id"` // 用户id
}

type FriendListResponse struct {
	Response user.Response `json:"response"`  // 状态码
	UserList []user.User   `json:"user_list"` // 用户列表
}
