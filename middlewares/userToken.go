package middlewares

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"silence/myRedis"
	"silence/tools"
	"strconv"
	"strings"
)

type UserTokenMiddleware struct {
	BaseMiddleWare
}

func (this UserTokenMiddleware)CheckToken(c *gin.Context){
	header := c.Request.Header
	token,ok := header["Token"]
	if !ok {
		c.Abort()
		this.ReturnJsonWithErr(c,401,"please login again")
		return
	}
	fmt.Printf("token:%#v \n",token)
	tokenStr := fmt.Sprintf(strings.Join(token,""))
	userId,err := tools.DecryptToken(tokenStr)
	if err != nil{
		c.Abort()
		this.ReturnJsonWithErr(c,401,err.Error())
		return
	}
	//验证token与redis上的token是否相同
	tokenRedis,err := redis.String(myRedis.MyRedis.Do("Get", myRedis.USER_TOKEN+strconv.FormatInt(userId,10)))
	if err != nil{
		c.Abort()
		this.ReturnJsonWithErr(c,401,"please login again")
		return
	}
	if tokenRedis != tokenStr{
		c.Abort()
		this.ReturnJsonWithErr(c,401,"please login again")
		return
	}
	c.Set("userId",userId)
}
