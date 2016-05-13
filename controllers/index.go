package controllers

import (
  "github.com/astaxie/beego"
)


type IndexController struct {
  BaseController
}


func (this *IndexController) Get() {
  beego.Debug("In IndexController.Get() - Start")
}
