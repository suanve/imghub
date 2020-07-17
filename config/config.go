package config

import (
	"fmt"
	"os"
)

var (
	RootDir = "./web"
	SecretKey string
	DB string
	HubHost = "127.0.0.1"
)

func init() {
	var mysqlHost string
	DockerPass := "xingyu"
	if os.Getenv("MysqlHost") != "" {
		mysqlHost = os.Getenv("MysqlHost")
	} else {
		mysqlHost = "127.0.0.1"
	}

	if os.Getenv("DockerPass") != "" {
		DockerPass = os.Getenv("DockerPass")
	}

	DB = fmt.Sprintf("%v:%v@tcp(%v:3306)/%v?charset=utf8",
		"root",
		DockerPass,
		mysqlHost,
		"imgHub")
	SecretKey = "biubiubiu"
	if os.Getenv("HubHost") != "" {
		HubHost = os.Getenv("HubHost")
	} else {
		HubHost = "127.0.0.1"
	}
}
