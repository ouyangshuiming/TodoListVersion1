package service

import (
	"toDoListDemo/model"
	"toDoListDemo/serializer"
	"toDoListDemo/utils"
)

type UserService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required, min=3,max=15"`
	Password string `form:"password" json:"password" binding:"required, min=3,max=16"`
}

func (service *UserService) Register() serializer.Response {
	var user model.User

	var count int
	//先查找一下数据库中有没有这个人,如果已经有了肯定不能让他注册
	model.DB.Model(&model.User{}).Where("user_name=?", service.UserName).First(&user).Count(&count)

	if count == 1 { //已经有这个人，不用再注册了
		return serializer.Response{Status: 400, Msg: "已经有这个人了，无需再注册了"}
	}
	user.UserName = service.UserName

	//然后对密码进行加密
	user.SetPassword(service.Password)

	//创建用户
	model.DB.Create(&user)

	return serializer.Response{Status: 200, Msg: "用户注册成功"}
}

func (service *UserService) Login() serializer.Response {
	var user model.User

	//先查找一下数据库中有没有这个人,如果有才能登陆
	error := model.DB.Where("user_name=?", service.UserName).First(&user).Error

	if error != nil {
		return serializer.Response{Status: 200, Msg: "用户不存在请先登录"}
	}

	if user.CheckPassword(service.Password) == false { //密码填错了
		return serializer.Response{Status: 400, Msg: "密码错误"}
	}

	//密码输入正确，发送一个token给客户端
	token, _ := utils.GenerateToken(user.ID, service.UserName, service.Password)

	return serializer.Response{
		Status: 200,
		Data:   serializer.TokenData{User: serializer.BuildUser(user), Token: token},
		Msg:    "登录成功",
	}
}
