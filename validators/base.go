package validators

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"net/http"
	"reflect"
	"silence/controllers"
	"strings"
)

type Valid struct {

}

// 定义一个全局翻译器T
var trans ut.Translator

func init(){
	if err := InitTrans("zh"); err != nil {
		return
	}
}

func InitTrans(locale string)(err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		//// 注册一个获取json tag的自定义方法
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		zhT := zh.New() // 中文翻译器
		enT := en.New() // 英文翻译器

		// 第一个参数是备用（fallback）的语言环境
		// 后面的参数是应该支持的语言环境（支持多个）
		// uni := ut.New(zhT, zhT) 也是可以的
		uni := ut.New(enT, zhT, enT)

		// locale 通常取决于 http 请求头的 'Accept-Language'
		var ok bool
		// 也可以使用 uni.FindTranslator(...) 传入多个locale进行查找
		trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
		}

		// 注册翻译器
		switch locale {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		}
		return
	}
	return
}

func removeTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}

func(this Valid)Validate(c *gin.Context, rule interface{})(ok bool){
	err := c.ShouldBind(rule)
	if err != nil{
		errs, ok := err.(validator.ValidationErrors)
		var msg string
		if !ok {
			msg = err.Error()
		}else{
			msg = getMessageFromMap(removeTopStruct(errs.Translate(trans)))
		}
		c.JSON(http.StatusOK,controllers.Rtn{
			Code: http.StatusForbidden,
			Msg: msg,
			Data: "",
		})
		return false
	}
	return true
}

func getMessageFromMap(errMap validator.ValidationErrorsTranslations)(msg string){
	for _,v := range errMap{
		msg = v
		break
	}
	return msg
}
