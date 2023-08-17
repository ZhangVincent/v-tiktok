// @description 发布视频
// @author zkp15
// @date 2023/8/10 21:23
// @version 1.0

package publish

import (
	"mime/multipart"
	"v-tiktok/model/user"
)

type ActionRequest struct {
	Token string                `json:"token" form:"token" binding:"required"`
	Data  *multipart.FileHeader `json:"data" form:"data" binding:"required"`
	Title string                `json:"title" form:"title"`
}

type ActionResponse struct {
	Response user.Response `json:"response"` // 状态码
}

type ListRequest struct {
	Token  string `json:"token" form:"token" binding:"required"`     // 用户鉴权token
	UserID int64  `json:"user_id" form:"user_id" binding:"required"` // 用户id
}

type ListResponse struct {
	Response  user.Response `json:"response"`   // 状态码
	VideoList []Video       `json:"video_list"` // 用户发布的视频列表
}

// Video
type Video struct {
	Id            int64     `json:"id,omitempty"`
	Author        user.User `json:"author"`
	PlayUrl       string    ` json:"play_url,omitempty"`
	CoverUrl      string    `json:"cover_url,omitempty"`
	FavoriteCount int64     `json:"favorite_count,omitempty"`
	CommentCount  int64     `json:"comment_count,omitempty"`
	IsFavorite    bool      `json:"is_favorite,omitempty"`
	Title         string    `json:"title"`
}
