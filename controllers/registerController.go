package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"hellobeego/db_mysql"
	"hellobeego/models"
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
		////result := models.Result{
		//	Code:    3,
		//	Message: "数据接受错误，请重试",
		//	Data:    nil,
		//}
		//r.Data["json"] = &result
		//r.ServeJSON()
	}
	var user models.User
	err = json.Unmarshal(dataBytes,&user)
	if err != nil{
		r.Ctx.WriteString("数据解析错误，请重试")
		//result := models.Result{
		//	Code:    2,
		//	Message: "数据解析错误，请重试",
		//	Data:    nil,
		//}
		//r.Data["json"] = &result
		//r.ServeJSON()
		return
	}
	//一切正常，将用户信息保存到数据库中去
	//直接调用保存数据的一个函数，并判断保存后的结果
	//_, err = db_mysql.AddUser(user)
	//if err != nil{
	//	//r.Ctx.WriteString("注册用户信息失败，请重试")
	//	result := models.Result{
	//		Code:    0,
	//		Message: "注册用户信息信息失败，请重试",
	//		Data:    nil,
	//	}
	//	r.Data["json"] = &result
	//	r.ServeJSON()
	//}
	////fmt.Println(row)
	////md5Hash := md5.New()
	////md5Hash.Write([]byte(user.Password))
	//user.Password = hex.EncodeToString(md5Hash.Sum(nil))
	//result  := models.Result{
	//	Code: 1,
	//	Message: "恭喜，注册用户信息成功",
	//	Data:    user,
	//}
	////json.Marshal()编码
	//r.Data["json"] = &result
	//r.ServeJSON()//将result编码为json格式返回给前端
	////r.Ctx.WriteString("恭喜，注册用户信息成功")
	////向客户端返回字符串
	//
	////http.HandleFunc("/",)

	row, err :=db_mysql.AddUser(user)
	if err != nil {
		r.Ctx.WriteString("注册用户信息失败，请重试")
		return
	}
	fmt.Println(row)
	r.Ctx.WriteString("恭喜，注册用户信息成功")

}

//func Save(resp http.ResponseWriter,req *http.Request){
//	resp.Write([]byte)
//}