package main

import (
	"github.com/astaxie/beego"
	_ "github.com/zituocn/UserAPI/docs"
	_ "github.com/zituocn/UserAPI/routers"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
