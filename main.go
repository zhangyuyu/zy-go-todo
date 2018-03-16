package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/context"
	_ "github.com/zhangyuyu/zy-go-todo/routers"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	mysqlReg := getEnvValue("DB_USERNAME") + ":" +
		getEnvValue("DB_PASSWORD") + "@tcp(" +
		getEnvValue("DB_HOST") + ":3306)/" +
		getEnvValue("DB_DATABASE")

	orm.RegisterDataBase("default", "mysql", mysqlReg)
}

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.InsertFilter("*", beego.BeforeRouter, FilterUser)

	port := getEnvValue("PORT")
	beego.Run(":" + port)
}

func getEnvValue(key string) string {
	if os.Getenv(key) != "" {
		return os.Getenv(key)
	} else {
		return beego.AppConfig.String(key)
	}
}

var FilterUser = func(ctx *context.Context) {
	_, ok := ctx.Input.Session("userLogin").(int64)
	if !ok && ctx.Request.RequestURI != "/login" {
		ctx.Redirect(302, "/login")
	}
}
