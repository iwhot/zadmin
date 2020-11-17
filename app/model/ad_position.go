package model

type AdPosition struct {
	ID uint32 `gorm:"column:id;primary_key;AUTO_INCREMENT"json:"id"`
	Name string `gorm:"column:name;" json:"name"`
	Width uint32 `gorm:"column:width;"json:"width"`
	Height uint32 `gorm:"column:height;"json:"height"`
	Type uint8 `gorm:"column:type;"json:"type"`
	Ctime uint32 `gorm:"column:ctime;"json:"ctime"`
	Utime uint32 `gorm:"column:utime;"json:"utime"`
	Desc string `gorm:"column:desc;"json:"desc"`
	Style string `gorm:"column:style;"json:"style"`
}

func (a AdPosition) TableName() string {
	return Prefix + "ad_position"
}