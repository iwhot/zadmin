package model

type Article struct {
	ID         uint32 `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Title      string `gorm:"column:title;" json:"title"`
	ShortTitle string `gorm:"column:short_title;" json:"short_title"`
	CategoryID uint32 `gorm:"column:category_id;" json:"category_id"`
	Content    string `gorm:"column:content;" json:"content"`
	Remark     string `gorm:"column:remark;" json:"remark"`
	Thumb      string `gorm:"column:thumb;" json:"thumb"`
	Ctime      uint32 `gorm:"column:ctime;" json:"ctime"`
	Utime      uint32 `gorm:"column:utime;" json:"utime"`
	UserID     uint32 `gorm:"column:user_id;" json:"user_id"`
	Click      uint32 `gorm:"column:click;" json:"click"`
	Like       uint32 `gorm:"column:like;" json:"like"`
	Type       uint8  `gorm:"column:type;" json:"type"`
	State      uint8  `gorm:"column:state;" json:"state"`
	SeoTitle   string `gorm:"column:seo_title;" json:"seo_title"`
	SeoKwds    string `gorm:"column:seo_kwds;" json:"seo_kwds"`
	SeoDesc    string `gorm:"column:seo_desc;" json:"seo_desc"`
	Uuid       string `gorm:"column:uuid;" json:"uuid"`
}

func (a Article) TableName() string {
	return Prefix + "article"
}

//创建

//修改

//删除

//获取一条记录

//获取列表

//点赞

//喜欢
