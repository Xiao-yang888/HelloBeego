package db_mysql

import (
	"BeegoProject0603/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	//mysql驱动
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB
/**
 *连接mysql数据库
 */

func Connect() {

	//项目配置
	config := beego.AppConfig
	dbDriver := config.String("db_driverName")//数据库驱动
	dbUser := config.String("db_user")
	dbPassed := config.String("db_password")
	dbIp := config.String("db_ip")
	dbName := config.String("db_name")
	fmt.Println(dbDriver,dbUser,dbPassed)

	//连接数据库
	connUrl := dbUser + ":" +  dbPassed+"@tcp("+dbIp+")/"+dbName+"?charset+=utf8"
	db,err := sql.Open(dbDriver,connUrl)
	if err != nil {// err不为nil，表示连接数据库时出现了错误, 程序就在此中断就可以，不用再执行了。
		//早解决，早解决
		panic("数据库连接错误，请检查配置")
	}
	Db = db
	fmt.Println(db)
}

/**
 *将用户信息保存到数据库中去的函数
 */
func AddUser(u models.User)(int64, error){
	//1,将密码进行hash计算，得到密码hash值，然后再存
	md5Hash := md5.New()
	md5Hash.Write([]byte(u.Password))
	passwordBytes := md5Hash.Sum(nil)
	u.Password = hex.EncodeToString(passwordBytes)

	result, err := Db.Exec("insert into user(name,birthday,password)" +
	   "value(?,?,?,?)",u.Name,u.Birthday,u.Address,u.Password)
    if err != nil{
    	return -1, err
	}
	row,err := result.RowsAffected()
	if err != nil{
		return -1, err
	}
	return row,nil
}


