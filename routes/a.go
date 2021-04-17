package routes

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

var routerMap []Core

func init(){
	fmt.Println("a.go")
	r = gin.Default()
	store := cookie.NewStore([]byte("hangUp"))
	r.Use(sessions.Sessions("hangUpSession",store))
	r.LoadHTMLGlob("tem/*/*/*")
	r.Static("/asset", "./asset")
	register()
	r.Run(":8090")

}

func register(){
	addModule(ApiRouter{})
	addModule(AdminRouter{})
	run()
}

func addModule(module Core){
	routerMap = append(routerMap,module)
}

func run(){
	for _,module := range routerMap{
		module.Register()
	}
}
