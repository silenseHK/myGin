package api

import (
	"github.com/garyburd/redigo/redis"
	"silence/myRedis"
	"silence/repository/api"
	"silence/tools"
	"silence/validators"
	"strconv"
	"time"
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
	verify,err := redis.String(myRedis.MyRedis.Do("Get", myRedis.REGISTER+valid.Phone))
	if err != nil{
		this.rtnWithError("Phone verification code error")
		return
	}
	if verify != valid.Verify{
		this.rtnWithError("Phone verification code error.")
		return
	}
	//注册用户
	registerUser := userRep.Register(valid)
	if registerUser.Id == 0{
		this.rtnWithError("Registration failed")
		return
	}
	//添加邀请码
	//邀请码
	code := tools.TestInviteCode2(1)
	registerUser.Code = code
	userRep.SaveUser(registerUser)

	Data = registerUser
	return
}

func (this UserService)GetVerify(valid validators.UserVerify){
	//判断验证码类型
	_,err := redis.String(myRedis.MyRedis.Do("Get", myRedis.REGISTER+valid.Phone))
	if err == nil{
		this.rtnWithError("Do not send SMS repeatedly within the validity period of the verification code")
		return
	}
	n,_ := redis.Int(myRedis.MyRedis.Do("Get", myRedis.REGISTER_COUNT+valid.Phone))
	if  n >= 5{
		this.rtnWithError("You can only get the verification code 5 times a day")
		return
	}
	//生成验证码
	verify := tools.CreateVerify(6)
	//存储redis
	myRedis.MyRedis.Do("Set", myRedis.REGISTER+valid.Phone,verify,"EX",1800)
	myRedis.MyRedis.Do("Set", myRedis.REGISTER_COUNT+valid.Phone,n+1,"EX",tools.GetDayEndUnix()-time.Now().Unix())
	Data = map[string]interface{}{
		verify: verify,
	}
	return
}

func (this UserService)Login(valid validators.UserLogin){
	//检查用户是否存在
	user := userRep.GetUserByPhone(valid.Phone)
	if user.Id == 0{
		this.rtnWithError("Incorrect username or password")
		return
	}
	if !tools.CheckPwd(user.Password,[]byte(valid.Password)){
		this.rtnWithError("Incorrect username or password.")
		return
	}
	if user.Status == 0{
		this.rtnWithError("Account has been frozen")
		return
	}
	user.LoginTime = time.Now().Unix()
	userRep.SaveUser(user)
	user.Password = ""
	token := this.doLogin(user.Id)
	Data = map[string]interface{}{
		"token": token,
		"userInfo": user,
	}
	return
}

func (this UserService)doLogin(userId int64)(token string){
	token = tools.EncryptToken(userId)
	expireTime := 7 * 24 * 60 * 60
	myRedis.MyRedis.Do("Set", myRedis.USER_TOKEN+strconv.FormatInt(userId,10),token,"EX",expireTime)
	return
}