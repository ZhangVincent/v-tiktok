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
	&User{}, &Message{}, &Video{},
}

type GormModel struct {
	ID        int64          `gorm:"primaryKey;autoIncrement" json:"id" form:"id"`
	CreatedAt time.Time      `json:"createAt" form:"createAt"` // 创建时间
	UpdatedAt time.Time      `json:"updateAt" form:"updateAt"` // 更新时间
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
	IsFollow        bool   `json:"is_follow"`        // true-已关注，false-未关注
	TotalFavorited  string `json:"total_favorited"`  // 获赞数量
	WorkCount       int64  `json:"work_count"`       // 作品数
}

type Message struct {
	GormModel
	Content    string `json:"content"`      // 消息内容
	FromUserID int64  `json:"from_user_id"` // 消息发送者id
	ToUserID   int64  `json:"to_user_id"`   // 消息接收者id
}

type Video struct {
	GormModel
	AuthorId      int64  `json:"author_id"`
	PlayUrl       string ` json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
	Title         string `json:"title"`
}
