package controllers

import "github.com/astaxie/beego"

// CatalogController контроллер основых действий
type CatalogController struct {
	beego.Controller
}

// Get позволяет общаться сайту по протоколу Get
func (c *CatalogController) Get() {
	var q КоммерческаяИнформация
	q = LoadFile()
	c.Data["Hi"] = "catalog.go"
	c.Data["Tovar"] = q.Каталог.Товары.Товар
	c.TplName = "catalog.tpl"
}
