package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func InitMySQL()(err error){
	//dsn := "root:poi$5574@tcp(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:123456@tcp(192.168.199.145:3307)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	DB.SingularTable(true)
	return DB.DB().Ping()
}

func Close(){
	DB.Close()
}