package init

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	. "beego-training/application/student/common"
)

func InitSql() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:111111@tcp(127.0.0.1:3306)/article?charset=utf8")
	orm.RegisterDataBase(DBTrain, "mysql", "root:111111@tcp(127.0.0.1:3306)/train?charset=utf8")
}
