package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/d16izhevsk/1coutlet/models"
)

// CatalogController контроллер основых действий
type CatalogController struct {
	beego.Controller
	CI models.КоммерческаяИнформация
}

// Get позволяет общаться сайту по протоколу Get
func (c *CatalogController) Get() {
	c.Data["SecretCode"] = models.SecretCode
	models.SecretCode++
	// c.Data["Tovar"] = models.CI.Каталог.Товары.Товар

	o := orm.NewOrm()

	tovar := new(models.Tovar)
	qs := o.QueryTable(tovar)

	var tovars []*models.Tovar
	c.Data["Count"], c.Data["Err"] = qs.Limit(300).All(&tovars)
	c.Data["Tovars"] = tovars
	c.TplName = "catalog.tpl"
}
