package Router

import (
	"github.com/gin-gonic/gin"
	"imgHub/web/Controllers"
	"imgHub/web/Middlewares"
)

func InitRouter() {
	router := gin.Default()
	router.Use(Middlewares.Cors())

	router.GET("/", Controllers.Index)
	router.POST("/upload", Controllers.Upload)
	router.GET("/uploads/:IMGID", Controllers.ReadImg)

	manage := router.Group("/manage")
	{
		manage.GET("",Controllers.ManageIndex)
	}

	router.Run(":8081")
}
