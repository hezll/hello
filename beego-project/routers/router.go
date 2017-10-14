package routers

import (
	"github.com/hezll/hello/beego-project/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
