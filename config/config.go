package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var (
	RootDir             = "./web"
	SecretKey           []byte
	DB                  string
	HubHost             = "127.0.0.1"
	Port                = 8081
	MysqlHost           = "127.0.0.1"
	MysqlUser           = "root"
	MysqlPass           = "root"
	MysqlPort           = 3307
	TokenExpireDuration = time.Hour * 24
)

func init() {
	if os.Getenv("mysqlHost") != "" {
		MysqlHost = os.Getenv("mysqlHost")
	}

	if os.Getenv("mysqlUser") != "" {
		MysqlUser = os.Getenv("mysqlUser")
	}

	if os.Getenv("mysqlPass") != "" {
		MysqlPass = os.Getenv("mysqlPass")
	}

	if os.Getenv("mysqlPort") != "" {
		MysqlPort, _ = strconv.Atoi(os.Getenv("mysqlPort"))
	}

	if os.Getenv("hubHost") != "" {
		HubHost = os.Getenv("hubHost")
	}

	DB = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8",
		MysqlUser,
		MysqlPass,
		MysqlHost,
		MysqlPort,
		"imgHub")
	SecretKey = []byte("夏天夏天悄悄过去")

}
