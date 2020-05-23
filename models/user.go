package models

import (
	"go-api/dao"
)

//类型要大写
type UserParamsReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//不能命名为UserId,如此严格
type User struct {
	Id        uint   `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Telephone string `json:"telephone"`
}
//单个信息
func QueryUser()(userInfo *User,err error)  {
	userInfo = new(User) //这里要新建对象
	if err = dao.DB.First(&userInfo).Error; err != nil{
		return nil, err
	}
	return
}
//保存用户信息
func SaveUser(userRegister *UserParamsReq)(err error)  {
	user :=User{Username:userRegister.Username,Password:userRegister.Password,Telephone:""}
	err = dao.DB.Create(&user).Error
	return
}
//数组
func QueryUserList()(userList []*User,err error)  {
	// 因为是slice，猜测是包含新建的意思，所以不用像新建对象一样新建
	if err = dao.DB.Find(&userList).Error; err != nil{
		return nil, err
	}
	return
}