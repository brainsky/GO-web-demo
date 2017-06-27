package routers

import (
	"hello/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/hello", &controllers.MainController{}, "get:MyHello")
	beego.Router("/hello/register", &controllers.RegisterController{}, "get:RegisterGet;post:Post")
	beego.Router("/hello/login", &controllers.LoginController{}, "get:LoginSignup;post:LoginPost")
}
