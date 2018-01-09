package routers

import (
	"beego-training/application/student/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.StudentController{},"get:Show")
}
