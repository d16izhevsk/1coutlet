package main

import (
	"github.com/astaxie/beego"
	_ "github.com/d16izhevsk/1coutlet/routers"
)

// main() запускает главный контроллер приложения
func main() {
	beego.Run()
}
