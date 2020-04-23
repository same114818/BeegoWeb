package main

import (
	"BeegoWeb/models"
	_ "BeegoWeb/routers"

	"github.com/astaxie/beego"
)

func main() {
	models.InitDatabase()
	beego.Run()
}
