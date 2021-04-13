package admin

import (
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
