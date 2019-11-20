package models

import (
	"time"

	"github.com/biankemi/bugbang/db"
)

// User 用户
type User struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Avator         string    `gorm:"default:''" json:"avator"`
	Integral       int       `gorm:"default:0" json:"integral"`
	Tel            string    `gorm:"default:''" json:"tel"`
	CompanyAddress string    `gorm:"default:''" json:"company_address"`
	HomeAddress    string    `jgorm:"default:''" son:"home_address"`
	RegisterDate   time.Time `json:"register_date"`
}

// Login xxx
func (user User) Login() (info User, err error) {
	count := 0
	if err = db.Conn.Where("tel = ?", user.Tel).First(&info).Count(&count).Error; err != nil {
		return info, err
	}
	if count > 0 {
		return info, nil
	}
	db.Conn.Create(&info)
	return
}

func (user User) getRow(tel string) (userInfo User, err error) {
	if err = db.Conn.Where("tel = ?", tel).First(&userInfo).Error; err != nil {
		return
	}
	return
}

// Edit 修改
func (user User) Edit(id int) error {
	db.Conn.Model(&user).Updates(&user)
	return nil
}
