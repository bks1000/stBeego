package routers

import (
	"mygoweb/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.DefaultController{})
	beego.Router("/home", &controllers.HomeController{})
	beego.Router("/api", &controllers.ApiController{})

}
