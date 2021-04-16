package api

import (
	"github.com/garyburd/redigo/redis"
	"silence/myRedis"
	"silence/repository/api"
	"silence/validators"
)

type UserService struct{
	Service
}

func init(){
	userRep = new(api.UserRep)
}

var (
	userRep *api.UserRep
)

func (this UserService)Register(valid validators.UserRegister){
	user := userRep.GetUserByPhone(valid.Phone)
	if user.Id != 0{
		this.rtnWithError("The phone number has been registered, please log in directly")
		return
	}
	//验证码
	//myRedis.MyRedis.Do("Set", myRedis.REGISTER+valid.Phone,666666)
	code,err := redis.String(myRedis.MyRedis.Do("Get", myRedis.REGISTER+valid.Phone))
	if err != nil{
		this.rtnWithError("Phone verification code error")
		return
	}
	if code != valid.Verify{
		this.rtnWithError("Phone verification code error.")
		return
	}
	//注册用户
	registerUser := userRep.Register(valid)
	if registerUser.Id == 0{
		this.rtnWithError("Registration failed")
		return
	}
	return
}
