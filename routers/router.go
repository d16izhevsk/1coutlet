package routers

import (
	"github.com/astaxie/beego"
	"github.com/d16izhevsk/1coutlet/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/catalog", &controllers.CatalogController{})
	beego.Router("/exchange", &controllers.ExchangeController{})
	beego.Router("/tovar", &controllers.TovarController{})
	// json апи
	beego.Router("/v1/grups/", &controllers.CatalogController{}, "get:ListGrups")
}
