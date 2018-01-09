package main

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"

	mysql "github.com/tongyuehong1/golang-project/application/blog/init"
)

func main() {
	mysql.InitSql()
	beego.Run()
}
