package main

import (
	"fmt"
	_ "mygoweb/routers"
	//"net/http"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func main() {

	//过滤中间件 开始

	//beego过滤器，通过cookie验证是否登录成功
	//https://blog.csdn.net/skh2015java/article/details/53944542

	var FilterUser = func(ctx *context.Context) {
		fmt.Println(ctx.Request.URL)
		if ctx.Request.URL.Path == "/home" {
			return
		}
		if ctx.GetCookie("token") == "" {
			ctx.Redirect(302, "/home")
		}
	}
	beego.InsertFilter("*", 1, FilterUser)
	//过滤中间件 结束

	//添加模板函数（在模板中处理数据） 开始
	var Hello = func(in string) (out string) {
		out = "hello " + in
		return
	}

	beego.AddFuncMap("hi", Hello)
	//添加模板函数		结束

	beego.SetStaticPath("/static", "static") //静态文件注册

	//beego默认情况下支持tpl和html后缀名的模板文件
	//beego.AddTemplateExt("html")

	beego.Run()
}
