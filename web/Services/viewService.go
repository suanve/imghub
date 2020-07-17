package Services

import (
	"imgHub/web/Models"
)

type ViewLogs struct {
	ID    int
	IMGID string
	TIME  int64
	IP    string
	UA    string
}

// 访问量增加
func (this *ViewLogs) AddViewLOG() (id int,err error) {
	var ViewLogsModel Models.ViewLogs
	ViewLogsModel.IMGID = this.IMGID
	ViewLogsModel.IP = this.IP
	ViewLogsModel.TIME = this.TIME
	ViewLogsModel.UA = this.UA
	id, err = ViewLogsModel.Insert()
	return
}
