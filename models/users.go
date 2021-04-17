package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Users struct{
	Id int64 `json:"id"`
	Phone string `json:"phone"`
	Password string `json:"password"`
	CreatedAt int64 `json:"create_time" gorm:"default:null"`
	DeletedAt *time.Time `json:"-" gorm:"default:null"`
	Balance float64 `json:"balance" gorm:"default:0"`
	Status uint `json:"status" gorm:"default:1"`
	Code string `json:"code"`
	Pid int64 `json:"pid"`
	Ppid int64 `json:"ppid"`
	LeaderId int64 `json:"leader"`
	IsTest uint `json:"is_test"`
	LoginTime int64 `json:"login_time"`
}

func (this Users)BeforeCreate(scope *gorm.Scope)error{
	//创建时间
	scope.SetColumn("CreatedAt", time.Now().Unix())
	scope.SetColumn("LoginTime",time.Now().Unix())
	return nil
}
