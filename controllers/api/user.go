package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"silence/services/api"
	"silence/validators"
)

type User struct{
	Core
}

func init(){
	userService = new(api.UserService)
}

var (
	validate  = new(validators.Valid)
	userService *api.UserService
)

func (this User)Get(c *gin.Context){
	c.JSON(http.StatusOK, rtn)
}


func (this User)Register(c *gin.Context){
	defer this.ReceiveErr(c)
	var registerRule validators.UserRegister
	if !validate.Validate(c, &registerRule){
		return
	}
	fmt.Println(userService.GetReturn())
	userService.Register(registerRule)
	c.JSON(http.StatusOK,userService.GetReturn())
}
