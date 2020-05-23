package controller

import (
	"github.com/gin-gonic/gin"
	"go-api/models"
	"net/http"
)

func Login(c *gin.Context)  {
	var userParams models.UserParamsReq
	c.BindJSON(&userParams)
	userInfo,err := models.QueryUser()
	if err!= nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	if userParams.Username==userInfo.Username&& userParams.Password==userInfo.Password {
		c.JSON(http.StatusOK, gin.H{"message":"登录成功"})
	}else {
		c.JSON(http.StatusOK, gin.H{"message":"账号或密码错误"})
	}
}

func Register(c *gin.Context)  {
	var userParams models.UserParamsReq
	c.BindJSON(&userParams)
	userInfo,err := models.QueryUser()
	if err!= nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	if userParams.Username==userInfo.Username{
		c.JSON(http.StatusOK, gin.H{"message":"账号已存在"})
	}else {
		models.SaveUser(&userParams)
		c.JSON(http.StatusOK, gin.H{"message":"注册成功"})
	}
}

func QueryUserList(c *gin.Context)  {
	userList,err := models.QueryUserList()
	if err!= nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}else {
		c.JSON(http.StatusOK, userList)
	}
}