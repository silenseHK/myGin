package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Models struct{

}

var DB *gorm.DB

func init(){
	db,err := gorm.Open("mysql", "root:root@(127.0.0.1:3306)/gin?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		panic("mysql conn err :" + err.Error())
	}
	db.SingularTable(true)   //设置全局表名禁用复数
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string)string{
		return "gin_" + defaultTableName
	}
	DB = db
}

func (this Models)Insert(data interface{}){
	DB.Create(data)
}
