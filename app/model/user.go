package model

import (
	"errors"
	page2 "github.com/iwhot/zadmin/system/page"
	"github.com/jinzhu/gorm"
	"time"
)

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
	var offset = page2.GetOffset(page, pageSize)
	var mod = DB.Model(&u)
	//where条件
	if u.Username != "" {
		mod.Where("username like ?", "%"+u.Username+"%")
	}

	if u.Email != "" {
		mod.Where("email like ?", "%"+u.Email+"%")
	}

	if u.ID != 0 {
		mod.Where("id = ?", u.ID)
	}

	if err := mod.Where("is_del = ?",0).Offset(offset).Limit(pageSize).Order("utime desc,id desc").Find(&userSlice).Error; err != nil {
		return userSlice, err
	}

	return userSlice, nil
}

//添加用户
func (u User) AddUser(DB *gorm.DB) error {
	return DB.CreateTable(&u).Error
}

//更新用户
func (u User) UpdateUser(DB *gorm.DB) error {
	return DB.Model(&u).Updates(u).Error
}

//删除用户
func (u User) DeleteUser(DB *gorm.DB) error {
	_, err := u.GetOneUserInfo(DB)
	if err != nil {
		return err
	}
	u.IsDel = 1
	u.Dtime = uint32(time.Now().Unix())
	return u.UpdateUser(DB)
}

//获取一个用户信息
func (u User) GetOneUserInfo(DB *gorm.DB) (*User, error) {
	if u.ID == 0 {
		return nil, errors.New("用户不存在")
	}

	users, err := u.GetUserList(DB, 0, 1)
	if err != nil {
		return nil, err
	}

	return users[0], nil
}
