package main

import (
	"github.com/astaxie/beego"
	_ "github.com/d16izhevsk/1coutlet/routers"
)

func init() {
}

// main() запускает главный контроллер приложения
func main() {
	beego.Run()
}
