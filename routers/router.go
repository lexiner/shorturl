package routers

import (
	"shorturl/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/create", &controllers.MainController{}, "post:CreateUrl")
	beego.Router("/:code", &controllers.MainController{}, "get:TargetUrl")
	beego.Router("/", &controllers.MainController{})
}
