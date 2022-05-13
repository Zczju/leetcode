package service

import (
	"github.com/NoCLin/douyin-backend-go/model"
	"github.com/NoCLin/douyin-backend-go/repository"
)

func NewGetUserInfoFlowInstance(queryUserId int64) *GetUserInfoFlow {
	return &GetUserInfoFlow{
		UserId: queryUserId,
	}
}

type GetUserInfoFlow struct {
	UserId int64
}

func (f *GetUserInfoFlow) GetUserInfo() *model.UserResponse {
	user, err := repository.NewUserDaoInstance().QueryUserById(f.UserId)
	if err != nil {
		return &model.UserResponse{
			Response: model.NewResponse(-1, "fail to find user"),
		}
	}
	// todo 获取follow情况
	return &model.UserResponse{
		Response: model.NewResponse(0, "success to find user"),
		User: model.UserInfo{
			User:          *user,
			FollowCount:   0,
			FollowerCount: 0,
			IsFollow:      false,
		},
	}

}
