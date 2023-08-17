// @description 数据库对象建模
// @author zkp15
// @date 2023/8/9 10:51
// @version 1.0

package model

import (
	"gorm.io/gorm"
	"time"
)

var Models = []interface{}{
	&User{}, &Message{}, &Video{}, &Comment{}, &Favorite{}, &Follow{},
}

type GormModel struct {
	ID        int64          `gorm:"primaryKey;autoIncrement" json:"id" form:"id"`
	CreatedAt time.Time      `json:"createAt" form:"createAt"` // 创建时间
	UpdatedAt time.Time      `json:"updateAt" form:"updateAt"` // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type GormTimeModel struct {
	ID        int64          `gorm:"primaryKey;autoIncrement" json:"id" form:"id"`
	CreatedAt int64          `json:"createAt" form:"createAt"` // 创建时间
	UpdatedAt int64          `json:"updateAt" form:"updateAt"` // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type User struct {
	GormModel
	Name            string `json:"name"`             // 用户名称
	Password        string `json:"password"`         // 密码
	Avatar          string `json:"avatar"`           // 用户头像
	BackgroundImage string `json:"background_image"` // 用户个人页顶部大图
	Signature       string `json:"signature"`        // 个人简介
	FavoriteCount   int64  `json:"favorite_count"`   // 喜欢数
	FollowCount     int64  `json:"follow_count"`     // 关注总数
	FollowerCount   int64  `json:"follower_count"`   // 粉丝总数
	//IsFollow        bool   `json:"is_follow"`        // true-已关注，false-未关注
	TotalFavorited string `json:"total_favorited"` // 获赞数量
	WorkCount      int64  `json:"work_count"`      // 作品数
}

type Message struct {
	GormTimeModel
	Content    string `json:"content"`      // 消息内容
	FromUserID int64  `json:"from_user_id"` // 消息发送者id
	ToUserID   int64  `json:"to_user_id"`   // 消息接收者id
}

type Video struct {
	GormTimeModel
	AuthorId      int64  `json:"author_id"`
	PlayUrl       string ` json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	//IsFavorite    bool   `json:"is_favorite"`
	Title   string `json:"title"`
	IsAudit int64  `json:"is_audit"` // 0-不在审核状态 1-在审核状态
}

type Comment struct {
	GormModel
	Content string `json:"content"`  // 评论内容
	UserId  int64  `json:"user_id"`  // 评论用户信息
	VideoId int64  `json:"video_id"` // 视频信息
}

type Favorite struct {
	GormModel
	UserId  int64 `json:"user_id"`  // 点赞的用户信息
	VideoId int64 `json:"video_id"` // 点赞的视频信息
}

type Follow struct {
	GormModel
	FromUserId int64 `json:"from_user_id"` // 关注的用户信息
	ToUserId   int64 `json:"to_user_id"`   // 被关注的用户信息
}
