package Controllers

import (
	"encoding/json"
	"fmt"
	"imgHub/Models"
	"imgHub/Services"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ManageIndex(c *gin.Context) {
	c.JSON(200, gin.H{"message": "It's work!"})

}

func ManageLogin(c *gin.Context) {

	data, _ := ioutil.ReadAll(c.Request.Body)
	var user Services.ManageUser

	if err := json.Unmarshal(data, &user); err == nil {
		// fmt.Println(user.Username)
		// fmt.Println(user.Password)
	}
	if user.Username == "" || user.Password == "" {
		c.JSON(http.StatusOK, gin.H{"code": 401})
	}

	if user.Check() {
		var claims Models.Claims
		tokenString, err := claims.GenToken(user.Username)
		if err != nil {
			fmt.Println(err.Error())
		}
		// fmt.Println(user.Username)

		c.JSON(http.StatusOK, gin.H{"code": 200, "token": tokenString})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 401})
	}
}

// ManageGet 用来获取数据
func ManageGet(c *gin.Context) {
	action := c.Param("action")
	var imgModel Models.Imgs
	switch action {
	case "RowsAffected":
		c.JSON(200, gin.H{"RowsAffected": imgModel.RowsAffected()})
		return
	case "All":
		imgs := imgModel.Rows()
		c.JSON(200, gin.H{
			"message": "success",
			"total":   len(imgs),
			"data":    imgs,
		})
		return
	case "filter":
		filterstr := c.Query("filter")
		value := c.Query("value")
		if filterstr == "" || value == "" {
			filterstr = "id"
			value = "1"
		}
		imgs, err := imgModel.Where(filterstr, value)
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
		}
		c.JSON(200, gin.H{
			"message": "success",
			"total":   len(imgs),
			"data":    imgs,
		})
		return
	}
	c.JSON(200, gin.H{"message": "ok"})
}

// ManageDel 用来删除数据
func ManageDel(c *gin.Context) {
	action := c.Param("action")
	var img Services.ManageImg
	var tpl map[string]string
	tpl = make(map[string]string)
	tpl["code"] = "-1"

	switch action {
	case "img":
		data, _ := ioutil.ReadAll(c.Request.Body)
		if err := json.Unmarshal(data, &img); err == nil {
			if img.Delete() {
				tpl["code"] = "200"
			}
		}
		return
	case "All":
		return
	case "filter":
		return
	}
	c.JSON(200, gin.H{"message": "ok"})
}
