package controllers

import (
	"net/http"

	"github.com/biankemi/bugbang/models"
	"github.com/gin-gonic/gin"
)

// List 关系列表
func List(c *gin.Context) {

	list := models.Relation.RelationList()
	c.SecureJSON(http.StatusOK, list)
	return
}
