package logic

import (
	"github.com/iwhot/zadmin/app/dao"
	"github.com/iwhot/zadmin/app/model"
)

//菜单树结构体
type MenuTree struct {
	ID       uint32
	Mname    string
	PID      uint32
	Type     uint8
	Children []MenuTree
}

//面包屑
type Breadcrumbs struct {
	ID    uint32
	Mname string
}

//递归获取菜单树形结构
func GetMenuTree(m []*model.Menu, pid uint32) []MenuTree {
	treeList := []MenuTree{}

	for _, v := range m {
		if v.PID == pid {
			child := GetMenuTree(m, v.ID)
			node := MenuTree{
				ID:    v.ID,
				Mname: v.Mname,
				PID:   v.PID,
				Type:  v.Type,
			}
			node.Children = child
			treeList = append(treeList, node)
		}
	}

	return treeList
}

//面包屑导航
func GetMenuBreadcrumbs(id uint32) []Breadcrumbs {
	var b = []Breadcrumbs{}

	var m1, err = dao.DefaultMenuDao.FindOneToID(id)
	if err != nil {
		return b
	}

	node := Breadcrumbs{
		ID:    m1.ID,
		Mname: m1.Mname,
	}

	b = append(b, node)

	if m1.PID != 0 {
		b2 := GetMenuBreadcrumbs(m1.PID)
		b = append(b, b2...)
	}

	return b
}

func GetMenuBreadcrumbsFan(id uint32) []Breadcrumbs {
	arr := GetMenuBreadcrumbs(id)
	length := len(arr)

	for i := 0; i < length/2; i++ {
		temp := arr[length-1-i]
		arr[length-1-i] = arr[i]
		arr[i] = temp
	}

	return arr
}
