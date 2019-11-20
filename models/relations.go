package models

import (
	"github.com/biankemi/bugbang/db"
)

// Relation 关系
type Relation struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//RelationList 关系列表
func (Relation) RelationList() (relations []Relation, err error) {
	db.Conn.Find(&relations)
	return
}
