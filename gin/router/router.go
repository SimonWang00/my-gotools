package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"my-gotools/gin/controller"
	_ "my-gotools/gin/docs"
)

func SetRouters(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))
	announcement := &controller.AnnouncementController{}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1:=r.Group("/v1")
	{
		//公告相关业务
		v1.GET("/announcement/:id", announcement.GetById)
		v1.GET("/announcement/", announcement.GetAll)
		v1.POST("/announcement", announcement.Add)
		v1.PUT("/announcement/:id", announcement.UpDate)
		v1.DELETE("/announcement/:id", announcement.Del)
	}
}
