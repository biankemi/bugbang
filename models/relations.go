package models

import (
	"github.com/biankemi/bugbang/db"
)

// Relation 关系
type Relation struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//List 关系列表
func (r Relation) List() (relations []Relation, err error) {
	db.Conn.Find(&relations)
	return
}
