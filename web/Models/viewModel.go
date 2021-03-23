package Models

import (
	Sqlite "imgHub/web/Databases"
)

type ViewLogs struct {
	ID    int
	IMGID string
	TIME  int64
	IP    string
	UA    string
}

func init() {
	Sqlite.DB.AutoMigrate(&ViewLogs{})
}

// 添加访问日志
func (this *ViewLogs) Insert() (id int, err error) {
	result := Sqlite.DB.Create(&this)
	id = this.ID
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}
