package controllers

import (
  "github.com/astaxie/beego"
  "github.com/astaxie/beego/orm"
  "github.com/astaxie/beego/validation"
  "github.com/astaxie/beego/utils"
  _ "github.com/astaxie/beego/cache/redis"
  "github.com/twinj/uuid"
  "github.com/TalLannder/beego-ureg/models"
)


type ForgotPasswordController struct {
  BaseController
}


func (this *ForgotPasswordController) Get() {
  beego.Debug("In ForgotPasswordController:Get - Start")
}//end ForgotPasswordController:Get()


func (this *ForgotPasswordController) Post() {

  beego.Debug("In ForgotPasswordController:Post - Start")

  flash := beego.NewFlash()

  email := this.GetString("email")

  captchaVerification := cpt.VerifyReq(this.Ctx.Request)

  if !captchaVerification {
    errormap := make(map[string]string)
    beego.Debug("In ForgotPasswordController:Post - captchaVerification Got wrong captcha")
    errormap["Captcha"] = "Sorry but the characters you endered didn't match. Please try again"
    this.Data["Errors"] = errormap
    return
  }

  valid := validation.Validation{}
  valid.Email(email, "email")
  if valid.HasErrors() {
    errormap := make(map[string]string)
    for _, err := range valid.Errors {
      errormap[err.Key] = err.Message
    }
    this.Data["Errors"] = errormap
    return
  }
  beego.Debug("In ForgotPasswordController:Post - After Validation")

  o := orm.NewOrm()
  o.Using("default")
  user := models.Account{Email: email}
  err := o.Read(&user, "Email")
  if err != nil {
    flash.Error("Sorry but the email seems to be incorrect.")
    flash.Store(&this.Controller)
    return
  }

  reset_uid := uuid.NewV4().String()

  user.Password_reset_uid = reset_uid

  _, err = o.Update(&user)
  if err != nil {
    flash.Error("Internal error")
    flash.Store(&this.Controller)
    return
  }

  link := "http://" + base_url + "/accounts/resetpassword/" + reset_uid

  mail := utils.NewEMail(email_config)
  mail.To = []string{"user1@cupi.co"}
  mail.From = "spongebob@gmail.com"
  mail.Subject = "Beego-Ureg - Password Reset Request"
  mail.HTML = "To reset your password, please click on the link: <a href=\""+link+
  "\">"+link+"</a><br><br>Best Regards,<br>Awesome team"

  err = mail.Send()

  if err != nil {
    beego.Error("Unable to send password reset email")
    flash.Error("Unable to send password reset email")
    flash.Store(&this.Controller)
    this.Redirect("/accounts/signin", 303)
  }

  flash.Notice("We sent you a password reset link. Please check your email to reset the password.")
  flash.Store(&this.Controller)
  this.Redirect("/accounts/signin", 303)

} //end ForgotPassword() func
