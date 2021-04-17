package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"silence/services/api"
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
