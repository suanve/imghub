package Controllers

import (
	"imgHub/Services"
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
	re, _ := regexp.Compile(`^[a-z0-9]{32}$`)
	if re.MatchString(imgService.IMGID) {
		path, err := imgService.GetPath(imgService.IMGID)
		if err != nil {
			c.JSON(403, gin.H{"message": "img not found"})
			return
		}
		c.File(path)
		if !imgService.AddViewLOG() {
			// 报错日志不对外输出
			// c.JSON(403, gin.H{"message": "addlog error"})
		}

	} else {
		c.JSON(403, gin.H{"message": "img id not found"})
	}
}
