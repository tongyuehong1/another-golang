package models

import (
	"github.com/astaxie/beego/orm"
)

func Init() {
	orm.RegisterModel(new(First))
}
type First struct{
	Id		int
	Name    string
	Other   string
}