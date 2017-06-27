package main

import (
	"fmt"
	"hello/models"
	_ "hello/routers"

	"github.com/astaxie/beego"
)

func main() {
	models.Init()
	beego.Run()
}
