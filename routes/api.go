package routes

import (
	"silence/controllers/api"
	"silence/middlewares"
)

type ApiRouter struct{

}

var (
	apiIndexCon = new(api.Index)
	apiGameCon = new(api.Game)
	apiUserCon = new(api.User)
	apiRecharge = new(api.Recharge)

	middleUserToken = new(middlewares.UserTokenMiddleware)
)

func (this ApiRouter)Register(){
	apiGroup := r.Group("/api")

	apiGroup.POST("/register",apiUserCon.Register)
	apiGroup.GET("/verify",apiUserCon.GetVerify)
	apiGroup.GET("/test",apiUserCon.Test)
	apiGroup.POST("/login",apiUserCon.Login)

	registerGroup := apiGroup.Group("/recharge", middleUserToken.CheckToken)
	registerGroup.GET("/type",apiRecharge.TypeList)
}
