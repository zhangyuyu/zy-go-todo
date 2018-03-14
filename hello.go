package main

import (
	"github.com/astaxie/beego"
	_ "github.com/zhangyuyu/hello/routers"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}
	beego.Run(":" + port)
}
