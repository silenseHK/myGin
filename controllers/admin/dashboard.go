package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Dashboard struct{
	Core
}

func (this Dashboard)Index(c *gin.Context){
	c.HTML(http.StatusOK,"dashboard/dashboard.html","")
}
