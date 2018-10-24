package controllers

import (
	//"encoding/json"
	"fmt"

	"mygoweb/models"

	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "login.tpl"
}

func (c *HomeController) Post() {

	//前端传递json格式数据({"LogonName":"changxb","PassWord":"123456"})，用下面一组解析
	//data := make(map[string]string)
	//json.Unmarshal(c.Ctx.Input.RequestBody, &data)
	//body := c.Ctx.Input.RequestBody
	//beego.Debug(fmt.Println(body))
	//username := data["LogonName"]
	//beego.Debug(fmt.Println(username))
	//password := data["PassWord"]
	//beego.Debug(fmt.Println(password))

	//前端传递Form Data(LogonName=changxb&PassWord=123456) 用下面一组解析
	username := c.GetString("LogonName")
	password := c.GetString("PassWord")
	beego.Debug(fmt.Println(username))
	beego.Debug(fmt.Println(password))

	users, _ := models.UserQuery("张三")
	beego.Debug(users)
	c.Ctx.SetCookie("token", username)
	//c.Ctx.ResponseWriter.Write([]byte(username))
	c.Data["username"] = username
	//c.Layout = "layout.tpl" //设置index.tpl使用layout.tpl作为模板
	//c.TplName = "index.tpl" //如果使用了模板，需要删除部分重复的html内容
	c.Layout = "layout.tpl"
	c.TplName = "main.tpl"
}
