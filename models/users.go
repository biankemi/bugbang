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
func (user User) Login() error {

	db.Conn.Create(&user)
	return nil
}
