// @description 用户数据库查询
// @author zkp15
// @date 2023/8/9 23:28
// @version 1.0

package repository

import (
	"errors"
	"gorm.io/gorm"
	"v-tiktok/model"
)

func GetUserByName(db *gorm.DB, name string) (*model.User, error) {
	var user model.User
	result := db.Where("name = ?", name).First(&user)
	if result.RowsAffected <= 0 {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func SaveUser(db *gorm.DB, user model.User) (*model.User, error) {
	result := db.Create(&user)
	if result.RowsAffected <= 0 {
		return nil, errors.New("save user error")
	}
	return &user, nil
}
