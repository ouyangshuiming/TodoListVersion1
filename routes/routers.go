package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"toDoListDemo/api"
	"toDoListDemo/middleware"
)

func NewRouter() *gin.Engine {
	g := gin.Default()

	store := cookie.NewStore([]byte("something-very-secret")) //参数是加密session的密钥，返回的是一个实例对象
	g.Use(sessions.Sessions("mysession", store))              //使用g.Use()设置中间件，session的名字是”mysession“，
	//这个中间件会从http请求的cookie中获取名为"mysession"的数据，然后将数据存储在当前http请求的上下文Context中

	//对路由进行分组
	//跟用户有关的路由
	v1 := g.Group("api/v1")
	{
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)

		authed := v1.Group("/")
		authed.Use(middleware.JWT()) //jwt鉴权
		{
			authed.POST("task", api.CreateTask)
			authed.POST("task/:id", api.ShowTask) //展示一条备忘录的内容

		}
	}
	return g
}
