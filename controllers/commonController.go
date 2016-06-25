package controllers

import (
	"fmt"

	m "cstore/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type CommonController struct {
	beego.Controller
}

func init() {
	fmt.Println("register access is start \n")
	m.CheckAccessAndRegisterRes()
}

func (this *CommonController) Rsp(status bool, str string) {
	this.Data["json"] = &map[string]interface{}{"status": status, "info": str}
	this.ServeJSON()
}

//
func (this *CommonController) GetTree() []Tree {
	resources, _ := m.GetResourceTree(0, 1)
	tree := make([]Tree, len(resources))
	for k, v := range resources {
		tree[k].Id = v["Id"].(int64)
		tree[k].Text = v["Title"].(string)
		children, _ := m.GetResourceTree(v["Id"].(int64), 2)
		tree[k].Children = make([]Tree, len(children))
		for k1, v1 := range children {
			tree[k].Children[k1].Id = v1["Id"].(int64)
			tree[k].Children[k1].Text = v1["Title"].(string)
			tree[k].Children[k1].Url = "/" + v["Name"].(string) + "/" + v1["Name"].(string)
		}
	}
	return tree
}
func (this *CommonController) GetResList(uname string, Id int64) []Tree { //这里的ID是角色ID
	var cnt, length int = 0, 0 //cnt代表的是级别，length代表的是数量
	var resources []orm.Params
	adminUser := beego.AppConfig.String("admin_user")
	if uname == adminUser {
		_, resources = m.GetAllResource() //判断是不是管理员，采用不同的取得资源的方式
	} else {
		resources, _ = m.GetResourcesByRoleId(Id)
	}

	for _, v := range resources {
		if v["Pid"].(int64) == 0 { //从接口里面取出来的数据需要转换成自己需要的数据类型
			length = length + 1 //这里是统计以及目录的数量，统计到一个增加一个
		}
	}

	//length
	tree := make([]Tree, length) //建立一个有长度的数组,
	for k, v := range resources {
		fmt.Printf("\n key is %d and id  is %d Pid is %d\n", k, v["Id"], v["Pid"])
		if v["Pid"].(int64) == 0 {
			k = cnt
			cnt = cnt + 1
			tree[k].Id = v["Id"].(int64)
			tree[k].Index = cnt
			tree[k].Text = v["Name"].(string)
			// 1代表菜单（目录下面的所有资源）没有把一些不需要的权限去掉
			// children, _ := m.GetResourceTree(v["Id"].(int64), 1)

			var childCnt int = 0
			children := make([]map[string]interface{}, 4)
			for _, v3 := range resources {
				if v3["Pid"].(int64) == v["Id"].(int64) {
					children[childCnt] = v3
					childCnt++
				}
			}

			tree[k].Children = make([]Tree, childCnt)
			for k1, v1 := range children {
				fmt.Printf("\n kv1 is %d\n", v1)
				if v1 == nil {
					fmt.Printf("\n data is over is %d\n")
				} else {
					if v1["Pid"].(int64) == v["Id"].(int64) {
						tree[k].Children[k1].Id = v1["Id"].(int64)
						tree[k].Children[k1].Text = v1["Name"].(string)
						tree[k].Children[k1].Url = "/" + v1["Url"].(string)
					}
				}
			}

		}

	}
	return tree
}
