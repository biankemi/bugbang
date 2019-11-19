package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Conn mysql connection
var Conn *gorm.DB

func init() {
	var err error
	Conn, err = gorm.Open("mysql", "lisi:3q213Q21lisi#@tcp(122.152.220.87:3306)/bugbang?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	Conn.DB().SetMaxIdleConns(10)
	Conn.DB().SetMaxOpenConns(100)
}
