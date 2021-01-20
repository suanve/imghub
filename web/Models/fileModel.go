package Models

import (
	"fmt"
	Mysql "imgHub/web/Databases"
)

type Imgs struct {
	ID       int `gorm:"primary_key`
	IMGID    string
	FilePath string
	OldName  string
	ViewNum  int
	AddTime  int64
	AddIP    string
	AddUA    string
}

func init() {
	Mysql.DB.AutoMigrate(&Imgs{})
}

// 插入数据
func (this *Imgs) Insert() (id int, err error) {
	result := Mysql.DB.Create(&this)
	id = this.ID
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

// 根据传入内容 返回查询结果
func (this *Imgs) Where(where, value string) (imgs []Imgs, err error) {
	result := Mysql.DB.Where(fmt.Sprintf("%s = ?", where), value).Find(&imgs)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

// 更新内容
func (this *Imgs) Update(column string, value int) (err error) {
	result := Mysql.DB.Model(&this).Update(column, value)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

// 获取所有行数
func (this *Imgs) RowsAffected() int64 {
	result := Mysql.DB.Find(&this).RowsAffected
	return result
}

// 获取所有数据
func (this *Imgs) Rows() []Imgs {
	var imgs []Imgs
	Mysql.DB.Find(&imgs)
	return imgs
}

func (this *Imgs) Delete(where, value string) (flag bool) {
	flag = false
	result := Mysql.DB.Where(fmt.Sprintf("%s = ?", where), value).Delete(&this)
	if result.Error != nil {
		flag = false
		return
	}
	flag = true
	return
}
