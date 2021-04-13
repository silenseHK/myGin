package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Game struct{
	Core
}

func (con Game)Default(c *gin.Context){
	rtn.Msg = "error"
	c.JSON(http.StatusOK,rtn)
}