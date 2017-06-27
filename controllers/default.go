package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (main *MainController) MyHello() {
	main.Data["Website"] = "Hello Web"
	main.Data["Email"] = "livetoend@163.com"
	main.TplName = "hello.tpl"
	name := main.GetString("name")
	if name != "" {
		main.Data["Name"] = name
	}
}
