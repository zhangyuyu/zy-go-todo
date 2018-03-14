package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/zhangyuyu/zy-go-todo/routers"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	mysqlReg := beego.AppConfig.String("mysqluser") + ":" +
		beego.AppConfig.String("mysqlpass") + "@tcp(127.0.0.1:3306)/" +
		beego.AppConfig.String("mysqldb")
	orm.RegisterDataBase("default", "mysql", mysqlReg)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}
	beego.Run(":" + port)
}