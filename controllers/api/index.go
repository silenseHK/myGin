package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Index struct{
	Core
}

func (con Index)Default(c *gin.Context){
	c.JSON(http.StatusOK,rtn)
	return
}
