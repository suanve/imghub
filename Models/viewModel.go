package Models

import Mysql "imgHub/Databases"

type ViewLogs struct {
	ID    int
	IMGID string
	TIME  int64
	IP    string
	UA    string
}

func init() {
	Mysql.DB.AutoMigrate(&ViewLogs{})
}

// 添加访问日志
func (this *ViewLogs) Insert() (id int, err error) {
	result := Mysql.DB.Create(&this)
	id = this.ID
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}
