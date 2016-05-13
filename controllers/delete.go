package controllers

import (
  "github.com/astaxie/beego"
  "github.com/astaxie/beego/orm"
  "github.com/astaxie/beego/validation"
  _ "github.com/astaxie/beego/cache/redis"
  "golang.org/x/crypto/bcrypt"
  "github.com/TalLannder/beego-ureg/models"
)


type DeleteController struct {
  AccountController
}


func (this *DeleteController) Get() {
  beego.Debug("In DeleteController:Get - Start")

  this.Data["ProfileActive"] = true
}


func (this *DeleteController) Post() {
  beego.Debug("In DeleteController:Post - Start")

  this.Data["ProfileActive"] = true

  currentpassword := this.GetString("currentpassword")
  valid := validation.Validation{}
  valid.Required(currentpassword, "currentpassword")

  if valid.HasErrors() {
    errormap := make(map[string]string)
    for _, err := range valid.Errors {
      errormap[err.Key] = err.Message
    }
    this.Data["Errors"] = errormap
    return
  }

  flash := beego.NewFlash()

  //******** Read account from the database
  o := orm.NewOrm()
  o.Using("write")
  user := models.Account{Uid: this.session["uid"].(string)}

  err := o.Read(&user)

  if err != nil {
    beego.Error("In DeleteController:Post - Got err reading user from DB: ", err)
    flash.Error("Internal server error, Please try again later")
    flash.Store(&this.Controller)
    this.DelSession("session")
  this.Redirect("/accounts/signin", 303)
    return
  }

  //******** Compare submitted password with the saved hash
  err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(currentpassword))

  if err != nil {
    flash.Error("The current password seems to be incorrect, please try again.")
    flash.Store(&this.Controller)
    return
  }


 //******** Delete user record from the database
  _, err = o.Delete(&user)
  if err != nil {
    flash.Error("Internal server error")
    flash.Store(&this.Controller)
    this.DelSession("session")
    this.Redirect("/accounts/signin", 303)
    return
  }

  flash.Notice("Your account is deleted.")
  flash.Store(&this.Controller)
  this.DelSession("session")
  this.Redirect("/accounts/signin", 303)

}//end DeleteController:Post() func
