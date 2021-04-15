package main

import (
	"silence/models"
	_ "silence/models"
	_ "silence/routes"
)

func main(){
	defer models.DB.Close()
}
