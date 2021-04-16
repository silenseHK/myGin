package tools

import (
	"golang.org/x/crypto/bcrypt"
)

func EncryptPwd(pwd string)(password string, error error){
	hash,err := bcrypt.GenerateFromPassword([]byte(pwd),bcrypt.MinCost)
	if err != nil{
		return "", nil
	}
	return string(hash),nil
}

func CheckPwd(hashedPwd string, plainPwd []byte)bool{
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return false
	}
	return true
}
