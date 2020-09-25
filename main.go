package main

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	_ "hellobeego/routers"
)

func main() {
	//定义congig变量，接收并赋值为全局配置变量
	config := beego.AppConfig
	//获取配置选项
	appname := config.String("appname")
	fmt.Println("项目应用名称",appname)
	port,err := config.Int("httpport")
	if err != nil{
		//配置信息解析错误
		panic("项目配置信息解析错误，请查验后重试")
	}
	fmt.Println("应用监听窗口",port)

	driver := config.String("db_driver")//数据库驱动
	dbUser := config.String("db_user")
	dbPassed := config.String("db_password")
	dbIp := config.String("db_ip")
	dbName := config.String("db_name")

	db,err := sql.Open(driver,dbUser+":"+dbPassed+"@tcp("+dbIp+")/"+dbName+"?charset=utf8")
	if err != nil {
		panic("数据连接打开失败，请重试")
	}
	fmt.Println(db)

	beego.Run()
}

