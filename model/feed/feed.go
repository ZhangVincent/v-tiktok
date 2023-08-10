// @description 视频流
// @author zkp15
// @date 2023/8/10 21:22
// @version 1.0

package feed

import (
	"v-tiktok/model/publish"
	"v-tiktok/model/user"
)

type FeedRequest struct {
	LatestTime string `json:"latest_time,omitempty"` // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
	Token      string `json:"token,omitempty"`       // 用户登录状态下设置
}

type FeedResponse struct {
	Response  user.Response   `json:"response"`             // 状态码
	VideoList []publish.Video `json:"video_list,omitempty"` // 视频列表
	NextTime  int64           `json:"next_time,omitempty"`  // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}
