package routers

import (
	"github.com/zhangyuyu/zy-go-todo/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/login",&controllers.UserLoginController{}, "get:Get;post:Login")
	beego.Router("/logout",&controllers.UserLogoutController{})
	beego.Router("/tasks/", &controllers.TaskController{}, "get:ListTasks;post:NewTask")
	beego.Router("/tasks/:id:int", &controllers.TaskController{}, "get:GetTask;put:UpdateTask;delete:DeleteTask")

}
