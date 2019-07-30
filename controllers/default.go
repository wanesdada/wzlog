package controllers

import (
	"github.com/astaxie/beego"
	"wzlog/wanzlog"
)


var apiLog *log.Log


type MainController struct {

	beego.Controller

}

func (c *MainController) Get() {
	apiLog = log.Init("20060102.error")
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"

	apiLog.Println("URI:", c.Ctx.Input.URI())
}
