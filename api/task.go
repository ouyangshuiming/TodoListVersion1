package api

import (
	"github.com/gin-gonic/gin"
	"time"
	"toDoListDemo/model"
	"toDoListDemo/serializer"
	"toDoListDemo/utils"
)

type CreateTaskService struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Status  int    `json:"status" form:"status"` //0是未做，1是已做
}

// 创建一条备忘录
func CreateTask(c *gin.Context) {
	var createTaskService CreateTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	c.ShouldBind(&createTaskService) //数据绑定

	var user model.User
	model.DB.First(&user, claim.Id) //查到这个用户，赋值给user对象
	task := model.Task{
		User:      user,
		Uid:       user.ID,
		Title:     createTaskService.Title,
		Status:    0,
		Content:   createTaskService.Content,
		StartTime: time.Now().Unix(),
		EndTime:   0,
	}
	model.DB.Create(&task) //在task表中插入一行数据

	result := serializer.Response{
		Status: 200,
		Msg:    "创建成功",
	}
	c.JSON(200, result)
}

type ShowTaskService struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Status  int    `json:"status" form:"status"` //0是未做，1是已做
}

// 展示一条备忘录
// 传进来的是task的id
func ShowTask(c *gin.Context) {
	var task model.Task
	model.DB.First(&task, taskId)
	tasa
	return serializer.Response{
		Status: 200,
		Data:
	}

}

