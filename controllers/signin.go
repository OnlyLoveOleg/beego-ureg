package controllers

import (
  "golang.org/x/crypto/bcrypt"
  "github.com/astaxie/beego"
  "github.com/astaxie/beego/orm"
  "github.com/astaxie/beego/validation"
  _ "github.com/astaxie/beego/cache/redis"
  "github.com/TalLannder/beego-ureg/models"
)


type SigninController struct {
  BaseController
}


func (this *SigninController) Get() {
  beego.Debug("In SigninController:Get - Start")
}


func (this *SigninController) Post() {
  beego.Debug("In SigninController:Post - Start")

  flash := beego.NewFlash()
  email := this.GetString("email")
  submittedPassword := this.GetString("password")
  redirect_after_login := this.GetString("redirect_after_login")

  valid := validation.Validation{}

  valid.Email(email, "email")
  valid.Required(submittedPassword, "password")

  this.Data["Email"] = email

  if valid.HasErrors() {
    errormap := make(map[string]string)
    for _, err := range valid.Errors {
      errormap[err.Key] = err.Message
    }
    this.Data["Errors"] = errormap
    return
  }

  //******** Read user account from the database
  o := orm.NewOrm()
  o.Using("default")

  user := models.Account{Email: email}

  err := o.Read(&user, "Email")
  
  if err == orm.ErrNoRows || err == orm.ErrMissPK {

    flash.Error("The email or password you entered is incorrect.")
    flash.Store(&this.Controller)
    return

  } else if err != nil {

    flash.Error("Internal server error - Sorry but we're unable to process your request at the moment. Please try later or contact support.")
    flash.Store(&this.Controller)
    return
  }

  if user.Registration_uid != "" {

    flash.Error("Please check your email to activate your account before login.")
    flash.Store(&this.Controller)
    return
  }

  //******** Compare submitted password with the saved hash
  err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(submittedPassword))

  if err != nil {
    flash.Error("The email or password you entered is incorrect.")
    flash.Store(&this.Controller)
    return
  }

  //******** Create session and go back to previous page
  m := make(map[string]interface{})

  m["uid"] = user.Uid
  m["firstname"] = user.First
  m["lastname"] = user.Last
  m["username"] = email

  beego.Debug("In SigninController:Post - Creating new session")
  this.SetSession("session", m)

  if redirect_after_login != "" {
    this.Redirect(redirect_after_login, 303)
  }

  this.Redirect("/accounts/profile", 303)

}//EOF SigninController:Post
