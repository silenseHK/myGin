package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"html/template"
	"net/http"
	"silence/controllers"
)

type Core struct{
	controllers.Controller
}

var rtn controllers.Rtn
var TempPath = "tem/admin/"

func init(){
	rtn.Data = ""
	rtn.Code = http.StatusOK
	rtn.Msg = "success"

}

func (this Core)HtmlExec(writer gin.ResponseWriter, temp string, name string)error{
	T,_ := template.New("").Delims("<<",">>").ParseFiles(
		"tem/admin/public/menu.html",
		"tem/admin/public/header.html",
		"tem/admin/public/footer.html",
		"tem/admin/public/layout.html",
	)
	t,err := T.ParseFiles(TempPath + temp)
	if err != nil{
		return err
	}
	err = t.ExecuteTemplate(writer,name,nil)
	return err
}

func (this Core)ReceiveErr(c *gin.Context){
	err := recover()
	msg := cast.ToString(err)
	if msg != ""{
		rtn.Code = http.StatusForbidden
		rtn.Msg = msg
		c.JSON(http.StatusOK,rtn)
	}
}
