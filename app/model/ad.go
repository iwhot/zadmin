package model

type Ad struct {
	ID uint32 `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Title string `gorm:"column:title;" json:"title"`
	Url string `gorm:"column:url;" json:"url"`
	PID uint32 `gorm:"column:pid;" json:"pid"`
	Ctime uint32 `gorm:"column:ctime;" json:"ctime"`
	Utime uint32 `gorm:"column:utime;" json:"utime"`
}

func (a Ad) TableName() string {
	return Prefix + "ad"
}