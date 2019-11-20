package controllers

import (
	"net/http"

	"github.com/biankemi/bugbang/models"
	"github.com/gin-gonic/gin"
)

// RelationList 关系列表
func RelationList(c *gin.Context) {
	var re models.Relation
	list, _ := re.List()
	c.JSON(http.StatusOK, list)
	return
}
