package controllers

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
)

// Tovar элемент базы данных товар
type Tovar struct {
	gorm gorm.Model
	ID   string
	Name string
}

// ГруппаВложеная2 подгруппа
type ГруппаВложеная2 struct {
	Ид           string
	Наименование string
}

// ГруппыВложенные2 структура от 1С
type ГруппыВложенные2 struct {
	Группа []ГруппаВложеная2
}

// ГруппаВложеная структура от 1С
type ГруппаВложеная struct {
	Ид           string
	Наименование string
	Группы       ГруппыВложенные2
}

// ГруппыВложеные странная структура в 1с файле
type ГруппыВложеные struct {
	Группа []ГруппаВложеная
}

// Группа товаров вложенная
type Группа struct {
	Ид           string
	Наименование string
	Группы       ГруппыВложеные
}

// Группы описания групп
type Группы struct {
	Ид     string
	Группа []Группа
}

// Изготовитель по версии 1С
type Изготовитель struct {
	Ид           string
	Наименование string
}

// Пересчет единица измерения товара
type Пересчет struct {
	Единица     int
	Коэффициент float32
}

// БазоваяЕдиница минимальная единица товара
type БазоваяЕдиница struct {
	Пересчет []Пересчет
}

// ЗначениеРеквизита из 1с
type ЗначениеРеквизита struct {
	Наименование string
	Значение     string
}

// ЗначенияРеквизитов конкретные из 1с
type ЗначенияРеквизитов struct {
	ЗначениеРеквизита []ЗначениеРеквизита
}

// Товар основная единица для обработки
type Товар struct {
	Ид                 string
	Артикул            string
	Наименование       string
	Группы             []Группы
	Описание           string
	Изготовитель       Изготовитель
	БазоваяЕдиница     БазоваяЕдиница
	ЗначенияРеквизитов ЗначенияРеквизитов
}

// Товары слайс товаров
type Товары struct {
	Товар []Товар
}

// Каталог товаров из 1с
type Каталог struct {
	Ид               string
	ИдКлассификатора string
	Наименование     string
	Товары           Товары
}

// Владелец каталога товаров
type Владелец struct {
	Ид                      string
	Наименование            string
	ОфициальноеНаименование string
	ИНН                     string
	КПП                     string
	ОКПО                    string
}

// Классификатор содержит фирмы и группы товаров
type Классификатор struct {
	Наименование string
	Владелец     Владелец
	Группы       Группы
}

// КоммерческаяИнформация основной пакет данных из 1С
type КоммерческаяИнформация struct {
	Классификатор Классификатор
	Каталог       Каталог
}

// LoadFile загружает файл xml формата 1С обмена commerceml_2, парсит каталог и вносить в базу данных
func LoadFile(filename string) КоммерческаяИнформация {

	xmlFile, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	defer xmlFile.Close()
	stat, _ := os.Stat(filename)
	data := make([]byte, stat.Size())
	xmlFile.Read(data)

	var q КоммерческаяИнформация
	err = xml.Unmarshal(data, &q)
	if err != nil {
		log.Printf("Ошибка демаршализации: %v\n", err)
		return q
	}
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
	reqtype := c.GetString("type")
	reqmode := c.GetString("mode")
	reqfilename := c.GetString("filename")
	beego.Info("Request:", reqtype, reqmode, reqfilename)
	switch reqmode {
	case "checkauth":
		// beego.Info("A. Начало сеанса")
		c.TplName = "checkauth.tpl"
	case "init":
		// beego.Info("B. Запрос параметров от сайта")
		c.TplName = "init.tpl"
	case "import":
		// beego.Info("D. Пошаговая загрузка данных")
		// Загрузка из файла в базу данных
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
		n, err := io.Copy(file, c.Ctx.Request.Body)
		file.Close()
		if err != nil {
			beego.Info(err)
		} else {
			beego.Info("Прочитано байт: ", n, " из ", c.Ctx.Request.ContentLength, "байт файла: ", reqfilename)
		}
	}
	c.TplName = "post.tpl"
}
