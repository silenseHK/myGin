package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"silence/services/api"
	"silence/validators"
)

type Recharge struct{
	Core
}

var (
	rechargeService *api.RechargeService
)

func init(){
	rechargeService = new(api.RechargeService)
}

func (this Recharge)TypeList(c *gin.Context){
	rechargeService.GetTypeList()
	c.JSON(http.StatusOK,rechargeService.GetReturn())
}

func (this Recharge)Recharge(c *gin.Context){
	defer this.ReceiveErr(c)
	var verifyRule validators.RechargeRecharge
	if !validate.Validate(c, &verifyRule){
		return
	}
	rechargeService.Recharge(verifyRule)
	c.JSON(http.StatusOK,userService.GetReturn())
}
