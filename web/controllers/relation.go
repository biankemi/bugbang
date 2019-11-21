package controllers

import (
	"net/http"

	"github.com/biankemi/bugbang/models"
	"github.com/gin-gonic/gin"
)

var r models.Relation

// RelationList 关系列表
func RelationList(c *gin.Context) {

	list, _ := r.List()
	c.JSON(http.StatusOK, list)
	return
}
