package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/d16izhevsk/1coutlet/models"
)

// TovarController главный контроллер веб-сервера нашего приложения
type TovarController struct {
	beego.Controller
}

// Get контроллера TovarController
func (c *TovarController) Get() {
	tovid, err := c.GetInt("tovid")
	if err == nil {
		c.Data["Tovid"] = tovid

		o := orm.NewOrm()
		tektov := models.Tovar{Id: tovid}
		err := o.Read(&tektov)
		if err == nil {
			c.Data["Tovar"] = tektov
		}

		c.TplName = "tovar.tpl"
	} else {

	}
}
