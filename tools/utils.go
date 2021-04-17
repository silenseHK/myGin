package tools

import (
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"strconv"
	"strings"
	"time"
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

func CreateVerify(n int)(verify string){
	var verifySlice []string
	rand.Seed(time.Now().Unix())
	for i:=n;i>0;i--{
		verifySlice = append(verifySlice,strconv.Itoa(rand.Intn(9)))
	}
	return strings.Join(verifySlice,"")
}

func GetDayStartUnix()int64{
	currentTime := time.Now()
	startTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location()).Unix()
	return startTime
}

func GetDayEndUnix()int64{
	currentTime := time.Now()
	endTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 23, 59, 59, 0, currentTime.Location()).Unix()
	return endTime
}
