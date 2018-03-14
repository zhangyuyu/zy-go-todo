package main

import (
	"github.com/astaxie/beego"
	_ "github.com/zhangyuyu/hello/routers"
)

func main() {
	beego.Run()
}
