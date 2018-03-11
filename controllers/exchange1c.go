package controllers

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/d16izhevsk/1coutlet/models"
)

var id = 0

// LoadGruppa2 рекурсивно загружает группы в ORM
func LoadGruppa2(Группы []models.ГруппаВложеная2, o orm.Ormer, pid int) {
	for _, группа := range Группы {
		// beego.Info(группы2.Наименование)
		id++
		gruppa := new(models.Gruppa)
		gruppa.Id = id
		gruppa.Idc = группа.Ид
		gruppa.Name = группа.Наименование
		gruppa.Pid = pid
		o.Insert(gruppa)
		// LoadGruppa(группа.Группы.Группа, o)
	}
}

// LoadGruppa рекурсивно загружает группы в ORM
func LoadGruppa(Группы []models.ГруппаВложеная, o orm.Ormer, pid int) {
	for _, группа := range Группы {
		// beego.Info(группы2.Наименование)
		id++
		gruppa := new(models.Gruppa)
		gruppa.Id = id
		gruppa.Idc = группа.Ид
		gruppa.Name = группа.Наименование
		gruppa.Pid = pid
		o.Insert(gruppa)
		LoadGruppa2(группа.Группы.Группа, o, id)
	}
}

// LoadFile загружает файл xml формата 1С обмена commerceml_2, парсит каталог и вносить в базу данных
func LoadFile(filename string) models.КоммерческаяИнформация {

	// beego.Info("Загружаю файл:", filename)
	xmlFile, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	defer xmlFile.Close()
	stat, _ := os.Stat(filename)
	data := make([]byte, stat.Size())
	xmlFile.Read(data)

	var q models.КоммерческаяИнформация
	err = xml.Unmarshal(data, &q)
	if err != nil {
		log.Printf("Ошибка демаршализации: %v\n", err)
		return q
	}
	// beego.Info(q.Каталог.Товары)

	o := orm.NewOrm()
	o.Using("default") // Using default, you can use other database
	for _, группа := range q.Классификатор.Группы.Группа {
		id++
		gruppa := new(models.Gruppa)
		gruppa.Id = id
		gruppa.Idc = группа.Ид
		gruppa.Name = группа.Наименование
		gruppa.Pid = 0
		o.Insert(gruppa)
		LoadGruppa(группа.Группы.Группа, o, id)
	}
	beego.Info("Группы загружены")

	id = 1
	for _, товар := range q.Каталог.Товары.Товар {
		tovar := new(models.Tovar)
		tovar.Id = id
		id++
		tovar.Articul = товар.Артикул
		tovar.Idc = товар.Ид
		tovar.Name = товар.Наименование
		tovar.Opisanie = товар.Описание
		tovar.Gruppa = товар.Группы[0].Ид
		_, err := o.Insert(tovar)
		if err != nil {
			beego.Error(err)
		}
	}
	beego.Info("Товар загружен")

	return q
}

// ExchangeController позволяет обмениваться Get и Post протоколами
type ExchangeController struct {
	beego.Controller
}

// Get контроллера ExchangeController
func (c *ExchangeController) Get() {
	// log := logs.NewLogger(1000)
	// log.SetLogger("console", "")
	// reqtype := c.GetString("type")
	reqmode := c.GetString("mode")
	reqfilename := c.GetString("filename")
	// beego.Info("Request:", reqtype, reqmode, reqfilename)
	switch reqmode {
	case "checkauth":
		// beego.Info("A. Начало сеанса")
		c.TplName = "checkauth.tpl"
	case "init":
		// beego.Info("B. Запрос параметров от сайта")
		c.TplName = "init.tpl"
	case "import":
		// Загрузка из файла в базу данных
		// beego.Info("D. Пошаговая загрузка данных")
		if reqfilename == "import0_1.xml" { // Пока загружаем только товары с артикулами и описаниями
			models.CI = LoadFile(reqfilename)
		}
		c.TplName = "import.tpl"
	default:
		// beego.Error("Ошибка что-то еще")
		c.TplName = "index.tpl"
	}
}

// Post контроллера ExchangeController
func (c *ExchangeController) Post() {

	// beego.Info("C. Выгрузка на сайт файлов обмена")
	reqfilename := c.GetString("filename")
	file, err := os.Create(reqfilename)
	if err == nil {
		n, _ := io.Copy(file, c.Ctx.Request.Body)
		file.Close()
		if err != nil {
			beego.Info(err)
		} else {
			beego.Info("Прочитано байт: ", n, " из ", c.Ctx.Request.ContentLength, "байт файла: ", reqfilename)
		}
	}
	c.TplName = "post.tpl"
}
