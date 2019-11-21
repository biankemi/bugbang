package models

import (
	"time"

	"github.com/biankemi/bugbang/db"
)

// BroadcastReward 广播
type BroadcastReward struct {
	ID          int       `json:"id"`
	BroadcastID int       `form:"broadcast_id"`
	UserID      int       `form:"user_id"`
	Integral    int       `form:"integral"`
	Prove       string    `form:"prove"`
	CreateDate  time.Time `time_format:"2006-01-02 15:04:05"`
}

//List 关系列表
func (BroadcastReward) List() (broadcastRewards []BroadcastReward, err error) {
	db.Conn.Find(&broadcastRewards)
	return
}

// Create 创建
func (br BroadcastReward) Create() error {
	db.Conn.Create(&br)
	return nil
}
