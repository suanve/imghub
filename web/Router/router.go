package Router

import (
	"imgHub/web/Controllers"
	"imgHub/web/Middlewares"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()
	router.Use(Middlewares.Cors())
	router.Static("/static/", "./web/Views/static/")
	router.LoadHTMLGlob("./web/Views/*.html")

	router.GET("/", Controllers.Index)
	router.POST("/upload", Controllers.Upload)
	router.GET("/uploads/:IMGID", Controllers.ReadImg)

	manage := router.Group("/manage")
	manage.GET("/", Controllers.ManageIndex)
	manage.POST("/login", Controllers.ManageLogin)
	manage.Use(Middlewares.JWTAuth())
	{
		manage.GET("/get/:action", Controllers.ManageGet)
		manage.POST("/del/:action", Controllers.ManageDel)
	}

	router.Run(":8081")
}
