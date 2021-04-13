package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct{
	Core
}

func (this Login)Index(c *gin.Context){
	c.HTML(http.StatusOK,"login.html","")
}

func (this Login)Login(c *gin.Context){
	c.JSON(http.StatusOK,rtn)
}
