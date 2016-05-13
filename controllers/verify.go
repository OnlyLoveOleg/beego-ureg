package controllers

import (
  "github.com/astaxie/beego"
  "github.com/astaxie/beego/orm"
  _ "github.com/astaxie/beego/cache/redis"
  "github.com/TalLannder/beego-ureg/models"
)


type AccountVerifyController struct {
  BaseController
}


func (this *AccountVerifyController) Get() {

  beego.Debug("In AccountVerifyController:Get - Start")

  flash := beego.NewFlash()

	uuid := this.Ctx.Input.Param(":uuid")

	o := orm.NewOrm()
	o.Using("write")

	user := models.Account{Registration_uid: uuid}

	err := o.Read(&user, "Registration_uid")

  if err != nil {
    flash.Error("We've been unable to verify your account at the moment please try later or contact us for help.")
    flash.Store(&this.Controller)
    this.DelSession("session")
    this.Redirect("/accounts/signin", 303)
  }

  user.Registration_uid = ""

  if _, err := o.Update(&user); err != nil {

    beego.Error("In AccountVerifyController:Get - Got err updating user", err)
    flash.Error("We've been unable to verify your account at the moment please try later or contact us for help.")
    flash.Store(&this.Controller)
    this.DelSession("session")
    this.Redirect("/accounts/signin", 303)

  }

  flash.Notice("Your account successfully activated, feel free to sign in.")
  flash.Store(&this.Controller)
  this.Redirect("/accounts/signin", 303)
}