//
//import (
//	"github.com/gin-gonic/gin"
//	"net/http"
//	"toDoListDemo/service"
//)
//
//// CreateTaskHandler @Tags TASK
//// @Summary 创建任务
//// @Produce json
//// @Accept json
//// @Header 200 {string} Authorization "必备"
//// @Param data body service.CreateTaskService true  "title"
//// @Success 200 {object} serializer.ResponseTask "{"success":true,"data":{},"msg":"ok"}"
//// @Failure 500 {json} {"status":500,"data":{},"Msg":{},"Error":"error"}
//// @Router /task [post]
////func CreateTaskHandler() gin.HandlerFunc {
////	return func(ctx *gin.Context) {
////		var req service.CreateTaskService
////
////		if err := ctx.ShouldBind(&req); err == nil {
////			// 参数校验
////			l := service.GetTaskSrv()
////			resp, err := l.CreateTask(ctx.Request.Context(), &req)
////			if err != nil {
////				ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
////				return
////			}
////			ctx.JSON(http.StatusOK, resp)
////		} else {
////			util.LogrusObj.Infoln(err)
////			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
////		}
////
////	}
////}
//
//// ListTaskHandler @Tags TASK
//// @Summary 获取任务列表
//// @Produce json
//// @Accept json
//// @Header 200 {string} Authorization "必备"
//// @Param data body service.ListTasksService true "rush"
//// @Success 200 {object} serializer.ResponseTask "{"success":true,"data":{},"msg":"ok"}"
//// @Failure 500 {json} {"status":500,"data":{},"Msg":{},"Error":"error"}
//// @Router /tasks [get]
//func ListTaskHandler() gin.HandlerFunc {
//	return func(ctx *gin.Context) {
//		var req types.ListTasksReq
//		if err := ctx.ShouldBind(&req); err == nil {
//			// 参数校验
//			if req.Limit == 0 {
//				req.Limit = consts.BasePageLimit
//			}
//			l := service.GetTaskSrv()
//			resp, err := l.ListTask(ctx.Request.Context(), &req)
//			if err != nil {
//				ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
//				return
//			}
//			ctx.JSON(http.StatusOK, resp)
//		} else {
//			util.LogrusObj.Infoln(err)
//			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
//		}
//
//	}
//}
//
//// ShowTaskHandler @Tags TASK
//// @Summary 展示任务详细信息
//// @Produce json
//// @Accept json
//// @Header 200 {string} Authorization "必备"
//// @Param data body service.ShowTaskService true "rush"
//// @Success 200 {object} serializer.ResponseTask "{"success":true,"data":{},"msg":"ok"}"
//// @Failure 500 {json} {"status":500,"data":{},"Msg":{},"Error":"error"}
//// @Router /task/:id [get]
//func ShowTaskHandler() gin.HandlerFunc {
//	return func(ctx *gin.Context) {
//		var req types.ShowTaskReq
//		if err := ctx.ShouldBind(&req); err == nil {
//			// 参数校验
//			l := service.GetTaskSrv()
//			resp, err := l.ShowTask(ctx.Request.Context(), &req)
//			if err != nil {
//				ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
//				return
//			}
//			ctx.JSON(http.StatusOK, resp)
//		} else {
//			util.LogrusObj.Infoln(err)
//			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
//		}
//
//	}
//}
//
//// DeleteTaskHandler @Tags TASK
//// @Summary 删除任务
//// @Produce json
//// @Accept json
//// @Header 200 {string} Authorization "必备"
//// @Param data body service.DeleteTaskService true "用户信息"
//// @Success 200 {object} serializer.Response "{"success":true,"data":{},"msg":"ok"}"
//// @Failure 500 {json} {"status":500,"data":{},"Msg":{},"Error":"error"}
//// @Router /task/:id [delete]
//func DeleteTaskHandler() gin.HandlerFunc {
//	return func(ctx *gin.Context) {
//		var req types.DeleteTaskReq
//		if err := ctx.ShouldBind(&req); err == nil {
//			// 参数校验
//			l := service.GetTaskSrv()
//			resp, err := l.DeleteTask(ctx.Request.Context(), &req)
//			if err != nil {
//				ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
//				return
//			}
//			ctx.JSON(http.StatusOK, resp)
//		} else {
//			util.LogrusObj.Infoln(err)
//			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
//		}
//
//	}
//}
//
//// UpdateTaskHandler @Tags TASK
//// @Summary 修改任务
//// @Produce json
//// @Accept json
//// @Header 200 {string} Authorization "必备"
//// @Param	data	body	service.DeleteTaskService true "2"
//// @Success 200 {object} serializer.Response "{"success":true,"data":{},"msg":"ok"}"
//// @Failure 500 {json} {"status":500,"data":{},"Msg":{},"Error":"error"}
//// @Router /task [put]
//func UpdateTaskHandler() gin.HandlerFunc {
//	return func(ctx *gin.Context) {
//		req := new(types.UpdateTaskReq)
//		if err := ctx.ShouldBind(&req); err == nil {
//			// 参数校验
//			l := service.GetTaskSrv()
//			resp, err := l.UpdateTask(ctx.Request.Context(), req)
//			if err != nil {
//				ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
//				return
//			}
//			ctx.JSON(http.StatusOK, resp)
//		} else {
//			util.LogrusObj.Infoln(err)
//			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
//		}
//
//	}
//}
//
//// SearchTaskHandler @Tags TASK
//// @Summary 查询任务
//// @Produce json
//// @Accept json
//// @Header 200 {string} Authorization "必备"
//// @Param data body service.DeleteTaskService true "2"
//// @Success 200 {object} serializer.Response "{"success":true,"data":{},"msg":"ok"}"
//// @Failure 500 {json} {"status":500,"data":{},"Msg":{},"Error":"error"}
//// @Router /search [post]
//func SearchTaskHandler() gin.HandlerFunc {
//	return func(ctx *gin.Context) {
//		var req types.SearchTaskReq
//		if err := ctx.ShouldBind(&req); err == nil {
//			// 参数校验
//			l := service.GetTaskSrv()
//			resp, err := l.SearchTask(ctx.Request.Context(), &req)
//			if err != nil {
//				ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
//				return
//			}
//			ctx.JSON(http.StatusOK, resp)
//		} else {
//			util.LogrusObj.Infoln(err)
//			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
//		}
//
//	}
//}
