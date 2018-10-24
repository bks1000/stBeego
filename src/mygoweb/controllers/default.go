package controllers

import (
	"github.com/astaxie/beego"
)

type DefaultController struct {
	beego.Controller
}

func (c *DefaultController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
