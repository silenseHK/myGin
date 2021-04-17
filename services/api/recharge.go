package api

import (
	"silence/repository/api"
	"silence/validators"
)

type RechargeService struct{
	Service
}

var (
	confRep *api.ConfRep
)

func init(){
	confRep = new(api.ConfRep)
}

func (this RechargeService)GetTypeList(){
	types,err := confRep.GetRechargeConf()
	if err != nil{
		this.rtnWithError(err.Error())
		return
	}
	Data = map[string]interface{}{
		"list": types,
	}
	return
}

func (this RechargeService)Recharge(valid validators.RechargeRecharge){

}