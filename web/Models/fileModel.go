package Models

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"imgHub/web/Databases"
)

//用户数据结构
type User struct {
	Id       int    `json:id`
	Username string `json:username`
	Password string `json:password`
	Level    int    `json:level`
}

// JWT认证数据结构
type Claims struct {
	Id       int    `json:"id"`       // id
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
	Level    int    `json:"level"`    // 密码
	jwt.StandardClaims
}

type Imgs struct {
	ID	int `gorm:"primary_key`
	IMGID    string
	FilePath string
	OldName  string
	ViewNum  int
	AddTime  int64
	AddIP    string
	AddUA    string
}



func init(){
	Mysql.DB.AutoMigrate(&Imgs{})
}


// 添加文件
func (this *Imgs) Insert() (id int,err error){
	result := Mysql.DB.Create(&this)
	id = this.ID
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

// 根据传入内容 返回查询结果
func (this *Imgs) Where(where ,value string) (img *Imgs,err error){
	result := Mysql.DB.Where(fmt.Sprintf("%s = ?",where), value).First(&this)
	img = this
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

// 更新内容
func (this *Imgs) Update(column string,value int) (err error){
	result := Mysql.DB.Model(&this).Update(column, value)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}