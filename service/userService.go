// @description user业务逻辑
// @author zkp15
// @date 2023/8/10 15:28
// @version 1.0

package service

import (
	"errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
	"v-tiktok/model"
	"v-tiktok/model/user"
	"v-tiktok/pkg/bcrypt"
	"v-tiktok/pkg/sqls"
	"v-tiktok/pkg/strs"
	"v-tiktok/repository"
)

func SaveUser(registerRequest user.RegisterRequest) (*model.User, error) {
	// 数据校验
	if strs.RuneLen(registerRequest.Username) > 32 || strs.RuneLen(registerRequest.Password) > 32 {
		return nil, errors.New("length of username and password should not over 32")
	}

	// 查询用户名是否已注册
	if _, err := repository.GetUserByName(sqls.DB(), registerRequest.Username); err == nil {
		return nil, errors.New("username already exits")
	}

	// 保存数据库
	userInfo, err := repository.SaveUser(sqls.DB(), model.User{
		GormModel: model.GormModel{
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		Name:     registerRequest.Username,
		Password: bcrypt.EncodePassword(registerRequest.Password),
	})
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	logrus.Info("new user create success: ", userInfo.Name)

	return userInfo, nil
}

func GetUser(loginRequest user.LoginRequest) (*model.User, error) {
	if strs.RuneLen(loginRequest.Username) > 32 || strs.RuneLen(loginRequest.Password) > 32 {
		return nil, errors.New("length of username and password should not over 32")
	}

	// 校验用户身份
	userInfo, err := repository.GetUserByName(sqls.DB(), loginRequest.Username)
	if err != nil {
		return nil, err
	}
	if err = bcrypt.ValidatePassword(userInfo.Password, loginRequest.Password); err != nil {
		return nil, errors.New("password error")
	}

	return userInfo, nil
}

func GetUserByName(name string) (*model.User, error) {
	userInfo, err := repository.GetUserByName(sqls.DB(), name)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return userInfo, nil
}
