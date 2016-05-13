package controllers

import (
  "github.com/astaxie/beego"
  _ "github.com/astaxie/beego/cache/redis"
)


type SignoutController struct {
  beego.Controller
}


func (this *SignoutController) Get() {
  beego.Debug("In SignoutController:Get - Start")

  this.DelSession("session")

  beego.Debug("In SignoutController:Get - Deleted user session redirecting back to /accounts/signin")

  this.Redirect("/accounts/signin", 303)
}
