package model

import "github.com/jinzhu/gorm"

type User struct {
	ID       uint32 `gorm:"primary_key;AUTO_INCREMENT" json:"id"`                              //id
	Username string `gorm:"type:varchar(30);unique_index;not null;default:''" json:"username"` //用户名
	Password string `gorm:"type:varchar(64);not null;default:''" json:"password"`              //密码
	Salt     string `gorm:"type:varchar(10);not null;default:'zadmin'" json:"salt"`            //盐值
	Email    string `gorm:"type:varchar(255);unique_index;not null;default:''" json:"email"`   //邮箱
	Rname    string `gorm:"type:varchar(64);not null;default:''" json:"rname"`                 //昵称
	Ctime    uint32 `gorm:"type:int(10);not null;default:0" json:"ctime"`                      //创建时间
	Utime    uint32 `gorm:"type:int(10);not null;default:0" json:"utime"`                      //更新时间
	Ltime    uint32 `gorm:"type:int(10);not null;default:0" json:"ltime"`                      //登录时间
	IsDel    uint8  `gorm:"type:tinyint(2);not null;default:0" json:"is_del"`                  //是否删除
	Dtime    uint32 `gorm:"type:int(10);not null;default:0" json:"dtime"`                      //删除时间
	Openid   string `gorm:"type:varchar(64);not null;default:''" json:"openid"`                //openid
}

func (u User) TableName() string {
	return pre + "user"
}

//用户列表
func (u User) GetUserList(DB *gorm.DB, page, pageSize int) ([]*User, error) {
	var userSlice []*User
	var offset int
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * pageSize
	}

	//where条件

	if err := DB.Model(&u).Offset(offset).Limit(pageSize).Find(&userSlice).Error; err != nil {
		return userSlice, err
	}
	return userSlice, nil
}
