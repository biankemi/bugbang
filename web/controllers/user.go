package controllers

import (
	"net/http"
	"time"

	"github.com/biankemi/bugbang/models"
	"github.com/biankemi/bugbang/services"
	"github.com/gin-gonic/gin"
)

// SendCode 发送验证码
func SendCode(c *gin.Context) {
	tel := c.PostForm("tel")
	isSendCode := services.SendCode(tel)

	if isSendCode == true {
		c.SecureJSON(http.StatusOK, "发送成功")
	}
	c.SecureJSON(http.StatusOK, "发送失败")
}

// Login 登录
func Login(c *gin.Context) {
	tel := c.PostForm("tel")
	code := c.PostForm("code")
	oldCode := "1234"
	isVerify := services.VerifyCode(oldCode, code)
	if isVerify == false {
		c.SecureJSON(http.StatusOK, "验证码校验失败")
	}
	user := models.User{Name: tel, Tel: tel, RegisterDate: time.Now()}
	error := user.Login()
	if error != nil {
		c.SecureJSON(http.StatusOK, error)
	}
	c.SecureJSON(http.StatusOK, "登录成功")
}
