package service

import (
	"github.com/NoCLin/douyin-backend-go/model"
	"github.com/NoCLin/douyin-backend-go/repository"
)

func NewUserLoginInstance(username, password string) *UserLoginFlow {
	return &UserLoginFlow{
		username: username,
		password: password,
	}
}

type UserLoginFlow struct {
	username string
	password string
}

func (f *UserLoginFlow) Login() *model.UserLoginResponse {
	user, err := repository.NewUserDaoInstance().QueryUserByName(f.username)
	if err != nil {
		return &model.UserLoginResponse{
			Response: model.NewResponse(-1, err.Error()),
		}
	}
	if user.Password != f.password {
		return &model.UserLoginResponse{
			Response: model.NewResponse(-1, "password doesn't match"),
		}

	}
	// TODO token生成与校验

	return &model.UserLoginResponse{
		Response: model.NewResponse(0, "login success"),
		UserId:   user.Id,
		Token:    "",
	}
}
