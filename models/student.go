package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

//定义一个类似Java中间的数据库表的映射
type Student struct {
	Id   int64
	Stid int16
	Name string
	Sex  string
	Age  int16
}

//找到数据库里面这张表
func (s *Student) TableName() string {
	return "student"
}

//注册表模型
func init() {
	orm.RegisterModel(new(Student))
}

//查询所有学生
func GetAllStudent(start int64, length int64, filter map[string]interface{}) (students []orm.Params, count int64) {
	qs := orm.NewOrm().QueryTable("student")

	var cond *orm.Condition
	cond = orm.NewCondition()

	if filter["name"] != nil {
		name := filter["name"].(string)
		cond = cond.And("name__contains", name)
	}

	cond = cond.And("Id__gte", 0)

	qs.Limit(length, start).SetCond(cond).Values(&students)
	count, _ = qs.Count()
	return students, count
	// omodel := orm.NewOrm()
	// st := new(Student)
	// qx := omodel.QueryTable(st)

	// qx.Limit(length, start).Values(&students)
	// count, _ = qx.Count()
	// return students, count
}

//查询
// func Stud() (cout int64, st []orm.Params) {
// 	o := orm.NewOrm()
// 	std := new(Student)
// 	qx := o.QueryTable(st)
// 	count, _ := qx.Values(&st, "ID", "stid", "name", "sex", "age")
// 	return count, std
// }

//增加学生
func AddStudents(s *Student) (int64, error) {
	omodel := orm.NewOrm()
	st := new(Student)
	st.Stid = s.Stid
	st.Age = s.Age
	st.Name = s.Name
	st.Sex = s.Sex
	id, err := omodel.Insert(st)
	return id, err
}

//更新某个学生信息（编辑学生信息）
func UpdateSingle(s *Student) (int64, error) {

	se := orm.NewOrm()
	stud := make(orm.Params)
	if len(s.Name) >= 0 {
		stud["Name"] = s.Name
	}
	if s.Age > 0 {
		stud["Age"] = s.Age
	}
	if len(s.Sex) > 0 {
		stud["Sex"] = s.Sex
	}
	if s.Stid > 0 {
		stud["Stid"] = s.Stid
	}
	if s.Id > 0 {
		stud["Id"] = s.Id
	}
	var table Student
	num, err := se.QueryTable(table).Filter("Id", s.Id).Update(stud)
	return num, err
}

//删除某个学生
func DeleteStudentById(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&Student{Id: Id})
	return status, err
}

//根据姓名查询某一个学生
func SelectStudentByName(sname string) (students Student) {
	o := orm.NewOrm()
	st := Student{Name: sname}
	err := o.Read(&st)
	if err == orm.ErrNoRows {
		fmt.Printf("\n  查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("\n 找不到主键")
	} else {
		fmt.Println(err)
		fmt.Println("\n")
		fmt.Println(st.Id, st.Name, st.Sex, st.Stid, st.Age)
	}
	return st
}

//根据ID得到某一个学会少年宫
func GetStudentById(studentId int64) (students Student) {
	fmt.Printf("student id is %d", studentId)
	o := orm.NewOrm()
	sthu := new(Student)
	err := o.QueryTable(sthu).One(&students)
	if err != nil {
		fmt.Println(err)
	}
	return students
}
