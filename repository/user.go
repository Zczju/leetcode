package repository

import (
	"github.com/NoCLin/douyin-backend-go/config/global"
	"github.com/NoCLin/douyin-backend-go/model"
	"gorm.io/gorm"
	"sync"
)

type UserDao struct {
}

var (
	userDao *UserDao
	once    sync.Once
)

func NewUserDaoInstance() *UserDao {
	once.Do(func() {
		userDao = &UserDao{}
	})
	return userDao
}

func (*UserDao) InsertUser(user *model.User) (*model.User, error) {
	err := global.DB.Create(user).Error
	return user, err
}

func (*UserDao) QueryUserByName(username string) (*model.User, error) {
	var user model.User
	err := global.DB.Where("name = ?", username).Take(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (*UserDao) QueryUserById(id int64) (*model.User, error) {
	var user model.User
	err := global.DB.Where("id = ?", id).Take(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}
