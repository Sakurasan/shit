package controllers

import (
	"github.com/astaxie/beego"
	"net/http"
	"oneisall/utils"
)

type MainController struct {
	beego.Controller
	http.Request
}

type One struct {
	Id   int64
	Vol  int64
	Time string
	Img  string
	Word string

	Atitle  string
	Aurl    string
	Note    string
	Author  string
	Article string

	Qtitle string
	Qurl   string
	Qdesc  string
	Answer string
}

func (this *MainController) Get() {
	this.TplName = "index.html"

	a := utils.One2three()
	this.Data["Website"] = "OneIsAll"
	this.Data["Email"] = "astaxie@gmail.com"
	this.Data["Date"] = a

}
