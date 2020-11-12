package model

type Category struct {
	ID uint32 `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	PID uint32 `gorm:"column:pid;" json:"pid"`
	Name string `json:"name"`
	Ename string `json:"ename"`
	Icon string `json:"icon"`
	Remark string `json:"remark"`
	Ctime uint32 `json:"ctime"`
	Utime uint32 `json:"utime"`
	SeoTitle string `json:"seo_title"`
	SeoKwds string `json:"seo_kwds"`
	SeoDesc string `json:"seo_desc"`
	Thumb string `json:"thumb"`
}
