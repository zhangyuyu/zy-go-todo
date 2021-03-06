package controllers

import (
	"github.com/astaxie/beego"
	"github.com/zhangyuyu/zy-go-todo/models"
)

type UserLoginController struct {
	beego.Controller
}

func (this *UserLoginController) Get() {
	this.TplName = "login.html"
	this.Render()
}

func (this *UserLoginController) Login() {
	var user models.User
	inputs := this.Input()
	user.Username = inputs.Get("username")
	password := inputs.Get("password")
	err := models.ValidateUser(user, password)
	if err == nil {
		u, _ := models.FindUser(user.Username)
		this.SetSession("userLogin", u.Id)
		this.SetSession("username", u.Username)
		this.Redirect("/", 302)
	} else {
		this.TplName = "login.html"
		this.Data["Error"] = err
		this.Render()
	}
}

type UserLogoutController struct {
	beego.Controller
}

func (this *UserLogoutController) Logout() {
	this.DelSession("userLogin")
	this.DelSession("username")
	this.Data["json"] = Response{Status: 302, Location: "/login"}
	this.ServeJSON()
}

type Response struct {
	Status   int64
	Location string
}