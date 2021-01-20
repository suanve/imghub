package main

import (
	Mysql "imgHub/Databases"
	"imgHub/Router"
)

func main() {
	Router.InitRouter()
	defer Mysql.DB.Close()
}
