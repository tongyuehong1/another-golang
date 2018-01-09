package models

import (
	"github.com/astaxie/beego/orm"
)

type StudentServiceProvider struct {
}

var StudentServer *StudentServiceProvider

func init() {
	orm.RegisterModel(new(Student))
}
type Student struct {
	Id   uint64 `form:"id"`
	Name string `form:"name"`
	Age  int    `form:"age"`
	Sex  string `form:"sex"`
}
//
//func (this *StudentServiceProvider) Insert(students []*Student) error {
//	var student Student
//	o := orm.NewOrm()
//	qs := o.QueryTable(students)
//	i, err := qs.PrepareInsert()
//	if err == nil {
//		for _, students = range students {
//			_, err := i.Insert(student)
//			return err
//		}
//		i.Close()
//	}
//	return err
//}

func (this *StudentServiceProvider) Insert(student Student) error {
	o := orm.NewOrm()
	o.Using("Student")
	newstu := new(Student)
	newstu.Name = student.Name
	newstu.Age = student.Age
	newstu.Sex = student.Sex
	_, err := o.Insert(&newstu)
	return err
}
