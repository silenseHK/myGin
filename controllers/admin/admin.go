package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"silence/models"
)

type Admin struct{
	Core
}

func (this Admin)AccountList(c *gin.Context){
	defer this.ReceiveErr(c)
	err := this.HtmlExec(c.Writer,"admin/account.html","account.html")
	if err != nil{
		panic(err)
	}
}

func (this Admin)AccountAdd(c *gin.Context){
	defer this.ReceiveErr(c)
	account := c.PostForm("account")
	if account == ""{
		panic("缺少参数:account")
		return
	}
	password := c.PostForm("password")
	if password == ""{
		panic("缺少参数:password")
		return
	}
	pwd,err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err != nil{
		panic("加密失败:" + err.Error())
		return
	}
	fmt.Println(string(pwd))
	data := &models.Admin{
		Account: account,
		Password: string(pwd),
	}
	new(models.Models).Insert(data)
	c.JSON(http.StatusOK,rtn)
}


