package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Admin struct{
	Core
}

func (this Admin)AccountList(c *gin.Context){
	c.HTML(http.StatusOK,"admin/account.html","")
}
