package api

import (
	"silence/models"
	"silence/tools"
	"silence/validators"
)

type UserRep struct{

}

func (this UserRep)GetUserByPhone(phone string)(user models.Users){
	models.DB.Where(&models.Users{Phone: phone}).First(&user)
	return
}

func (this UserRep)Register(valid validators.UserRegister)models.Users{
	password,_ := tools.EncryptPwd(valid.Password)
	var user = models.Users{
		Phone: valid.Phone,
		Password: password,
	}
	new(models.Models).Insert(&user)
	return user
}

func (this UserRep)SaveUser(userData models.Users){
	models.DB.Save(&userData)
}
