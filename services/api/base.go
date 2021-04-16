package api

import (
	"net/http"
	"silence/controllers"
)

type Service struct{

}

var (
	Code = http.StatusOK
	Msg = "success"
	Data interface{}
)

func (this Service)rtnWithError(err string){
	Code = http.StatusForbidden
	Msg = err
}

func (this Service)GetReturn()controllers.Rtn{
	var rtn = controllers.Rtn{Code: Code,Msg: Msg,Data: Data}
	Init()
	return rtn
}

func Init(){
	Code = http.StatusOK
	Msg = "success"
	Data = ""
}
