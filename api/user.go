package api

import (
	"github.com/gin-gonic/gin"
	"toDoListDemo/service"
)

func UserRegister(c *gin.Context) {
	var userRegister service.UserService
	c.ShouldBind(&userRegister)
	//从HTTP请求中获取数据(也就是用户名和密码)（其实是从表单中获取数据，也就是从用户名和密码两个框框中获取数据）
	//然后将数据绑定到对象userRegister上
	res := userRegister.Register() //调用注册方法
	c.JSON(200, res)

}

func UserLogin(c *gin.Context) {
	var userLogin service.UserService
	c.ShouldBind(&userLogin)
	//从HTTP请求中获取数据(也就是用户名和密码)（其实是从表单中获取数据，也就是从用户名和密码两个框框中获取数据）
	//然后将数据绑定到对象userRegister上
	res := userLogin.Login() //调用注册方法
	c.JSON(200, res)
}
