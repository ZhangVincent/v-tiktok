// @description 消息
// @author zkp15
// @date 2023/8/12 16:45

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"v-tiktok/model/constants"
	"v-tiktok/model/message"
	"v-tiktok/pkg/render"
	"v-tiktok/pkg/strs"
	"v-tiktok/pkg/token"
	"v-tiktok/service"
)

// MessageAction @description 聊天记录
// @author zkp15
// @date 2023/8/16 18:43
func MessageAction(c *gin.Context) {
	//获取并校验数据
	var request message.RelationActionRequest
	if err := c.ShouldBind(&request); err != nil || strs.IsAnyBlank(request.Token, request.ActionType, request.Content) || request.ToUserID <= 0 {
		render.Error(c, constants.MessageErrorCode, "input params error")
		return
	}

	//解析token
	userClaims, err := token.ParseToken(request.Token)
	if err != nil {
		logrus.Error("token parse error")
		render.Error(c, constants.MessageErrorCode, "token parse error")
		return
	}

	//插入消息记录
	if request.ActionType == "1" {
		if err := service.SaveMessage(userClaims.UserId, request.ToUserID, request.Content); err != nil {
			render.Error(c, constants.MessageErrorCode, err.Error())
		}
	} else {
		render.Error(c, constants.MessageErrorCode, "action type error")
	}

	//封装返回
	render.Response(c, message.RelationActionResponse{})
}

// MessageChat @description 消息操作
// @author zkp15
// @date 2023/8/16 18:45
func MessageChat(c *gin.Context) {
	//获取并校验数据
	var request message.ChatRequest
	if err := c.ShouldBind(&request); err != nil || strs.IsBlank(request.Token) || request.ToUserID <= 0 {
		render.Error(c, constants.MessageErrorCode, "input params error")
		return
	}

	//解析token
	userClaims, err := token.ParseToken(request.Token)
	if err != nil {
		logrus.Error("token parse error")
		render.Error(c, constants.MessageErrorCode, "token parse error")
		return
	}

	//查询消息记录
	messages, err := service.GetMessage(userClaims.UserId, request.ToUserID, request.PreMsgTime)
	if err != nil {
		render.Error(c, constants.MessageErrorCode, err.Error())
	}

	//封装返回
	messageList := make([]message.Message, len(messages))
	for i, x := range messages {
		messageList[i] = render.MessageConverter(x)
	}

	render.Response(c, message.ChatResponse{
		MessageList: messageList,
	})
}
