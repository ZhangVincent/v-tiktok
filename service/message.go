// @description 消息业务
// @author zkp15
// @date 2023/8/14 21:33
// @version 1.0

package service

import (
	"errors"
	"github.com/sirupsen/logrus"
	"v-tiktok/model"
	"v-tiktok/pkg/forbiddenWordCheck"
	"v-tiktok/pkg/sqls"
	"v-tiktok/pkg/strs"
	"v-tiktok/repository"
)

func SaveMessage(fromUserId, toUserId int64, content string) error {
	db := sqls.DB()
	//校验数据
	if fromUserId <= 0 || toUserId <= 0 || strs.IsBlank(content) {
		return errors.New("input params error")
	}
	if !repository.UserExist(db, fromUserId) || !repository.UserExist(db, toUserId) {
		return errors.New("user not found")
	}

	//对消息做违禁词检查
	hitWords := forbiddenWordCheck.Check(&content)
	if len(hitWords) != 0 {
		logrus.Errorf("user (id: %v) to (id: %v) message (%v) contains forbidden words (%v)", fromUserId, toUserId, content, hitWords)
	}

	//保存消息
	if err := repository.SaveMessage(db, fromUserId, toUserId, content); err != nil {
		return errors.New("save message error")
	}

	return nil
}

func GetMessage(fromUserId, toUserId, preMsgTime int64) ([]model.Message, error) {
	db := sqls.DB()
	//校验数据
	if fromUserId <= 0 || toUserId <= 0 {
		return nil, errors.New("input params error")
	}
	if !repository.UserExist(db, fromUserId) || !repository.UserExist(db, toUserId) {
		return nil, errors.New("user not found")
	}

	//查询消息记录
	//messages, err := repository.GetMessage(db, fromUserId, toUserId, preMsgTime)
	messages, err := repository.GetMessage(db, fromUserId, toUserId, 0)
	if err != nil {
		return nil, errors.New("get message error")
	}

	return messages, nil
}
