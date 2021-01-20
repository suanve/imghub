package Services

import (
	"imgHub/web/Models"
)

// var Imgs Models.Imgs

// 添加图片

type ManageImg struct {
	IMGID string `json:"img_id"`
}

type ManageUser struct {
	Username string `json:username`
	Password string `json:password`
}

func (this *ManageImg) Delete() (flag bool) {
	var img Models.Imgs
	return img.Delete("img_id", this.IMGID)
}

func (this *ManageUser) Check() bool {
	var user Models.Users
	user.Username = this.Username
	user.Password = this.Password
	user = user.Check()
	if user.Id > 0 {
		return true
	}
	return false
}
