package logic

import "github.com/iwhot/zadmin/app/model"

type RoleMenuTree struct {
	ID       uint32         `json:"id"`
	Title    string         `json:"title"`
	Field    string         `json:"field"`
	Checked  bool           `json:"checked"`
	Spread   bool           `json:"spread"`
	Disabled bool           `json:"disabled"`
	Children []RoleMenuTree `json:"children"`
}

//获取树形节点
func GetRoleMenuNode(rm []*model.Menu, pid uint32,chk []uint32) []RoleMenuTree {
	var rmt = []RoleMenuTree{}

	for _, v := range rm {
		if v.PID == pid {
			child := GetRoleMenuNode(rm, v.ID,chk)
			node := RoleMenuTree{
				ID:       v.ID,
				Title:    v.Mname,
				Field:    "menu[]",
				Spread:   false,
				Disabled: false,
			}

			for _,item := range chk {
				if item == v.ID{
					node.Checked = true
				}
			}


			node.Children = child
			rmt = append(rmt, node)
		}
	}

	return rmt
}
