package controller

import (
	"go-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context)  {
	userInfo,err := models.QueryUser()
	if err!= nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}else {
		c.JSON(http.StatusOK, userInfo)
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