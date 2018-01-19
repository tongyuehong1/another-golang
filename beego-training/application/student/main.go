package main

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"

	"github.com/astaxie/beego/orm"
	mysql "github.com/tongyuehong1/another-golang/beego-training/application/student/init"
	"time"
)

func init() {
	orm.RegisterModel(new(Student))
}

type Student struct {
	Id         uint64 `orm:pk`
	String168  string `orm:"size(168)"`
	Int64      int64
	Int32      int32
	String     string
	Bool       bool
	Time       time.Time `orm:"index"`
	Timedate   time.Time `orm:"type(date)"`
	Stringtext string    `orm:"type(text)"`
	Rune       rune
	Float32    float32
	Float64    float64
	Int        int
	Int8       int8
	Uint       uint
	Uint8      uint8
	Uint16     uint16
	Byte       byte
}

func createTable() {
	name := "train"
	force := true
	verbose := true
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		beego.Error(err)
	}
}

func main() {
	mysql.InitSql()
	o := orm.NewOrm()
	o.Using("train")
	createTable()
}
