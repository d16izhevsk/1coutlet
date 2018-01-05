package routers

import (
	"github.com/d16izhevsk/1coutlet/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
