package models

import (
	"time"

	"github.com/biankemi/bugbang/db"
	"github.com/jinzhu/gorm"
)

// Broadcast 广播
type Broadcast struct {
	ID            int    `json:"id"`
	Title         string `form:"title" json:"title" gorm:"column:title"`
	Content       string `form:"content" json:"content" gorm:"column:content"`
	Distance      string `form:"distance" json:"distance" gorm:"column:distance"`
	Longitude     string `form:"longitude" json:"longitude" gorm:"column:longitude"`
	Latitude      string `form:"latitude" json:"latitude" gorm:"column:latitude"`
	Integral      int    `form:"integral" json:"integral" gorm:"column:integral"`
	Num           int    `form:"num" json:"num" gorm:"column:num"`
	EffectiveTime string `form:"effective" json:"effective" gorm:"column:effective"`
	IsRandom      int    `form:"is_random" json:"is_random" gorm:"column:is_random"`
	CategoryID    int    `form:"category_id" json:"category_id" gorm:"column:category_id"`
	UserID        int    `form:"user_id" json:"user_id" gorm:"column:user_id"`
	ViewCount     int    `form:"view_count" json:"view_count" gorm:"column:view_count"`
	LikeCount     int    `form:"like_count" json:"like_count" gorm:"column:like_count"`
	UnlikeCount   int    `form:"unlike_count" json:"unlike_count" gorm:"column:unlike_count"`
	CommentCount  int    `form:"comment_count" json:"comment_count" gorm:"column:comment_count"`
	CreateDate    time.Time
}

//List 关系列表
func (b Broadcast) List() (broadcasts []Broadcast, err error) {
	db.Conn.Find(&broadcasts)
	return
}

// Create 创建
func (b Broadcast) Create() error {
	db.Conn.Create(&b)
	return nil
}

// Edit 编辑
func (b Broadcast) Edit() error {
	err := db.Conn.Model(&b).Updates(&b)
	if err != nil {
		return err.Error
	}
	return nil
}

// Incr 递增
func (b Broadcast) Incr(ID int, incrType string) (count int, msg string) {

	tmpStr := " + 1"
	if err := db.Conn.Model(&b).Where("id = ?", ID).Update(incrType, gorm.Expr(incrType+tmpStr)); err != nil {
		return 0, "写入出错"
	}
	return 1, "添加成功"
}
