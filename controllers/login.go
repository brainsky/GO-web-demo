package controllers

import (
	"fmt"
	"hello/models"

	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) LoginSignup() {
	fmt.Println("LoginSignup方法")
	this.TplName = "login.tpl"
	if flash != nil {
		flash.Store(&this.Controller)
	}

}
func (this *LoginController) LoginPost() {
	fmt.Println("Login Post方法")
	inputEmail := this.GetString("inputEmail")
	inputPassword := this.GetString("inputPassword")
	//连接数据库，获取数据库的数据
	fmt.Println(inputEmail, inputPassword)
	mysql := new(models.CruderImp)
	mysql.Con = models.Connect()
	customer := mysql.QueryNamePass(inputEmail, inputPassword)
	fmt.Println(customer)
	if customer.Name != "" {
		this.Data["name"] = customer.Name
		this.Redirect("/hello?name", 302)

	} else {
		flash = beego.NewFlash()
		flash.Error("邮箱或密码错误，请重新输入")
		this.Redirect("/hello/login", 302)
	}
}
