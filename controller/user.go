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

// Register @description 用户注册接口
// @author zkp15
// @date 2023/8/16 15:03
func Register(c *gin.Context) {
	// 数据获取校验
	var registerRequest user.RegisterRequest
	if err := c.ShouldBind(&registerRequest); err != nil || strs.IsAnyBlank(registerRequest.Username, registerRequest.Password) {
		render.Error(c, constants.UserRegisterErrorCode, "input params error")
		return
	}

	// 用户注册
	userInfo, err := service.SaveUser(registerRequest.Username, registerRequest.Password)
	if err != nil {
		render.Error(c, constants.UserRegisterErrorCode, err.Error())
		return
	}

	// 颁发令牌
	token, err := token.CreateToken(userInfo.Name, userInfo.ID)
	if err != nil {
		logrus.Error("token create error", err)
		render.Error(c, constants.UserRegisterErrorCode, "token create error")
		return
	}

	//封装返回
	render.Response(c, user.RegisterResponse{
		Token:  token,
		UserID: userInfo.ID,
	})
}

// Login @description 用户登录
// @author zkp15
// @date 2023/8/16 15:11
func Login(c *gin.Context) {
	// 数据校验
	var loginRequest user.LoginRequest
	if err := c.ShouldBind(&loginRequest); err != nil || strs.IsAnyBlank(loginRequest.Username, loginRequest.Password) {
		render.Error(c, constants.UserLoginErrorCode, "input params error")
		return
	}

	// 查询数据库
	userInfo, err := service.GetUserAndValidate(loginRequest.Username, loginRequest.Password)
	if err != nil {
		render.Error(c, constants.UserLoginErrorCode, err.Error())
		return
	}

	// 颁发令牌
	tokens, err := token.CreateToken(userInfo.Name, userInfo.ID)
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

// UserInfo @description 用户信息
// @author zkp15
// @date 2023/8/16 15:12
func UserInfo(c *gin.Context) {
	var userRequest user.InfoRequest

	// 数据校验
	if err := c.ShouldBind(&userRequest); err != nil || strs.IsAnyBlank(userRequest.Token) || userRequest.UserID <= 0 {
		render.Error(c, constants.UserInfoErrorCode, "input params error")
		return
	}

	// 解析token
	userClaims, err := token.ParseToken(userRequest.Token)
	if err != nil {
		render.Error(c, constants.UserInfoErrorCode, "token parse error")
		return
	}

	// 查询用户
	userInfo, err := service.GetUserById(userRequest.UserID)
	if err != nil {
		render.Error(c, constants.UserInfoErrorCode, "user not found")
		return
	}

	render.Response(c, user.InfoResponse{
		User: render.UserConverter(userInfo, service.CheckIsFollow(userClaims.UserId, userInfo.ID)),
	})
}
