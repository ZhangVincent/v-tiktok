// @description TODO
// @author zkp15
// @date 2023/8/10 21:26
// @version 1.0

package comment

import "v-tiktok/model/user"

type ActionRequest struct {
	ActionType  string `json:"action_type"`            // 1-发布评论，2-删除评论
	CommentID   string `json:"comment_id,omitempty"`   // 要删除的评论id，在action_type=2的时候使用
	CommentText string `json:"comment_text,omitempty"` // 用户填写的评论内容，在action_type=1的时候使用
	Token       string `json:"token"`                  // 用户鉴权token
	VideoID     string `json:"video_id"`               // 视频id
}

type ActionResponse struct {
	Response user.Response `json:"response"`          // 状态码
	Comment  Comment       `json:"comment,omitempty"` // 评论成功返回评论内容，不需要重新拉取整个列表
}

type ListRequest struct {
	Token   string `json:"token"`    // 用户鉴权token
	VideoID string `json:"video_id"` // 视频id
}

type ListResponse struct {
	Response    user.Response `json:"response"`               // 状态码
	CommentList []Comment     `json:"comment_list,omitempty"` // 评论列表
}

// Comment
type Comment struct {
	Content    string    `json:"content"`     // 评论内容
	CreateDate string    `json:"create_date"` // 评论发布日期，格式 mm-dd
	ID         int64     `json:"id"`          // 评论id
	User       user.User `json:"user"`        // 评论用户信息
}

type MessageSendEvent struct {
	UserId     int64  `json:"user_id,omitempty"`
	ToUserId   int64  `json:"to_user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}

type MessagePushEvent struct {
	FromUserId int64  `json:"user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}
