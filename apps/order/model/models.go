package models

import (
	"errors"
	"go_zero/common/utils"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Mobile int64  `gorm:"index:mobile;type:varchar(13)"`
	Passwd string `gorm:"type:varchar(64)"`
}

// BeforeCreate 在创建前检验验证一下密码的有效性
func (u *User) BeforeCreate(db *gorm.DB) error {
	if len(u.Passwd) < 6 {
		return errors.New("密码太简单了")
	}
	//对密码进行加密存储
	u.Passwd = utils.Password(u.Passwd)
	return nil
}
