package controllers

import (
	"github.com/astaxie/beego"
)

// MainController главный контроллер веб-сервера нашего приложения
type MainController struct {
	beego.Controller
}

// Get контроллера MainController
func (c *MainController) Get() {
	c.Data["Website"] = "liqw.ru"
	c.Data["Email"] = "dmylnikov@gmail.com"
	c.TplName = "index.tpl"
}
