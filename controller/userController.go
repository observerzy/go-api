package controller

import (
	"fmt"
	"go-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)
type UserParamsReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Telephone string `json:"telephone"`
}
func Login(c *gin.Context)  {
	userInfo,err := models.QueryUser()
	if err!= nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}else {
		c.JSON(http.StatusOK, userInfo)
	}
}

func Register(c *gin.Context)  {
	var userParams UserParamsReq
	userInfo,err := models.QueryUser()
	c.BindJSON(&userParams)
	fmt.Println(userInfo)
	fmt.Println(userParams)
	if err!= nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	if userParams.Username==userInfo.Username&& userParams.Password==userInfo.Password&&userParams.Telephone==userInfo.Telephone {
		c.JSON(http.StatusOK, gin.H{"message":"注册成功"})
	}else {
		c.JSON(http.StatusOK, gin.H{"message":"账号或密码错误"})
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