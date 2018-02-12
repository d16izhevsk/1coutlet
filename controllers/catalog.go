package controllers

import "github.com/astaxie/beego"

// CatalogController контроллер основых действий
type CatalogController struct {
	beego.Controller
}

// Get позволяет общаться сайту по протоколу Get
func (c *CatalogController) Get() {
	c.TplName = "catalog.tpl"
}
