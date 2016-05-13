package controllers

import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
    beego.Controller
}


// Here you can Go wild and customize the error pages for your server.
//
// Hint: you can also use separate templates instead of error_page.tpl :)

func (c *ErrorController) Error401() {
    c.Data["Content"] = "Unauthorized"
    c.TplName = "error_page.tpl"
}


func (c *ErrorController) Error403() {
    c.Data["Content"] = "Forbidden"
    c.TplName = "error_page.tpl"
}


func (c *ErrorController) Error404() {
    c.Data["Content"] = "My super cool 404 page"
    c.TplName = "error_page.tpl"
}


func (c *ErrorController) Error500() {
    c.Data["Content"] = "Oops something whent wrong"
    c.TplName = "error_page.tpl"
}


func (c *ErrorController) Error503() {
    c.Data["Content"] = "Service Unavailable"
    c.TplName = "error_page.tpl"
}
