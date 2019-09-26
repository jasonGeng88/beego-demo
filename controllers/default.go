package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"2
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
