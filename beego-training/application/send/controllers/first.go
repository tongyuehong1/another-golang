package controllers

import (
	"beego-training/application/send/models"
	"github.com/astaxie/beego"
)

type FirstController struct {
	beego.Controller
}

func (this *FirstController) Get() {
	name := this.GetString(":first")
	first := models.First{Name: name}
	this.Data["First"] = first
}
