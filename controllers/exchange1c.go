package controllers

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
)

// Tovar элемент базы данных товар
type Tovar struct {
	gorm gorm.Model
	Id   string
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

// -
type ГруппыВложеные struct {
	Группа []ГруппаВложеная
}

// -
type Группа struct {
	Ид           string
	Наименование string
	Группы       ГруппыВложеные
}

// -
type Группы struct {
	Ид     string
	Группа []Группа
}

// -
type Изготовитель struct {
	Ид           string
	Наименование string
}

// -
type Пересчет struct {
	Единица     int
	Коэффициент float32
}

// -
type БазоваяЕдиница struct {
	Пересчет []Пересчет
}

// -
type ЗначениеРеквизита struct {
	Наименование string
	Значение     string
}

// -
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
func LoadFile() {

	xmlFile, err := os.Open("test.xml")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	defer xmlFile.Close()
	stat, _ := os.Stat("test.xml")
	data := make([]byte, stat.Size())
	xmlFile.Read(data)

	var q КоммерческаяИнформация
	err = xml.Unmarshal(data, &q)
	if err != nil {
		fmt.Printf("Ошибка демаршализации: %v\n", err)
		return
	}
	fmt.Printf("Каталог: %#v\n", q.Классификатор)

	for _, tov := range q.Каталог.Товары.Товар {
		fmt.Printf("Товар: %#v\n\n", tov)
		// fmt.Printf("Артикул:%s\tНаименование: %s\tИзготовитель:%s\n", tov.Артикул, tov.Наименование, tov.Изготовитель.Наименование)
	}
}

// ExchangeController позволяет обмениваться Get и Post протоколами
type ExchangeController struct {
	beego.Controller
}

func (c *ExchangeController) Get() {
	c.TplName = "catalog.tpl"
}

func (c *ExchangeController) Post() {
	c.TplName = "catalog.tpl"
}
