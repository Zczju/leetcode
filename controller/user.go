package controller

import (
	G "github.com/NoCLin/douyin-backend-go/config/global"
	"github.com/NoCLin/douyin-backend-go/model"
	"github.com/NoCLin/douyin-backend-go/service"
	"github.com/NoCLin/douyin-backend-go/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var (
	usersLoginInfo = map[string]model.UserInfo{}
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
//
//var userIdSequence = int64(1)

func Register(c *gin.Context) {
	username := c.Query("name")
	if !utils.UserNameValid(username) {
		c.JSON(200, &model.UserRegisterResponse{
			Response: model.NewResponse(-1, "username invalid"),
		})
	}
	password := c.Query("password")
	if !utils.PasswordValid(password) {
		c.JSON(200, &model.UserRegisterResponse{
			Response: model.NewResponse(-1, "password invalid"),
		})
	}
	data := service.NewUserRegisterFlow(username, password).Register()
	c.JSON(200, data)
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	// TODO token 校验
	if !utils.UserNameValid(username) {
		c.JSON(200, &model.UserLoginResponse{
			Response: model.NewResponse(-1, "username invalid"),
		})
	}
	if !utils.PasswordValid(password) {
		c.JSON(200, &model.UserLoginResponse{
			Response: model.NewResponse(-1, "password invalid"),
		})
	}

	data := service.NewUserLoginInstance(username, password).Login()
	c.JSON(200, data)
}

func GetUserInfo(c *gin.Context) {
	userIdStr := c.Query("user_id")
	userId, err := strconv.ParseInt(userIdStr, 0, 64)
	if err != nil {
		c.JSON(200, &model.UserResponse{
			Response: model.NewResponse(-1, "fail to parse user id"),
		})
	}
	if !utils.UserIdValid(userId) {
		c.JSON(200, &model.UserResponse{
			Response: model.NewResponse(-1, "User Id invalid"),
		})
	}
	//TODO 获取token并对其进行解析
	//token := c.Query("token")
	//parseToken(token)

	data := service.NewGetUserInfoFlowInstance(userId).GetUserInfo()
	c.JSON(200, data)
}

func Test(c *gin.Context) {

	user := model.User{Name: "Jinzhu"}

	result := G.DB.Create(&user)
	print(result)
	c.JSON(http.StatusOK, model.UserResponse{
		Response: model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
	})
}
