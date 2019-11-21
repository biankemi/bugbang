package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/biankemi/bugbang/models"
	"github.com/biankemi/bugbang/services"
	"github.com/gin-gonic/gin"
)

// SendCode 发送验证码
func SendCode(c *gin.Context) {
	tel := c.Query("tel")
	isSendCode := services.SendCode(tel)

	if isSendCode == true {
		c.SecureJSON(http.StatusOK, "发送成功")
		return
	}
	c.SecureJSON(http.StatusOK, "发送失败")
	return
}

// Login 登录
func Login(c *gin.Context) {
	tel := c.Query("tel")
	code := c.Query("code")
	oldCode := "1234"
	isVerify := services.VerifyCode(oldCode, code)
	if isVerify == false {
		c.SecureJSON(http.StatusOK, "验证码校验失败")
		return
	}
	user := models.User{Name: tel, Tel: tel, RegisterDate: time.Now()}
	error := user.Login()
	if error != nil {
		c.SecureJSON(http.StatusOK, error)
		return
	}
	c.SetCookie("user_tel", tel, 3600, "/", "127.0.0.1", false, true)
	c.JSON(http.StatusOK, gin.H{"msg": "success"})
	return
}

// UserEdit 用户修改
func UserEdit(c *gin.Context) {

	var u models.User
	ID, _ := strconv.Atoi(c.Param("id"))
	u.Avator = c.PostForm("avator")
	u.Name = c.PostForm("name")
	u.Tel = c.PostForm("tel")
	u.Integral, _ = strconv.Atoi(c.PostForm("integral"))
	u.CompanyAddress = c.PostForm("company_address")
	u.HomeAddress = c.PostForm("home_address")
	if err := u.Edit(ID); err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
	return
}
