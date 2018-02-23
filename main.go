package main

import (
	"github.com/astaxie/beego"
	_ "github.com/d16izhevsk/1coutlet/routers"
)

// CI предоставляет доступ к каталогу
var CI CommInfo

// main() запускает главный контроллер приложения
func main() {
	beego.Run()
}
