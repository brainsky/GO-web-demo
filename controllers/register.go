package controllers

import (
	"bytes"
	"fmt"
	"hello/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type RegisterController struct {
	beego.Controller
}

var flash *beego.FlashData

func (this *RegisterController) Post() {
	var buffer bytes.Buffer
	flash = beego.NewFlash()
	cus := new(models.Customer)
	//将页面的form转为struct中的字段
	if err := this.ParseForm(cus); err != nil {
		fmt.Println(err)
	}
	valid := validation.Validation{}
	b, err := valid.Valid(cus)
	fmt.Println("validation方法")
	fmt.Println(b)
	if err != nil {
		fmt.Println(err)
	}
	if !b {
		for _, err := range valid.Errors {
			//把错误转发到页面
			fmt.Println(err.Key, "iii", err.Message)
			switch {
			case err.Key == "Name.Required":
				flash.Data["NameErr"] = "用户名不能为空"
			case err.Key == "Password.Required":
				flash.Data["PasswordErr"] = "密码不能为空"
			case err.Key == "Email.Email":
				flash.Data["EmailErr"] = "请输入正确的邮箱"
			case err.Key == "Phone.Mobile":
				flash.Data["PhoneErr"] = "请输入正确的手机号"
			}
		}
		fmt.Println(buffer.String())
		flash.Error(buffer.String())
		this.Redirect("/hello/register", 302)
		return
	}

	//存到数据库
	mysql := new(models.CruderImp)
	mysql.Con = models.Connect()
	mysql.AddCus(cus)
	this.Redirect("login", 302) //转向登录页面
	return
}

func (this *RegisterController) RegisterGet() {
	this.TplName = "register.tpl"
	if flash != nil {
		flash.Store(&this.Controller)
	}
}
