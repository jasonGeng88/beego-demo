package routers

import (
	"github.com/jasonGeng88/beego-demo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
