package main

import (
	"silence/models"
	"silence/myRedis"
	_ "silence/myRedis"
	_ "silence/routes"
	_ "silence/validators"
)

func main(){
	defer models.DB.Close()
	defer myRedis.MyRedis.Close()
}
