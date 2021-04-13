package routes

import (
	"silence/controllers/api"
)

type ApiRouter struct{

}

var (
	apiIndexCon = new(api.Index)
	apiGameCon = new(api.Game)
)

func (this ApiRouter)Register(){
	g := r.Group("/api")

	g.GET("/index",apiIndexCon.Default)
	g.GET("/index2",apiGameCon.Default)
}
