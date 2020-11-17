package model

type Tag struct {
	ID      uint32 `json:"id"`
	TagName string `json:"tag_name"`
}

func (t Tag) TableName() string {
	return Prefix + "tag"
}

//创建

//修改

//删除

//获取一条记录

//获取列表
