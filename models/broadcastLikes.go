package models

import (
	"time"

	"github.com/biankemi/bugbang/db"
)

// BroadcastLike 广播
type BroadcastLike struct {
	ID          int       `json:"id"`
	BroadcastID int       `form:"broadcast_id"`
	UserID      int       `form:"user_id"`
	CreateDate  time.Time `time_format:"2006-01-02 15:04:05"`
}

//List 关系列表
func (BroadcastLike) List() (broadcastLikes []BroadcastLike, err error) {
	db.Conn.Find(&broadcastLikes)
	return
}

// Create 创建
func (bl BroadcastLike) Create() error {
	db.Conn.Create(&bl)
	return nil
}
