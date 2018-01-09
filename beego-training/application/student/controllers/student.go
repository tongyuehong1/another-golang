package controllers

import (
	"beego-training/application/student/models"
	"beego-training/libs/logger"
	"github.com/astaxie/beego"
	"github.com/tongyuehong1/golang-project/application/blog/common"
	"fmt"
)

type StudentController struct {
	beego.Controller
}

//func (this *StudentController) Insert() {
//	stu := models.Student{}
//	if err := this.ParseForm(&stu); err != nil {
//		logger.Logger.Error("解析form数据出错：", err)
//		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
//	} else {
//		err := models.StudentServer.Insert(stu)
//		if err != nil {
//			logger.Logger.Error("Insert ", err)
//			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
//		}
//	}
//}
func (this *StudentController) Show() {
	stu := models.Student{}
	if err := this.ParseForm(&stu); err != nil {
		logger.Logger.Error("解析form数据出错：", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		this.Data["json"] = map[string]interface{}{"content":stu}
	}
	this.ServeJSON()
	fmt.Printf("get the data  %v\n",stu)
}
