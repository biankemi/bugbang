package models

import (
	"time"

	"github.com/biankemi/bugbang/db"
)

// UserRelation 用户关系
type UserRelation struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	ReUserID    int       `json:"re_user_id"`
	RelationsID int       `json:"relations_id"`
	CreateDate  time.Time `json:"create_date"`
}

//List 关系列表
func (ur UserRelation) List(userID int) (userRelations []UserRelation, err error) {
	db.Conn.Where("user_id", userID).Find(&userRelations)
	return
}

// Insert 增加
func (ur *UserRelation) Insert() error {
	db.Conn.Save(&ur)
	return nil
}
