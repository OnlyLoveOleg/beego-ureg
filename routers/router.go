package routers

import (
	"github.com/astaxie/beego"
	"github.com/TalLannder/beego-ureg/controllers"
)


func init() {

  beego.Router("/", &controllers.IndexController{})

	beego.Router("/accounts/signup", &controllers.SignupController{})

	beego.Router("/accounts/signin", &controllers.SigninController{})

	beego.Router("/accounts/signout", &controllers.SignoutController{})

	beego.Router("/accounts/verify/:uuid([0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12})", &controllers.AccountVerifyController{})

  beego.Router("/accounts/profile", &controllers.ProfileController{})

	beego.Router("/accounts/delete", &controllers.DeleteController{})

  beego.Router("/accounts/security", &controllers.SecurityController{})

  beego.Router("/accounts/forgotpassword", &controllers.ForgotPasswordController{})	

  beego.Router("/accounts/resetpassword/:uuid([0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12})", &controllers.ResetPasswordController{})
}
