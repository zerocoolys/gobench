package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "index.tpl"
}


type BenchmarkController struct {
	beego.Controller
}

func (this *BenchmarkController) Post(){
	this.Layout = "benchmark.html"
	this.TplNames = "benchmark.tpl"
	this.Data["Url"] = this.GetString("url")
	str,err := httplib.Post(this.Ctx.Request.Form.Get("url")).String()
	fmt.Println(str, err)
}
	

