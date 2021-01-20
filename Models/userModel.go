package Models

import (
	"fmt"
	Mysql "imgHub/Databases"
)

type Users struct {
	Id       int    `json:"id"`       // id
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
}

func init() {
	Mysql.DB.AutoMigrate(&Users{})
}

func (this *Users) Where(where, value string) (user []Users) {
	result := Mysql.DB.Where(fmt.Sprintf("%s = ?", where), value).Find(&user)
	if result.Error != nil {
		return
	}
	return
}

func (this *Users) Check() (user Users) {
	result := Mysql.DB.Where("username = ? AND password = ?", this.Username, this.Password).Find(&user)
	if result.Error != nil {
		return
	}
	return
}
