package serializer

import "toDoListDemo/model"

type User struct {
	ID       uint   `json:"id" form:"id" example:"1"`                  //用户ID
	UserName string `json:"user_name"form:"user_name"example:"FanOne"` //用户名
	status   string `json:"status" form:"Status"`                      //用户状态
	CreateAt int64  `json:"create_at" form:"create_at"`                //创建
}

func BuildUser(user model.User) User {
	return User{
		ID:       user.ID,
		UserName: user.UserName,
		CreateAt: user.CreatedAt.Unix(),
	}
}
