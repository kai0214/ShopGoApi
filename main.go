package main

import (
	_ "ShopGoApi/routers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)
func init() {
	//设置数据库类型
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//连接数据库
	err := orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("mysql"))
	if err == nil {
		fmt.Println("数据库连接成功")
	} else {
		fmt.Println("数据库连接失败")
	}
}
func main() {

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
