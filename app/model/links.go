package model

type Links struct {
	ID     uint32 `json:"id"`
	Name   string `json:"name"`
	Url    string `json:"url"`
	Target string `json:"target"`
	State  uint8  `json:"state"`
	Icon   string `json:"icon"`
	Type   uint8  `json:"type"`
	Ctime  uint32 `json:"ctime"`
	Utime  uint32 `json:"utime"`
}

func (l Links) TableName() string {
	return Prefix + "links"
}