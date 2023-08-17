// @description 用户关系业务操作
// @author zkp15
// @date 2023/8/12 16:54
// @version 1.0

package service

import (
	"errors"
	"gorm.io/gorm"
	"v-tiktok/model"
	"v-tiktok/model/relation"
	"v-tiktok/pkg/render"
	"v-tiktok/pkg/sqls"
	"v-tiktok/pkg/strs"
	"v-tiktok/repository"
)

func SaveRelation(fromUserId, toUserId int64) error {
	//数据校验
	if !repository.UserExist(sqls.DB(), fromUserId) || !repository.UserExist(sqls.DB(), toUserId) {
		return errors.New("input params error")
	}

	//数据库操作
	if err := sqls.DB().Transaction(func(tx *gorm.DB) error {
		//关系表中是否存在记录
		if isFollow := repository.RelationExist(tx, fromUserId, toUserId); isFollow {
			return errors.New("follow relation already exit")
		}
		//关注者关注数+1
		if err := repository.UpdateUserFollow(tx, fromUserId, 1); err != nil {
			return err
		}
		//被关注者粉丝数+1
		if err := repository.UpdateUserFollower(tx, toUserId, 1); err != nil {
			return err
		}
		//将关注信息插入follow表中
		if err := repository.SaveRelation(tx, fromUserId, toUserId); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}

func DeleteRelation(fromUserId, toUserId int64) error {
	//数据校验
	if !repository.UserExist(sqls.DB(), fromUserId) || !repository.UserExist(sqls.DB(), toUserId) {
		return errors.New("input params error")
	}

	//数据库操作
	if err := sqls.DB().Transaction(func(tx *gorm.DB) error {
		//关系表中是否存在记录
		if isFollow := repository.RelationExist(tx, fromUserId, toUserId); !isFollow {
			return errors.New("you not follow yet")
		}
		//关注者关注数-1
		if err := repository.UpdateUserFollow(tx, fromUserId, -1); err != nil {
			return err
		}
		//被关注者粉丝数-1
		if err := repository.UpdateUserFollower(tx, toUserId, -1); err != nil {
			return err
		}
		//将关注信息从follow表中删除
		if err := repository.DeleteRelation(tx, fromUserId, toUserId); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}

func GetFollowList(fromUserId int64) ([]model.User, error) {
	//数据校验
	if fromUserId <= 0 {
		return nil, errors.New("input params error")
	}

	//查询关注的信息
	follows, err := repository.GetFollowsByFromUserId(sqls.DB(), fromUserId)
	if err != nil {
		return nil, err
	}

	n := len(follows)
	if n == 0 {
		return nil, nil
	}

	//获取被关注的用户id
	ids := make([]int64, n)
	for i, f := range follows {
		ids[i] = f.ToUserId
	}

	//根据用户id查询用户信息
	users, err := repository.GetUserByIds(sqls.DB(), ids)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func GetFollowerList(toUserId int64) ([]model.User, error) {
	//数据校验
	if toUserId <= 0 {
		return nil, errors.New("input params error")
	}

	//查询粉丝的信息
	follows, err := repository.GetFollowersByToUserId(sqls.DB(), toUserId)
	if err != nil {
		return nil, err
	}

	n := len(follows)
	if n == 0 {
		return nil, nil
	}

	//获取粉丝的用户id
	ids := make([]int64, n)
	for i, f := range follows {
		ids[i] = f.FromUserId
	}

	//根据用户id查询用户信息
	users, err := repository.GetUserByIds(sqls.DB(), ids)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func GetFriendList(toUserId int64) ([]relation.FriendUserList, error) {
	//查询粉丝
	followUsers, err := GetFollowList(toUserId)
	if err != nil {
		return nil, err
	}

	//查询最近发消息的朋友和最新消息记录
	friendUsers := make([]relation.FriendUserList, len(followUsers))
	for i, user := range followUsers {
		isFollow := repository.RelationExist(sqls.DB(), toUserId, user.ID)
		messages1, _ := repository.GetLatestMessages(sqls.DB(), user.ID, toUserId)
		messages2, _ := repository.GetLatestMessages(sqls.DB(), toUserId, user.ID)
		b1 := strs.IsNotBlank(messages1.Content)
		b2 := strs.IsNotBlank(messages2.Content)
		if b1 && b2 {
			if messages1.CreatedAt > messages2.CreatedAt {
				//message1是最新消息
				friendUsers[i] = render.FriendConverter(user, isFollow, messages1.Content, 0)
			} else {
				//message2是最新消息
				friendUsers[i] = render.FriendConverter(user, isFollow, messages2.Content, 1)
			}
		} else if b1 {
			//用户接收到的消息
			friendUsers[i] = render.FriendConverter(user, isFollow, messages1.Content, 0)
		} else if b2 {
			//用户发送的消息
			friendUsers[i] = render.FriendConverter(user, isFollow, messages2.Content, 1)
		} else {
			//没有消息
			friendUsers[i] = render.FriendConverter(user, isFollow, "", 0)
		}
	}

	return friendUsers, nil
}

func CheckIsFollow(fromUserId, toUserId int64) bool {
	if fromUserId <= 0 || !repository.UserExist(sqls.DB(), fromUserId) || toUserId <= 0 || !repository.UserExist(sqls.DB(), toUserId) {
		return false
	}
	return repository.RelationExist(sqls.DB(), fromUserId, toUserId)
}
