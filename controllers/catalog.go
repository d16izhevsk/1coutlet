package controllers

import (
	"github.com/astaxie/beego"
	"github.com/d16izhevsk/1coutlet/models"
)

// CatalogController контроллер основых действий
type CatalogController struct {
	beego.Controller
	CI models.КоммерческаяИнформация
}

// Get позволяет общаться сайту по протоколу Get
func (c *CatalogController) Get() {
	c.Data["Hi"] = "Приветствуем вас в alpha версии 0.2"
	c.TplName = "catalog.tpl"
}
