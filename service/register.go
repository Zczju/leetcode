package service

import (
	"github.com/NoCLin/douyin-backend-go/model"
	"github.com/NoCLin/douyin-backend-go/repository"
)

func NewUserRegisterFlow(username, password string) *UserRegisterFlow {
	return &UserRegisterFlow{
		username: username,
		password: password,
	}
}

type UserRegisterFlow struct {
	username string
	password string

	userId int64
}

func (f *UserRegisterFlow) Register() *model.UserRegisterResponse {
	user, err := repository.NewUserDaoInstance().QueryUserByName(f.username)
	if err != nil {
		return &model.UserRegisterResponse{
			Response: model.NewResponse(-1, err.Error()),
		}
	}
	if user != nil {
		return &model.UserRegisterResponse{
			Response: model.NewResponse(-1, "username already exists"),
		}
	}

	user = &model.User{}
	user.Name = f.username
	user.Password = f.password
	newUser, err := repository.NewUserDaoInstance().InsertUser(user)
	if err != nil {
		return &model.UserRegisterResponse{
			Response: model.NewResponse(-1, "user register fails"),
		}
	}
	f.userId = newUser.Id

	// TODO 生成token

	return &model.UserRegisterResponse{
		Response: model.NewResponse(0, "user register success"),
		UserId:   f.userId,
		Token:    "",
	}
}
