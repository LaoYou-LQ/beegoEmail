package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "www.github.com"
	c.Data["Email"] = "hhsp.app"
	c.TplName = "index.tpl"
}
