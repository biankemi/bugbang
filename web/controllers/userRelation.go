package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/biankemi/bugbang/models"
	"github.com/gin-gonic/gin"
)

// UserRelationList 关系列表
func UserRelationList(c *gin.Context) {

	var re models.UserRelation
	userID, _ := strconv.Atoi(c.Param("user_id"))
	list, _ := re.List(userID)
	c.JSON(http.StatusOK, list)
	return
}

// UserRelationCreate 添加关系
func UserRelationCreate(c *gin.Context) {
	var re models.UserRelation
	re.UserID, _ = strconv.Atoi(c.PostForm("user_id"))
	re.ReUserID, _ = strconv.Atoi(c.PostForm("re_user_id"))
	re.RelationID, _ = strconv.Atoi(c.PostForm("relation_id"))
	re.CreateDate = time.Now()
	if err := re.Insert(); err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": "添加失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "添加成功"})
	return
}

// UserRelationDel 解除关系
func UserRelationDel(c *gin.Context) {
	var re models.UserRelation
	ID, _ := strconv.Atoi(c.PostForm("id"))
	err := re.Del(ID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": "解除失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "解除成功"})
	return
}

// UserRelationEdit 改变关系
func UserRelationEdit(c *gin.Context) {
	var re models.UserRelation
	ID, _ := strconv.Atoi(c.PostForm("id"))
	RelationID, _ := strconv.Atoi(c.PostForm("relation_id"))
	err := re.Edit(ID, RelationID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": "解除失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "解除成功"})
	return
}
