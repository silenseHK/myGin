package models

type Admin struct {
	Id int64 `mysql:"id"`
	Account string `mysql:"account"`
	Password string `mysql:"password"`
}
