package model

type Catalog struct {
	ID uint32 `json:"id"`
	Name string `json:"name"`
	Ename string `json:"ename"`
	Icon string `json:"icon"`
	Url string `json:"url"`
	Target string `json:"target"`
	State uint8 `json:"state"`
	Ctime uint32 `json:"ctime"`
	Utime uint32 `json:"utime"`
}

func (c Catalog) TableName() string {
	return Prefix + "catalog"
}