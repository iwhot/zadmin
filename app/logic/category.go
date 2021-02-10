package logic

import "github.com/iwhot/zadmin/app/model"

type CategoryTree struct {
	ID       uint32
	PID      uint32
	Name     string
	Ename    string
	Icon     string
	Remark   string
	State    uint8
	Ctime    uint32
	Utime    uint32
	Children []CategoryTree
}

//获取分类树
func GetCategoryTree(cate []*model.Category, pid uint32) ([]CategoryTree) {
	var cates = []CategoryTree{}

	for _, v := range cate {
		if v.PID == pid {
			child := GetCategoryTree(cate, v.ID)
			node := CategoryTree{
				ID:     v.ID,
				PID:    v.PID,
				Name:   v.Name,
				Ename:  v.Ename,
				Icon:   v.Icon,
				Remark: v.Remark,
				State:  v.State,
				Ctime:  v.Ctime,
				Utime:  v.Utime,
			}

			node.Children = child
			cates = append(cates, node)
		}
	}

	return cates
}
