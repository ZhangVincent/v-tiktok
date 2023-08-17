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

func UserExist(db *gorm.DB, userId int64) bool {
	var count int64
	db.Model(&model.User{}).Where("id = ?", userId).Count(&count)
	return count > 0
}

func GetUserByName(db *gorm.DB, name string) (model.User, error) {
	var user model.User
	result := db.Where("name = ?", name).First(&user)
	if result.RowsAffected <= 0 {
		return user, errors.New("user not found")
	}
	return user, nil
}

func GetUserById(db *gorm.DB, id int64) (model.User, error) {
	var user model.User
	result := db.First(&user, id)
	if result.RowsAffected <= 0 {
		return user, errors.New("user not found")
	}
	return user, nil
}

func GetUserByIds(db *gorm.DB, ids []int64) ([]model.User, error) {
	var users []model.User
	result := db.Find(&users, ids)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func SaveUser(db *gorm.DB, username, password string) (model.User, error) {
	user := model.User{
		Name:     username,
		Password: password,
	}
	result := db.Create(&user)
	if result.RowsAffected <= 0 {
		return user, errors.New("save user error")
	}
	return user, nil
}

func UpdateUserFavorite(db *gorm.DB, userId int64, changeScore int64) error {
	return db.Exec("update t_user set favorite_count = favorite_count + ? where id = ?", changeScore, userId).Error
}

func UpdateUserFavorited(db *gorm.DB, userId int64, changeScore int64) error {
	return db.Exec("update t_user set total_favorited = total_favorited + ? where id = ?", changeScore, userId).Error
}

func UpdateUserFollow(db *gorm.DB, fromUserId, changeScore int64) error {
	return db.Exec("update t_user set follow_count = follow_count + ? where id = ?", changeScore, fromUserId).Error
}

func UpdateUserFollower(db *gorm.DB, toUserId, changeScore int64) error {
	return db.Exec("update t_user set follower_count = follower_count + ? where id = ?", changeScore, toUserId).Error
}

func UpdateUserWorkCount(db *gorm.DB, userId, changeScore int64) error {
	return db.Exec("update t_user set work_count = work_count + ? where id = ?", changeScore, userId).Error
}
