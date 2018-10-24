package controllers

import (
	"fmt"

	"github.com/astaxie/beego"

	"encoding/json"
)

type ApiController struct {
	beego.Controller
}

//首字母大写是公有的，首字母小写是私有的
type People struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (c *ApiController) Get() {
	mystruct := new(People)
	mystruct.Age = 1
	mystruct.Name = "zs"

	mystruct2 := People{"ls", 18}
	fmt.Println(typeof(mystruct2))
	//c.Data["json"] = &mystruct2
	beego.Debug(*mystruct)
	beego.Debug(json.Marshal(*mystruct))
	beego.Debug(mystruct)
	beego.Debug(&mystruct)
	c.Data["json"] = map[string]interface{}{"code": 0, "message": "获取成功", "date": *mystruct}

	c.ServeJSON()
}

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}
