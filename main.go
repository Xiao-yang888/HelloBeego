package main

import "github.com/astaxie/beego"

func main() {
/**
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

	driverName := config.String("db_driverName")//数据库驱动
	dbUser := config.String("db_user")
	dbPassed := config.String("db_password")
	dbIp := config.String("db_ip")
	dbName := config.String("db_name")

    //连接数据库
    connUrl := dbUser + ":" +  dbPassed+"@tcp("+dbIp+")/"+dbName+"?charset+=utf8"
	db,err := sql.Open(driverName,connUrl)
	if err != nil {//err ！= nil 表示连接数据库时出现错误
		panic("数据连接打开失败，请重试")
	}
	fmt.Println

 */

	//其他配置

	beego.Run()//代码简洁
}

//代码封装：可以将重复的代码或者功能相对比较独立的代码 ，进行封装，
//以函数的形式进行封装变成一个代码块或者是功能包，供使用者进行调用