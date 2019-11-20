package controllers

import (
	"net/http"

	"github.com/biankemi/bugbang/models"
	"github.com/gin-gonic/gin"
)

// UserRelationList 关系列表
func UserRelationList(c *gin.Context) {
	var re models.UserRelation
	list, _ := re.List()
	c.JSON(http.StatusOK, list)
	return
}

// UserRelationCreate 添加关系
func UserRelationCreate(c *gin.Context) {
	re := models.UserRelation
	if err := c.ShouldBind(&re); err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": "参数出错"})
		return
	}

	if err := re.Insert(); err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": "添加失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "添加成功"})
	return
}
