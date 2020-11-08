package logic

type MenuTree struct {
	ID       uint32
	Mname    string
	PID      uint32
	Children []*MenuTree
}

//获取菜单树形结构
func GetMenuTree() {

}
