package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"hellobeego/db_mysql"
	"io/ioutil"
)

type RegisterController struct {
	beego.Controller
}

func (r *RegisterController)Post()  {
	dataBytes, err := ioutil.ReadAll(r.Ctx.Request.Body)
	if err != nil {
		r.Ctx.WriteString("数据接收错误，请重试")
		return
	}
	var user models.User
	err = json.Unmarshal(dataBytes,&user)
	if err != nil{
		r.Ctx.WriteString("数据解析错误，请重试")
		return
	}
	//一切正常，将用户信息保存到数据库中去
	//直接调用保存数据的一个函数，并判断保存后的结果
	row, err := db_mysql.AddUser(user)
	if err != nil{
		r.Ctx.WriteString("注册用户信息失败，请重试")
		return
	}
	fmt.Println(row)
	r.Ctx.WriteString("恭喜，注册用户信息成功")
	return
}
