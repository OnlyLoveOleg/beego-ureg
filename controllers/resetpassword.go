package controllers

import (
  "github.com/astaxie/beego"
  "github.com/astaxie/beego/orm"
  "github.com/astaxie/beego/validation"
  _ "github.com/astaxie/beego/cache/redis"
  "github.com/TalLannder/beego-ureg/models"
)


type ResetPasswordController struct {
  BaseController
}


func (this *ResetPasswordController) Get() {
  beego.Debug("In ResetPasswordController:Get - Start")
}


func (this *ResetPasswordController) Post() {

  beego.Debug("In ResetPasswordController:Post - Start")

  flash := beego.NewFlash()

  submitted_reset_uid := this.Ctx.Input.Param(":uuid")

  o := orm.NewOrm()
  o.Using("write")

  user := models.Account{Password_reset_uid: submitted_reset_uid}

  err := o.Read(&user, "Password_reset_uid")

  if err != nil {
    flash.Notice("Invalid key.")
    flash.Store(&this.Controller)
    this.Redirect("/accounts/signin", 303)
  }

  new_password := this.GetString("password")
  new_password_confirmation := this.GetString("password2")

  valid := validation.Validation{}
  valid.MinSize(new_password, 8, "new_password")
  valid.Required(new_password_confirmation, "new_password_confirmation")

  if valid.HasErrors() {
    errormap := make(map[string]string)
    for _, err := range valid.Errors {
      errormap[err.Key] = err.Message
    }
    this.Data["Errors"] = errormap
    return
  }

  if new_password != new_password_confirmation {
    flash.Error("Passwords don't match")
    flash.Store(&this.Controller)
    return
  }

  // Hashing the password with the default cost of 10
  bcryptPasswordHash, err := CryptPassword([]byte(new_password))

  if err != nil {
    beego.Error("In ResetPasswordController:Post - err hashing new_password: ", err)
    flash.Error("We're very sorry but we couldn't process your request.")
    flash.Store(&this.Controller)
    return
  }

  user.Password = string(bcryptPasswordHash)

  user.Password_reset_uid = ""

  if _, err := o.Update(&user); err != nil {
    flash.Error("Internal server error")
    flash.Store(&this.Controller)
    return
  }

  flash.Notice("Password updated.")
  flash.Store(&this.Controller)
  this.Redirect("/accounts/signin", 303)
}
