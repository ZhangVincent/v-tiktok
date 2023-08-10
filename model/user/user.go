// @description 用户模型
// @author zkp15
// @date 2023/8/9 22:57
// @version 1.0

package user

type Response struct {
	StatusCode int    `json:"status_code"`          // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg,omitempty"` // 返回状态描述
}

type RegisterRequest struct {
	Password string ` json:"password" form:"password" binding:"required"` // 密码，最长32个字符
	Username string ` json:"username" form:"username" binding:"required"` // 注册用户名，最长32个字符
}

type RegisterResponse struct {
	Response Response `json:"response"`          // 状态码
	UserID   int64    `json:"user_id,omitempty"` // 用户id
	Token    string   `json:"token"`             // 用户鉴权token
}

type LoginRequest struct {
	Password string ` json:"password" form:"password" binding:"required"` // 密码，最长32个字符
	Username string `json:"username" form:"username" binding:"required"`  // 注册用户名，最长32个字符
}

type LoginResponse struct {
	Response Response `json:"response"`          // 状态码
	UserID   int64    `json:"user_id,omitempty"` // 用户id
	Token    string   `json:"token"`             // 用户鉴权token
}

type InfoRequest struct {
	Token  string `json:"token" form:"token" binding:"required"`      // 用户鉴权token
	UserID string `json:"user_id"  form:"user_id" binding:"required"` // 用户id
}

type InfoResponse struct {
	Response Response `json:"response"` // 状态码
	User     User     `json:"user"`     // 用户信息
}

type User struct {
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
}
