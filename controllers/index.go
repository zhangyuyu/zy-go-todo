package controllers

import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
	sess := this.StartSession()
	username := sess.Get("username")
	if username == nil || username == "" {
		this.Redirect("/login", 302)
	} else {
		this.TplName = "index.html"
		this.Data["Username"] = username
		this.Render()
	}
}
