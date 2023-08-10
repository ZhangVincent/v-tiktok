// @description TODO
// @author zkp15
// @date 2023/8/10 21:28
// @version 1.0

package message

import "v-tiktok/model/user"

type RelationActionRequest struct {
	ActionType string `json:"action_type"` // 1-发送消息
	Content    string `json:"content"`     // 消息内容
	ToUserID   string `json:"to_user_id"`  // 对方用户id
	Token      string `json:"token"`       // 用户鉴权token
}

type RelationActionResponse struct {
	Response user.Response `json:"response"` // 状态码
}

type ChatRequest struct {
	ToUserID   string `json:"to_user_id"`   // 对方用户id
	Token      string `json:"token"`        // 用户鉴权token
	PreMsgTime int64  `json:"pre_msg_time"` //上次最新消息的时间
}

type ChatResponse struct {
	Response    user.Response `json:"response"`     // 状态码
	MessageList []Message     `json:"message_list"` // 用户列表
}

// Message
type Message struct {
	ID         int64  `json:"id"`           // 消息id
	Content    string `json:"content"`      // 消息内容
	CreateTime int64  `json:"create_time"`  // 消息发送时间 yyyy-MM-dd HH:MM:ss
	FromUserID int64  `json:"from_user_id"` // 消息发送者id
	ToUserID   int64  `json:"to_user_id"`   // 消息接收者id
}
