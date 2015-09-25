package main

import (
	_ "zyun-wechat-api/docs"
	_ "zyun-wechat-api/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDataBase("default", "mysql", "zyun_user:2wsx1qaz@tcp(127.0.0.1:3306)/zyun_platform")
}

func main() {
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

