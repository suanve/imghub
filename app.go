package main

import (
	Mysql "imgHub/web/Databases"
	"imgHub/web/Router"
)

func main() {
	Router.InitRouter()
	defer Mysql.DB.Close()
}
