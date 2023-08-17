// @description 视频点赞
// @author zkp15
// @date 2023/8/10 21:25
// @version 1.0

package favorite

import (
	"v-tiktok/model/publish"
	"v-tiktok/model/user"
)

type ActionRequest struct {
	ActionType string `json:"action_type" form:"action_type" binding:"required"` // 1-点赞，2-取消点赞
	Token      string `json:"token" form:"token" binding:"required"`             // 用户鉴权token
	VideoID    int64  `json:"video_id" form:"video_id" binding:"required"`       // 视频id
}

type ActionResponse struct {
	Response user.Response `json:"response"` // 状态码
}

type ListRequest struct {
	Token  string `json:"token" form:"token" binding:"required"`     // 用户鉴权token
	UserID int64  `json:"user_id" form:"user_id" binding:"required"` // 用户id
}

type ListResponse struct {
	Response  user.Response   `json:"response"`   // 状态码
	VideoList []publish.Video `json:"video_list"` // 用户点赞视频列表
}
