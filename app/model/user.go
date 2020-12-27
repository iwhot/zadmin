package model

import (
	"errors"
	page2 "github.com/iwhot/zadmin/system/page"
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	ID       uint32 `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`                                    //id
	Username string `gorm:"column:username;type:varchar(30);unique_index;not null;default:''" json:"username"` //用户名
	Password string `gorm:"column:password;type:varchar(64);not null;default:''" json:"password"`              //密码
	Salt     string `gorm:"column:salt;type:varchar(10);not null;default:'zadmin'" json:"salt"`                //盐值
	Email    string `gorm:"column:email;type:varchar(255);unique_index;not null;default:''" json:"email"`      //邮箱
	Rname    string `gorm:"column:rname;type:varchar(64);not null;default:''" json:"rname"`                    //昵称
	Ctime    uint32 `gorm:"column:ctime;type:int(10);not null;default:0" json:"ctime"`                         //创建时间
	Utime    uint32 `gorm:"column:utime;type:int(10);not null;default:0" json:"utime"`                         //更新时间
	Ltime    uint32 `gorm:"column:ltime;type:int(10);not null;default:0" json:"ltime"`                         //登录时间
	IsDel    uint8  `gorm:"column:is_del;type:tinyint(2);not null;default:0" json:"is_del"`                    //是否删除
	Dtime    uint32 `gorm:"column:dtime;type:int(10);not null;default:0" json:"dtime"`                         //删除时间
	Openid   string `gorm:"column:openid;type:varchar(64);not null;default:''" json:"openid"`                  //openid
	Avatar   string `gorm:"column:avatar;type:varchar(255);unique_index;not null;default:''" json:"avatar"`    //头像
	State    uint8  `gorm:"column:state;type:tinyint(2);not null;default:1" json:"state"`                      //状态
	Desc     string `gorm:"column:desc;type:varchar(1000);unique_index;not null;default:''" json:"desc"`       //描述
	Role     Role   //user属于role，role_id为外键
	RoleID   uint32
}

func (u User) TableName() string {
	return Prefix + "user"
}

//用户列表
func (u User) GetUserList(DB *gorm.DB, page, pageSize int) ([]*User, error) {
	var userSlice []*User
	var offset = page2.GetOffset(page, pageSize)
	var mod *gorm.DB

	mod = DB.Model(&u)

	if u.ID != 0 {
		mod = mod.Where("id = ?", u.ID)
	} else {
		//where条件
		if u.Username != "" {
			mod = mod.Where("username like ?", "%"+u.Username+"%")
		}

		if u.Rname != "" {
			mod = mod.Where("rname like ?", "%"+u.Rname+"%")
		}

		if u.Email != "" {
			mod = mod.Where("email like ?", "%"+u.Email+"%")
		}
	}

	if err := mod.Where("is_del = ?", 0).Offset(offset).Limit(pageSize).Order("utime desc,id desc").Find(&userSlice).Error; err != nil {
		return userSlice, err
	}

	return userSlice, nil
}

//添加用户
func (u User) AddUser(DB *gorm.DB) error {
	tx := DB.Begin()
	//用户入库
	if err := tx.Create(&u).Error; err != nil {
		tx.Rollback()
		return err
	}
	//todo 加入文件管理器
	tx.Commit()
	return nil
}

//更新用户
func (u User) UpdateUser(DB *gorm.DB) error {
	tx := DB.Begin()

	//更新用户
	if err := tx.Model(&u).Updates(u).Error; err != nil {
		tx.Rollback()
		return err
	}

	//todo 加入文件管理器

	tx.Commit()
	return nil
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

//获取记录数量
func (u User) Count(DB *gorm.DB) int {
	var count int
	var mod = DB.Model(&u)

	if u.ID != 0 {
		mod.Where("id = ?", u.ID)
	} else {
		//where条件
		if u.Username != "" {
			mod.Where("username like ?", "%"+u.Username+"%")
		}

		if u.Rname != "" {
			mod.Where("rname like ?", "%"+u.Rname+"%")
		}

		if u.Email != "" {
			mod.Where("email like ?", "%"+u.Email+"%")
		}
	}

	if err := mod.Count(&count).Error; err != nil {
		return 0
	}

	return count
}
