package controllers

type Controller interface {


}

type Rtn struct{
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}