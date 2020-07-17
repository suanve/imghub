package Controllers

import (
	"imgHub/web/Services"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
)

// ReadImg 获取图片
func ReadImg(c *gin.Context) {
	var imgService Services.Imgs

	imgService.IMGID = c.Param("IMGID")
	imgService.AddIP = c.ClientIP()              //获取访客ip
	imgService.AddTime = time.Now().Unix()       //时间
	imgService.AddUA = c.GetHeader("User-Agent") //ua
	re, _ := regexp.Compile(`[a-z0-9]{32}`)
	if re.MatchString(imgService.IMGID) {
		path, err := imgService.GetPath(imgService.IMGID)
		if err != nil {
			c.JSON(500, gin.H{"message": "server err"})
			return
		}
		c.File(path)
		if !imgService.AddViewLOG() {
			c.JSON(401, gin.H{"message": "not found"})
		}

	} else {
		c.JSON(401, gin.H{"message": "not found"})
	}
}
