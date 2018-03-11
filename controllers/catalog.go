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

// Catalog содержит базу данных
type Catalog struct {
	Id      int
	Idc     string
	Articul string
	Tname   string
	Gruppa  string
	Gname   string
}

// Get позволяет общаться сайту по протоколу Get
func (c *CatalogController) Get() {

	grpid, _ := c.GetInt("grpid")
	o := orm.NewOrm()
	c.Data["SecretCode"] = models.SecretCode
	models.SecretCode++

	// Выбираем подгруппы начиная с корня
	qbg, _ := orm.NewQueryBuilder("mysql")
	qbg.Select("gruppa.id", "gruppa.idc", "gruppa.name", "gruppa.pid").
		From("gruppa").Where("gruppa.pid=?").Limit(100)
	sql2 := qbg.String()

	var grp []*models.Gruppa
	c.Data["Countg"], _ = o.Raw(sql2, grpid).QueryRows(&grp)
	c.Data["Grupps"] = grp

	// Показываем товары в выбранных подгруппах
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("tovar.id", "tovar.idc", "tovar.articul", "tovar.name as tname", "tovar.gruppa", "gruppa.name as gname").
		From("tovar").LeftJoin("gruppa").On("tovar.gruppa=gruppa.idc").Where("tovar.Gruppa IN (SELECT idc FROM gruppa WHERE pid=? OR id=?)").Limit(500)
	sql := qb.String()
	// beego.Info(sql)
	var cat []*Catalog
	c.Data["Count"], _ = o.Raw(sql, grpid, grpid).QueryRows(&cat)
	c.Data["Tovars"] = cat

	// Получить группу и группы вышестоящие
	tekGruppa := models.Gruppa{Id: grpid}
	err := o.Read(&tekGruppa)
	if err != nil {
		c.Data["Gruppa"] = "1С витрина"
	} else {
		c.Data["Gruppa"] = tekGruppa.Name
	}
	// beego.Info(err)

	c.TplName = "catalog.tpl"
}
