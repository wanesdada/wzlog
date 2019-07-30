package main

import (
	"github.com/astaxie/beego"
	"wzlog/merror"
	_ "wzlog/routers"
)

func main() {
	merror.Log("bee run")
	beego.Run()
}

