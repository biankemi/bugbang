package models

import (
	"time"

	"github.com/biankemi/bugbang/db"
)

// BroadcastComment 广播
type BroadcastComment struct {
	ID          int       `json:"id"`
	PID         int       `form:"pid" gorm:"default:0;column:pid"`
	BroadcastID int       `form:"broadcast_id"`
	UserID      int       `form:"user_id"`
	Content     string    `form:"content"`
	CreateDate  time.Time `time_format:"2006-01-02 15:04:05" gorm:"default:0000-00-00 00:00:00"`
}

//List 关系列表
func (BroadcastComment) List() (broadcastComments []BroadcastComment, err error) {
	db.Conn.Find(&broadcastComments)
	return
}

// Create 创建
func (bc BroadcastComment) Create() error {
	db.Conn.Create(&bc)
	return nil
}
