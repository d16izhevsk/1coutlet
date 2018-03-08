package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:123456@/1coutlet?charset=utf8")
}

// Tovar содержит базу данных
type Tovar struct {
	Id       int
	Idc      string
	Articul  string
	Name     string
	Opisanie string
	Gruppa   string
}

// Gruppa содержит базу данных
type Gruppa struct {
	Id   int
	Idc  string
	Name string
}

func init() {
	// Need to register model in init
	orm.RegisterModel(new(Tovar), new(Gruppa))
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

// SecretCode переменная для вывода как счетчик
var SecretCode = 1

// CI база данных о товарах
var CI КоммерческаяИнформация
