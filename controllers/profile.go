package controllers

import (
  "golang.org/x/crypto/bcrypt"
  "github.com/astaxie/beego"
  "github.com/astaxie/beego/orm"
  "github.com/astaxie/beego/validation"
  _ "github.com/astaxie/beego/cache/redis"
  "github.com/TalLannder/beego-ureg/models"
)


type ProfileController struct {
  AccountController
}


func (this *ProfileController) Get() {
  beego.Debug("In ProfileController:Get - Start")

  o := orm.NewOrm()
  o.Using("default")

  user := models.Account{Uid: this.session["uid"].(string)}

  beego.Debug("In ProfileController:Get - Reading user from the database")

  err := o.Read(&user)

  if err != nil {
    flash := beego.NewFlash()
    flash.Error("Internal server error - Please try later or let us know that something whent wrong.")
    flash.Store(&this.Controller)
    this.DelSession("session")
    this.Redirect("/accounts/signin", 303)
  }

  this.Data["User"] = user
  this.Data["ProfileActive"] = true
  this.LayoutSections["Footer"] = "profilecontroller/footer.tpl"

}//end ProfileController:Get() func


func (this *ProfileController) Post() {
  beego.Debug("In ProfileController:Post - Start")

  this.Data["ProfileActive"] = true
  this.LayoutSections["Footer"] = "profilecontroller/footer.tpl"

  flash := beego.NewFlash()

  o := orm.NewOrm()
  o.Using("write")

  user := models.Account{Uid: this.session["uid"].(string)}

  beego.Debug("In ProfileController:Post - Reading user from the database")

  err := o.Read(&user)

  if err != nil {
    flash.Error("Internal server error - Please try later or let us know that something whent wrong.")
    flash.Store(&this.Controller)
    this.DelSession("session")
    this.Redirect("/accounts/signin", 303)
  }

  // In this case try to parse the submitted form if unable to
  // parse use the exsisting user information to generate the page.
  this.Data["User"] = user

  accountUpdateForm := models.FormAccountUpdate{}

  if err := this.ParseForm(&accountUpdateForm); err != nil {
    beego.Debug("In ProfileController:Post - Got err parsing the form", err)
    flash.Error("Internal server error - Please try later or let us know that something whent wrong.")
    flash.Store(&this.Controller)
    return
  }

  // After parsing save the form data to the controller to preserve
  // the information. This will help the user in case of validation failure.
  user = user.CopyUpdateForm(&accountUpdateForm)
  this.Data["User"] = user

  valid := validation.Validation{}

  if v, _ := valid.Valid(&accountUpdateForm); !v {
    beego.Debug("In ProfileController:Post - Got form validation err")
    this.Data["Errors"] = valid.ErrorsMap
    return
  }

  //******** Compare submitted password with the saved hash
  err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(accountUpdateForm.CurrentPassword))

  if err != nil {
    errormap := make(map[string]string)
    errormap["Current"] = "The current password seems to be incorrect, please try again."
    this.Data["Errors"] = errormap
    flash.Error("The current password seems to be incorrect, please try again.")
    flash.Store(&this.Controller)
    return
  }

  //******** Save user info to database
  user = user.CopyUpdateForm(&accountUpdateForm)

  _, err = o.Update(&user)

  if err != nil {

    beego.Error("In ProfileController:Post - Gor err updating user in database", err)
    flash.Error("Internal server error - Please try later or let us know that something whent wrong.")
    flash.Store(&this.Controller)
    this.DelSession("session")
    this.Redirect("/accounts/signin", 303)
  }

  flash.Notice("Profile updated")
  flash.Store(&this.Controller)

  //******** update session
  m := make(map[string]interface{})

  m["uid"] = user.Uid
  m["firstname"] = accountUpdateForm.First
  m["lastname"] = accountUpdateForm.Last
  m["username"] = accountUpdateForm.Email

  beego.Debug("In SigninController:Post - Creating new session")
  this.SetSession("session", m)

  this.session["username"] = accountUpdateForm.Email
  this.session["first"] = accountUpdateForm.First

  this.Data["User"] = user
  this.Redirect("/accounts/profile", 303)

}//end ProfileController:Post func
