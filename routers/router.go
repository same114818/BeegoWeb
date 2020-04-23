package routers

import (
	"BeegoWeb/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Include(&controllers.IndexController{}, &controllers.ContactController{})
}
