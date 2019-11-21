package models

import (
	"time"

	"github.com/biankemi/bugbang/db"
)

// UserRelation 用户关系
type UserRelation struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	ReUserID   int       `json:"re_user_id"`
	RelationID int       `json:"relation_id"`
	CreateDate time.Time `json:"create_date" gorm:"default:'0000-00-00 00:00:00'"` //'创建日期'
	DeleteDate time.Time `json:"delete_date" gorm:"default:'0000-00-00 00:00:00'"` //'创建日期'
}

//List 关系列表
func (ur UserRelation) List(userID int) (userRelations []UserRelation, err error) {
	db.Conn.Where("user_id = ?", userID).Find(&userRelations)
	return
}

// Insert 增加
func (ur *UserRelation) Insert() error {
	err := db.Conn.Create(&ur)
	if err != nil {
		return err.Error
	}
	return nil
}

// Edit 编辑
func (ur *UserRelation) Edit(id int, relationID int) error {
	err := db.Conn.Model(&ur).Where("id = ?", id).Update("relation_id", relationID)
	if err != nil {
		return err.Error
	}
	return nil
}

// Del 编辑
func (ur *UserRelation) Del(id int) error {
	deleteDate := time.Now()
	err := db.Conn.Model(&ur).Where("id = ?", id).Update("delete_date", deleteDate)
	if err != nil {
		return err.Error
	}
	return nil
}
