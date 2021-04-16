package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

var routerMap []Core

func init(){
	fmt.Println("a.go")
	r = gin.Default()
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
