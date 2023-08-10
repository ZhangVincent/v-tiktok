// @description 用户登录
// @author zkp15
// @date 2023/8/10 9:34

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"v-tiktok/model/constants"
	"v-tiktok/model/user"
	"v-tiktok/pkg/render"
	"v-tiktok/pkg/strs"
	"v-tiktok/pkg/token"
	"v-tiktok/service"
)

func Register(c *gin.Context) {
	var registerRequest user.RegisterRequest

	// 数据校验
	if err := c.ShouldBind(&registerRequest); err != nil || strs.IsAnyBlank(registerRequest.Username, registerRequest.Password) {
		render.Error(c, constants.UserRegisterErrorCode, "input params miss")
		return
	}

	// 保存数据库
	userInfo, err := service.SaveUser(registerRequest)
	if err != nil {
		render.Error(c, constants.UserRegisterErrorCode, err.Error())
		return
	}

	// 颁发令牌
	token, err := token.CreateToken(userInfo.Name, strs.ItoA(userInfo.ID))
	if err != nil {
		logrus.Error("token create error", err)
		render.Error(c, constants.UserRegisterErrorCode, "token create error")
		return
	}

	render.Response(c, user.RegisterResponse{
		Token:  token,
		UserID: userInfo.ID,
	})
}

func Login(c *gin.Context) {
	var loginRequest user.LoginRequest

	// 数据校验
	if err := c.ShouldBind(&loginRequest); err != nil || strs.IsAnyBlank(loginRequest.Username, loginRequest.Password) {
		render.Error(c, constants.UserLoginErrorCode, "input params miss")
		return
	}

	// 查询数据库
	userInfo, err := service.GetUser(loginRequest)
	if err != nil {
		render.Error(c, constants.UserLoginErrorCode, err.Error())
		return
	}

	// 颁发令牌
	tokens, err := token.CreateToken(userInfo.Name, strs.ItoA(userInfo.ID))
	if err != nil {
		logrus.Error("token create error", err)
		render.Error(c, constants.UserLoginErrorCode, "token create error")
		return
	}

	render.Response(c, user.LoginResponse{
		Token:  tokens,
		UserID: userInfo.ID,
	})
}

func UserInfo(c *gin.Context) {
	var userRequest user.InfoRequest

	// 数据校验
	if err := c.ShouldBind(&userRequest); err != nil || strs.IsAnyBlank(userRequest.Token, userRequest.UserID) {
		render.Error(c, constants.UserInfoErrorCode, "input params miss")
		return
	}

	// 解析token
	userClaims, err := token.ParseToken(userRequest.Token)
	if err != nil {
		render.Error(c, constants.UserInfoErrorCode, "token parse error")
		return
	}

	// 根据token查询user
	userInfo, err := service.GetUserByName(userClaims.Name)
	if err != nil || strs.IEqualsA(userInfo.ID, userRequest.UserID) {
		render.Error(c, constants.UserInfoErrorCode, "user not found")
		return
	}

	render.Response(c, user.InfoResponse{
		User: user.User{
			FollowCount:   userInfo.FollowCount,
			FollowerCount: userInfo.FollowerCount,
			ID:            userInfo.ID,
			IsFollow:      userInfo.IsFollow,
			Name:          userInfo.Name,
		},
	})
}
