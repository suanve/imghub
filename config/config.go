package config

import (
	"fmt"
	"os"
	"strconv"
)

var (
	RootDir   = "./web"
	SecretKey string
	DB        string
	HubHost   = "127.0.0.1"
	Port      = 8081
	MysqlHost = "127.0.0.1"
	MysqlPort = 3307
	MysqlPass = "root"
)

func init() {
	if os.Getenv("MysqlHost") != "" {
		MysqlHost = os.Getenv("MysqlHost")
	} else {
		MysqlHost = "127.0.0.1"
	}

	if os.Getenv("mysqlPass") != "" {
		MysqlPass = os.Getenv("mysqlPass")
	}
	if os.Getenv("mysqlPort") != "" {
		MysqlPort, _ = strconv.Atoi(os.Getenv("mysqlPort"))
	}

	DB = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8",
		"root",
		MysqlPass,
		MysqlHost,
		MysqlPort,
		"imgHub")
	SecretKey = "biubiubiu"
	if os.Getenv("HubHost") != "" {
		HubHost = os.Getenv("HubHost")
	}
}
