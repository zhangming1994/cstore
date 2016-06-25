package controllers

import (
	m "cstore/models"
	"fmt"
	"github.com/astaxie/beego"
)

type SudentController struct {
	CommonController
}

//显示所有的学生列表
func (this *SudentController) ShowAll() {
	if this.IsAjax() {
		sEcho := this.GetString("sEcho")

		iStart, _ := this.GetInt64("iDisplayStart")
		iLength, _ := this.GetInt64("iDisplayLength")

		searchName := this.GetString("sSearch_1") //搜索姓名

		filter := make(map[string]interface{})
		filter["name"] = searchName
		studentlist, count := m.GetAllStudent(iStart, iLength, filter)

		for _, resource := range studentlist {
			switch resource["Sex"] {
			case "0":
				resource["Typename"] = "女"
			case "1":
				resource["Typename"] = "男"
			}
		}

		data := make(map[string]interface{})
		data["aaData"] = studentlist
		data["iTotalDisplayRecords"] = count
		data["iTotalRecords"] = iLength
		data["sEcho"] = sEcho
		this.Data["json"] = &data
		this.ServeJSON()
	}
	this.TplName = "student/index.html"
}

//增加学生
func (this *SudentController) AddStudent() {
	isAction := this.GetString("isAction")
	if "0" == isAction {
		stid, _ := this.GetInt16("stid")
		name := this.GetString("key")
		sex := this.GetString("sex")
		age, _ := this.GetInt16("age")

		beego.Debug("----------------", name, sex, stid, age)
		// st := new(m.Student)
		st := m.SelectStudentByName(name)
		if st.Id == 0 {
			st.Stid = stid
			st.Name = name
			st.Sex = sex
			st.Age = age

			_, err := m.AddStudents(&st)
			if err != nil {
				fmt.Println("插入数据库错误")
				fmt.Println(err)
				return
			}
		} else {
			fmt.Println("已经插入")
			this.Rsp(false, "已经插入")
		}
		this.Ctx.Redirect(302, "/cstore/student/adds.html")
	} else {
		this.TplName = "res/addstu.html"
	}
}

//删除学生
func (this *SudentController) DeleteStudent() {
	Id, _ := this.GetInt64("Id")
	_, err := m.DeleteStudentById(Id)
	if err != nil {
		this.Rsp(false, err.Error())
	} else {
		this.Rsp(true, "Student is Delete Success")
	}
}

//编辑某个学生
func (this *SudentController) UpdateStudent() {
	isAction := this.GetString("isAction")
	Id, _ := this.GetInt64("id")

	if "0" == isAction {
		Name := this.GetString("name")
		Age, _ := this.GetInt16("age")
		Sex := this.GetString("sex")
		Stid, _ := this.GetInt16("stid")

		students := new(m.Student)
		students.Id = Id
		students.Name = Name
		students.Age = Age
		students.Sex = Sex
		students.Stid = Stid
		_, err := m.UpdateSingle(students)
		if err != nil {
			this.Rsp(false, "更新出现问题："+err.Error())
		} else {
			this.Ctx.Redirect(302, "/cstore/student/list.html")
		}
	} else {
		students := m.GetStudentById(Id)
		this.Data["Student"] = students
		this.TplName = "student/edit.html"
	}
}

//根据姓名查找学生
func (this *SudentController) SelectStudent() {
	isAction := this.GetString("isAction")
	if "0" == isAction {
		names := this.GetString("hu")
		a := m.SelectStudentByName(names)
		fmt.Println(a)
	}
	this.TplName = "res/hu.html"
}
