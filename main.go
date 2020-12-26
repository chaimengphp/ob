package main

import (
	"github.com/astaxie/beego/client/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "obapi/routers"

	beego "github.com/astaxie/beego/server/web"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default","mysql","root:123456@/ob?charset=utf8")
	orm.Debug = true
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
