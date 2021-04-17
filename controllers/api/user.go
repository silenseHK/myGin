package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"silence/services/api"
	"silence/tools"
	"silence/validators"
)

type User struct{
	Core
}

func init(){

}

var (
	validate  = new(validators.Valid)
	userService = new(api.UserService)
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
	userService.Register(registerRule)
	c.JSON(http.StatusOK,userService.GetReturn())
}

func (this User)GetVerify(c *gin.Context){
	defer this.ReceiveErr(c)
	var verifyRule validators.UserVerify
	if !validate.Validate(c, &verifyRule){
		return
	}
	userService.GetVerify(verifyRule)
	c.JSON(http.StatusOK,userService.GetReturn())
}

func (this User)Login(c *gin.Context){
	defer this.ReceiveErr(c)
	var verifyRule validators.UserLogin
	if !validate.Validate(c, &verifyRule){
		return
	}
	userService.Login(verifyRule)
	c.JSON(http.StatusOK,userService.GetReturn())
}

func (this User)Test(c *gin.Context){
	verify := tools.CreateVerify(6)
	c.JSON(http.StatusOK,verify)
}
