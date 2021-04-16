package routes

import (
	"silence/controllers/api"
)

type ApiRouter struct{

}

var (
	apiIndexCon = new(api.Index)
	apiGameCon = new(api.Game)
	apiUserCon = new(api.User)
)

func (this ApiRouter)Register(){
	apiGroup := r.Group("/api")

	apiGroup.POST("/register",apiUserCon.Register)
}
