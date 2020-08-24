package controller

import (
	"github.com/gin-gonic/gin"
	"go-api/models"
	"net/http"
)

type HeaderType struct {
	RetCode string `json:"retCode"`
	ErrorNum string `json:"errorNum"`
	ErrorMsg string `json:"errorMsg"`
}
type BodyTypeSub struct {
	Message string `json:"message"`
}

type responseType struct {
	Header HeaderType `json:"header"`
	Body BodyTypeSub `json:"body"`
}
type ResponseTypeNoBody struct {
	Header HeaderType `json:"header"`
}

func Login(c *gin.Context) {
	var userParams models.UserParamsReq
	c.BindJSON(&userParams)
	userInfo, err := models.QueryUser(userParams.Username)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	if userParams.Username == userInfo.Username && userParams.Password == userInfo.Password {
		c.JSON(http.StatusOK, gin.H{"message": "登录成功"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "账号或密码错误"})
	}
}

func Register(c *gin.Context) {
	var userParams models.UserParamsReq
	c.BindJSON(&userParams)
	_, err := models.QueryUser(userParams.Username)

	var response ResponseTypeNoBody
	response.Header.RetCode = "0"
	if err != nil {
		models.SaveUser(&userParams)
		response.Header.ErrorNum = "0"
		response.Header.ErrorMsg = "注册成功"
		c.JSON(http.StatusOK, response)
		return
	}
	response.Header.ErrorNum = "1"
	response.Header.ErrorMsg = "账号已存在"
	c.JSON(http.StatusOK, response)
}

func QueryUserList(c *gin.Context) {
	userList, err := models.QueryUserList()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, userList)
	}
}
