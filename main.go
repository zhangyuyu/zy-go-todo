package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/auth"
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
	authPlugin := auth.NewBasicAuthenticator(SecretAuth, "My Realm")
	beego.InsertFilter("*", beego.BeforeRouter, authPlugin)

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

func SecretAuth(username, password string) bool {
	// The username and password parameters comes from the request header,
	// make a database lookup to make sure the username/password pair exist
	// and return true if they do, false if they dont.

	// To keep this example simple, lets just hardcode "hello" and "world" as username,password
	if username == "hello" && password == "world" {
		return true
	}
	return false
}
