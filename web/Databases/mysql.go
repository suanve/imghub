package Mysql

import (
	"fmt"
	"imgHub/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func init()  {
	var err error
	DB, err = gorm.Open("mysql", config.DB)
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}
	if DB.Error != nil {
		fmt.Printf("database error %v", DB.Error)
	}
}
