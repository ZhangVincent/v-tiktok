// @description user业务逻辑
// @author zkp15
// @date 2023/8/10 15:28
// @version 1.0

package service

import (
	"errors"
	"github.com/sirupsen/logrus"
	"v-tiktok/model"
	"v-tiktok/pkg/bcrypt"
	"v-tiktok/pkg/redis"
	"v-tiktok/pkg/sqls"
	"v-tiktok/pkg/strs"
	"v-tiktok/repository"
	"v-tiktok/repository/redisDao"
)

func SaveUser(username, password string) (model.User, error) {
	// 数据校验
	if strs.RuneLen(username) > 32 || strs.RuneLen(password) > 32 {
		return model.User{}, errors.New("length of username and password should not over 32")
	}

	// 查询用户名是否已注册
	if _, err := repository.GetUserByName(sqls.DB(), username); err == nil {
		return model.User{}, errors.New("username already exits")
	}

	// 保存数据库
	userInfo, err := repository.SaveUser(sqls.DB(), username, bcrypt.EncodePassword(password))
	if err != nil {
		logrus.Error("save user error", err.Error())
		return userInfo, err
	}
	logrus.Info("new user create success: ", userInfo.Name)

	//将用户保存在redis中
	err = redisDao.SaveUser(redis.Client(), userInfo)
	if err != nil {
		logrus.Error("save user to redis error", err.Error())
	}

	return userInfo, nil
}

func GetUserAndValidate(username, password string) (model.User, error) {
	if strs.RuneLen(username) > 32 || strs.RuneLen(password) > 32 {
		return model.User{}, errors.New("length of username and password should not over 32")
	}

	//从数据库中查询用户
	userInfo, err := repository.GetUserByName(sqls.DB(), username)
	if err != nil {
		return userInfo, err
	}

	// 校验用户身份
	if err = bcrypt.ValidatePassword(userInfo.Password, password); err != nil {
		return userInfo, errors.New("password error")
	}

	return userInfo, nil
}

func GetUserById(id int64) (model.User, error) {
	if id <= 0 {
		return model.User{}, errors.New("input params error")
	}
	//从redis中查询用户
	userInfo, err := redisDao.GetUserById(redis.Client(), id)
	if err == nil {
		return userInfo, nil
	}
	//从数据库中查询用户
	userInfo, err = repository.GetUserById(sqls.DB(), id)
	if err != nil {
		return userInfo, err
	}
	//将用户保存在redis中
	err = redisDao.SaveUser(redis.Client(), userInfo)
	if err != nil {
		logrus.Error("save user to redis error", err.Error())
	}
	return userInfo, nil
}
