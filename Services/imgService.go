package Services

import (
	"imgHub/Models"
)

type Imgs struct {
	IMGID    string `json:id`
	FileName string `json:filename`
	FilePath string `json:filepath`
	OldName  string `json:oldname`
	ViewNum  int    `json:viewnum`
	AddTime  int64  `json:addtime`
	AddIP    string `json:addip`
	AddUA    string `json:addua`
}

// 添加图片
func (this *Imgs) AddImage() (id int, err error) {
	var imgModel Models.Imgs
	imgModel.IMGID = this.IMGID
	imgModel.FilePath = this.FilePath
	imgModel.OldName = this.OldName
	imgModel.ViewNum = 0
	imgModel.AddTime = this.AddTime
	imgModel.AddIP = this.AddIP
	imgModel.AddUA = this.AddUA
	id, err = imgModel.Insert()
	return
}

// 获取图片存储地址
func (this *Imgs) GetPath(id string) (path string, err error) {
	var imgModel Models.Imgs
	img, err := imgModel.Where("img_id", id)
	if len(img) > 0 {
		path = img[0].FilePath
	}
	return
}

// 访问量增加
func (this *Imgs) AddViewLOG() (flag bool) {
	var imgModel Models.Imgs
	var ViewLogsModel Models.ViewLogs
	flag = true
	img, err := imgModel.Where("img_id", this.IMGID)
	if err != nil {
		flag = false
	}
	if len(img) > 0 {
		img[0].Update("view_num", img[0].ViewNum+1)
		// imgModel.Update()
	}

	ViewLogsModel.IMGID = this.IMGID
	ViewLogsModel.IP = this.AddIP
	ViewLogsModel.TIME = this.AddTime
	ViewLogsModel.UA = this.AddUA
	ViewLogsModel.Insert()
	return
}
