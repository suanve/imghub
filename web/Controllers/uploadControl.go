package Controllers

import (
	"fmt"
	"imgHub/config"
	"imgHub/web/Common"
	"imgHub/web/Services"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// Upload 文件上传
func Upload(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file[]"]

	url := ""
	for _, file := range files {

		ext := strings.Split(file.Filename, ".")[len(strings.Split(file.Filename, "."))-1]
		var Img Services.Imgs
		Img.AddIP = c.ClientIP()              //获取访客ip
		Img.AddTime = time.Now().Unix()       //时间
		Img.AddUA = c.GetHeader("User-Agent") //ua
		Img.FileName = fmt.Sprintf("%s_%d.%s", Common.RandomString(5), Img.AddTime, ext)
		Img.FilePath = "./uploads/" + Img.FileName
		Img.OldName = file.Filename
		c.SaveUploadedFile(file, Img.FilePath)

		Img.IMGID = Common.Md5(file.Filename + strconv.FormatInt(Img.AddTime, 10) + Common.RandomString(20))

		fileurl := fmt.Sprintf("%s://%s:%d/uploads/%s", config.Protocol, config.HubHost, config.Port, Img.IMGID)
		url += fileurl + "\n"
		Img.AddImage()
	}
	c.String(http.StatusOK, fmt.Sprintf("Upload Success:\n%s", url))
}
