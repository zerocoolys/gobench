package main

import (
	"benchmark/controllers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/benchmark",&controllers.BenchmarkController{})
	beego.Router("/run",&controllers.RunController{})
	beego.Run()
}

