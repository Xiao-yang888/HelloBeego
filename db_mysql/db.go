package db_mysql

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
)

/**
 *连接mysql数据库
 */

func Connect() {

	//项目配置
	config := beego.AppConfig
	dbDriverName := config.String("db_driverName")//数据库驱动
	dbUser := config.String("db_user")
	dbPassed := config.String("db_password")
	dbIp := config.String("db_ip")
	dbName := config.String("db_name")
	fmt.Println(dbDriverName,dbUser,dbPassed)

	//连接数据库
	connUrl := dbUser + ":" +  dbPassed+"@tcp("+dbIp+")/"+dbName+"?charset+=utf8"
	db,err := sql.Open(dbDriverName,connUrl)
	if err != nil {//err ！= nil 表示连接数据库时出现错误
		panic("数据连接打开失败，请重试")
	}
	DB = db
	fmt.Println(db)
}

/**
 *将用户信息保存到
 */
func AddUser(u model.User)(int64, error){
	//1,将密码进行hash计算，得到密码hash值，然后再存
	md5Hash := md5.New()
	md5Hash.Write([]byte(u.Password)
	passwordBytes := md5Hash.Sum(nil)
	u.Password

	result, err := Db.Exec(insert into user(name,birthday,password))
	   "value(?,?,?,?)",u.Name,u.Birthday,u.Password)
    if err != nil{
    	return -1,err
	}
}

func main() {
	db_masql.Connect()




	//2,其他配置

	//3，启动应用
	beego.Run()
}


