package myRedis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"silence/tools"
)

var MyRedis redis.Conn

const (
	REGISTER = "HANG:REGISTER:"
)

func init(){
	databaseConf := tools.InitConfig("configs/db.yaml")
	if databaseConf == nil{
		fmt.Println("redis配置获取失败1")
		return
	}
	rConn,err := redis.Dial("tcp",databaseConf["redis_host"] + ":" + databaseConf["redis_port"])
	if err != nil{
		fmt.Println("redis链接失败")
		return
	}
	MyRedis = rConn
	if _,err = MyRedis.Do("AUTH",databaseConf["redis_password"]); err != nil{
		fmt.Println(err)
		return
	}
}


