package model

import (
	"errors"
	"gorm.io/gorm"
)

// User
// 系统登录用户
type User struct {
	gorm.Model
	Username string `gorm:"unique;<-:create;size:20;not null" json:"-"`
	Nickname string `gorm:"size:30;not null"`
	Password string `gorm:"size:32;not null" json:"-"`
	Role     int    `gorm:"type:tinyint;default:0"` //0普通：可进行违约认定，重生认定；1管理：对认定进行审核
}

func (u *User) Login() error {
	u.Password = Md5(u.Password)
	err := db.Model(User{}).Where("username = ? AND password = ?", u.Username, u.Password).First(u).Error
	if err != nil {
		return errors.New("用户名或密码错误")
	}
	return nil
}

// GetInfo
// 查询系统用户信息
func (u *User) GetInfo() error {
	err := db.Where("id = ?", u.ID).First(u).Error
	if err != nil {
		return errors.New("用户名或密码错误")
	}
	return nil
}

func PreInsertUser() {
	users := []User{
		{
			Username: "admin",
			Nickname: "管理员",
			Role:     1,
			Password: Md5("admin"),
		},
		{
			Username: "user",
			Nickname: "普通用户",
			Role:     0,
			Password: Md5("user"),
		},
	}
	db.CreateInBatches(&users, len(users))
}
