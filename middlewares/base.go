package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"silence/controllers"
)

type BaseMiddleWare struct{

}

func (this BaseMiddleWare)ReturnJsonWithErr(c *gin.Context, code int, msg string){
	var rtnCode int
	if code > 0{
		rtnCode = code
	}else{
		rtnCode = http.StatusForbidden
	}
	var rtn = controllers.Rtn{
		Code: rtnCode,
		Msg: msg,
		Data: "",
	}
	c.JSON(http.StatusOK,rtn)
}
