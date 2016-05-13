package controllers

import (
  "golang.org/x/crypto/bcrypt"
  "github.com/astaxie/beego"
  "github.com/astaxie/beego/orm"
  "github.com/astaxie/beego/validation"
  _ "github.com/astaxie/beego/cache/redis"
  "github.com/TalLannder/beego-ureg/models"
)


type SecurityController struct {
  AccountController
}


func (this *SecurityController) Get() {
  beego.Debug("In SecurityController:Get - Start")
  this.Data["SecurityActive"] = true
  this.LayoutSections["Footer"] = "securitycontroller/footer.tpl"
}


func (this *SecurityController) Post() {

  beego.Debug("In SecurityController:Post - Start")

  this.LayoutSections["Footer"] = "securitycontroller/footer.tpl"

  o := orm.NewOrm()
  o.Using("write")

  user := models.Account{Uid: this.session["uid"].(string)}

  beego.Debug("In SecurityController:Post - Reading user from the database")

  err := o.Read(&user)

  if err != nil {
    flash := beego.NewFlash()
    flash.Error("Internal server error - Please try later or let us know that something whent wrong.")
    flash.Store(&this.Controller)
    this.DelSession("session")
    this.Redirect("/accounts/signin", 303)
  }

  // In this case try to parse the submitted form if unable to
  // parse use the exsisting user information to generate the page.
  this.Data["User"] = user
  this.Data["SecurityActive"] = true

  flash := beego.NewFlash()

  passwordUpdateForm := models.FormPasswordUpdate{}

  if err := this.ParseForm(&passwordUpdateForm); err != nil {
    beego.Debug("In SecurityController:Post - Got err parsing the form", err)
    flash.Error("Internal server error - Please try later or let us know that something whent wrong.")
    flash.Store(&this.Controller)
    return
  }

  // After parsing save the form data to the controller to preserve
  // the information. This will help the user in case of validation failure.
  //user = user.CopyUpdateForm(&passwordUpdateForm)
  //this.Data["User"] = user
  

  valid := validation.Validation{}

  if v, _ := valid.Valid(&passwordUpdateForm); !v {
    beego.Debug("In SecurityController:Post - Got form validation err")
    this.Data["Errors"] = valid.ErrorsMap
    return
  }

  //******** Compare submitted password with the saved hash
  err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passwordUpdateForm.CurrentPassword))

  if err != nil {
    flash.Error("The email or password you entered is incorrect.")
    flash.Store(&this.Controller)
    return
  }

  bcryptPasswordHash, err := CryptPassword([]byte(passwordUpdateForm.NewPassword))

  if err != nil {
    beego.Debug("In SecurityController:Post - Got err hashing the password: ", err)
    flash.Error("Internal server error - Please try later or let us know that something whent wrong.")
    flash.Store(&this.Controller)
    return
  }

  user.Password = string(bcryptPasswordHash)

  //******** Save account to database
  _, err = o.Update(&user)

  if err != nil {

    beego.Error("In SecurityController:Post - Gor err updating user in database", err)
    flash.Error("Internal server error - Please try later or let us know that something whent wrong.")
    flash.Store(&this.Controller)
    return
  }

  flash.Notice("Password updated cheers!")
  flash.Store(&this.Controller)

  this.Redirect("/accounts/security", 303)

}//end SecurityController:Post() func
