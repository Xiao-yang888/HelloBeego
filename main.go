package main

import (
	"fmt"
	_ "hellobeego/routers"
	"github.com/astaxie/beego"
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
	beego.Run()
}

