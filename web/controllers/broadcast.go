package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/biankemi/bugbang/models"
	"github.com/gin-gonic/gin"
)

// BroadcastList 广播列表
func BroadcastList(c *gin.Context) {
	var b models.Broadcast
	list, _ := b.List()
	c.JSON(http.StatusOK, list)
	return
}

// BroadcastCreate 创建广播
func BroadcastCreate(c *gin.Context) {
	var b models.Broadcast
	b.Title = c.PostForm("title")
	b.Content = c.PostForm("content")
	b.Distance = c.PostForm("distance")
	b.Longitude = c.PostForm("longitude")
	b.Latitude = c.PostForm("latitude")
	b.Integral, _ = strconv.Atoi(c.PostForm("integral"))
	b.Num, _ = strconv.Atoi(c.PostForm("num"))
	b.EffectiveTime = c.PostForm("effective_time")
	b.IsRandom, _ = strconv.Atoi(c.PostForm("is_random"))
	b.CategoryID, _ = strconv.Atoi(c.PostForm("category_id"))
	b.UserID, _ = strconv.Atoi(c.PostForm("user_id"))
	b.ViewCount, _ = strconv.Atoi(c.PostForm("view_count"))
	b.LikeCount, _ = strconv.Atoi(c.PostForm("like_count"))
	b.UnlikeCount, _ = strconv.Atoi(c.PostForm("unlike_count"))
	b.CommentCount, _ = strconv.Atoi(c.PostForm("comment_count"))
	b.CreateDate = time.Now()
	b.Create()
}

// BroadcastEdit 编辑广播
func BroadcastEdit(c *gin.Context) {
	var form models.Broadcast

	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusOK, gin.H{"err_code": 0, "err_msg": err.Error()})
		return
	}
	form.ID, _ = strconv.Atoi(c.Param("id"))

	if err := form.Edit(); err != nil {
		c.JSON(http.StatusOK, gin.H{"err_code": 0, "err_msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"err_code": 0, "err_msg": "success"})
	return
}

// DoLike 喜欢
func DoLike(c *gin.Context) {
	bID, _ := strconv.Atoi(c.PostForm("id"))
	userID, _ := strconv.Atoi(c.PostForm("user_id"))

	var bl models.BroadcastLike
	bl.BroadcastID = bID
	bl.UserID = userID
	bl.CreateDate = time.Now()
	if err := bl.Create(); err != nil {
		c.JSON(http.StatusOK, gin.H{"err_code": 1, "err_msg": err.Error()})
		return
	}
	var b models.Broadcast
	b.Incr(bID, "like_count")

	c.JSON(http.StatusOK, gin.H{"err_code": 0, "err_msg": "点赞成功"})
	return

}

// DoComment 评论
func DoComment(c *gin.Context) {

	bID, _ := strconv.Atoi(c.PostForm("id"))
	pID, _ := strconv.Atoi(c.PostForm("pid"))
	userID, _ := strconv.Atoi(c.PostForm("user_id"))
	content := c.PostForm("content")

	var bc models.BroadcastComment

	bc.BroadcastID = bID
	bc.UserID = userID
	bc.PID = pID
	bc.CreateDate = time.Now()
	bc.Content = content
	if err := bc.Create(); err != nil {
		c.JSON(http.StatusOK, gin.H{"err_code": 1, "err_msg": err.Error()})
		return
	}
	// 评论数+1
	var b models.Broadcast
	b.Incr(bID, "comment_count")

	c.JSON(http.StatusOK, gin.H{"err_code": 0, "err_msg": "评论成功"})
	return
}

// DoReward 领取奖励
func DoReward(c *gin.Context) {
	bID, _ := strconv.Atoi(c.PostForm("id"))
	userID, _ := strconv.Atoi(c.PostForm("user_id"))
	prove := c.PostForm("prove")
	//todo:缓存获取可以获取的奖励
	integral := 1

	var br models.BroadcastReward
	br.UserID = userID
	br.BroadcastID = bID
	br.Integral = integral
	br.Prove = prove
	if err := br.Create(); err != nil {
		c.JSON(http.StatusOK, gin.H{"err_code": 1, "err_msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"err_code": 0, "err_msg": "领取" + string(integral) + "积分"})
	return

}
