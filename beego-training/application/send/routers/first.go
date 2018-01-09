package routers

import (
	"github.com/astaxie/beego"
	"beego-training/application/send/controllers"
	)

func init() {
	beego.Router("/", &controllers.FirstController{})
	beego.Router("/first/:first", &controllers.FirstController{})
}
