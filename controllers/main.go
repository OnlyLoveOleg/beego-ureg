package controllers

import (
  "github.com/astaxie/beego"
  "github.com/astaxie/beego/cache"
  _ "github.com/astaxie/beego/cache/redis"
  "github.com/astaxie/beego/utils/captcha"
  "github.com/twinj/uuid"
)


var(
  app_name_space, _ = uuid.Parse(beego.AppConfig.String("app_name_space"))
	base_url string = beego.AppConfig.String("base_url")
	gmail_account string = beego.AppConfig.String("gmail_account")
	gmail_account_password string = beego.AppConfig.String("gmail_account_password")
	mail_host string = beego.AppConfig.String("mail_host")
	mail_host_port, err = beego.AppConfig.Int("mail_host_port")
	CacheProvider string = beego.AppConfig.String("CacheProvider")
	CacheConnection string = beego.AppConfig.String("CacheConnection")
	cpt *captcha.Captcha
  email_config string = `{"username":"` + gmail_account + `","password":"` + gmail_account_password + `","host":"smtp.gmail.com","port":587}`
)


func init() {
  store, _ := cache.NewCache(CacheProvider, CacheConnection)
  cpt = captcha.NewWithFilter("/accounts/captca/", store)
}


type BaseController struct {
  beego.Controller
  logged bool
}


func (this *BaseController) Prepare() {
  beego.Debug("In BaseController:Prepare - Start")

  session := this.GetSession("session")
  
  if session != nil {
    beego.Debug("In BaseController:Prepare - Session found redirecting user to the /account/profile page")
    this.Redirect("/accounts/profile", 303)
    return
  }

  beego.ReadFromRequest(&this.Controller)

  this.Data["Title"] = "Beego-Ureg"
  this.Data["base_url"] = base_url

  this.Layout = "shared/layout.tpl"

  this.LayoutSections = make(map[string]string)
  this.LayoutSections["Header"] = "shared/header.tpl"
  this.LayoutSections["Footer"] = "shared/footer.tpl"
}


type AccountController struct {
  beego.Controller
  logged bool
  session map[string]interface{}
}

func (this *AccountController) Prepare() {
  beego.Debug("In AccountController:Prepare - Start")

  session := this.GetSession("session")

  if session == nil {
    beego.Debug("In AccountController:Prepare - Session not found kick the user back to /account/signin")
    this.Redirect("/accounts/signin", 303)
    return
  }

  beego.Debug("In AccountController:Prepare - found session setting logged flag to true")
  this.Data["logged"] = true
  this.Data["Session"] = session.(map[string]interface{})
  this.logged = true
  this.session = session.(map[string]interface{})

  beego.ReadFromRequest(&this.Controller)

  this.Data["Title"] = "Beego-Ureg"
  this.Data["base_url"] = base_url

  this.Layout = "shared/layout.tpl"

  this.LayoutSections = make(map[string]string)
  this.LayoutSections["Header"] = "shared/header.tpl"
  this.LayoutSections["Footer"] = "shared/footer.tpl"
}
