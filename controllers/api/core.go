package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
	"silence/controllers"
)

type Core struct{
	controllers.Controller
}

var rtn controllers.Rtn

func init(){
	rtn.Data = ""
	rtn.Code = http.StatusOK
	rtn.Msg = "success"
}

func (this Core)ReceiveErr(c *gin.Context){
	err := recover()
	if err != nil{
		fmt.Println(err)
	}
	msg := cast.ToString(err)
	if msg != ""{
		rtn.Code = http.StatusForbidden
		rtn.Msg = msg
		c.JSON(http.StatusOK,rtn)
	}
}
