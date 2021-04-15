package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Dashboard struct{
	Core
}

func (this Dashboard)Index(c *gin.Context){
	defer func(){
		err := recover()
		fmt.Println("err: ",err)
	}()
	err := this.HtmlExec(c.Writer,"dashboard/dashboard.html","dashboard.html")
	if err != nil{
		panic(err)
	}
}
