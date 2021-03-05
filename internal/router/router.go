package router

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"go-blog/global"
	"go-blog/internal/middleware"
	"go-blog/internal/router/api"
	"go-blog/internal/router/api/music"
	v1 "go-blog/internal/router/api/v1"
	"net/http"
)


func NewRouter() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	engine.Use(middleware.Translations())
	url := ginSwagger.URL("http:127.0.0.1:8000/swagger/doc.json")
	engine.GET("/swagger/*any",ginSwagger.WrapHandler(swaggerFiles.Handler,url))

	tag := v1.Tag{}
	article := v1.Article{}
	student :=  music.Student{}

	upload := NewUpload()
	engine.POST("/upload/file",upload.UploadFile)
	engine.StaticFS("/static",http.Dir(global.AppSetting.UploadSavePath))

	//获取token
	engine.GET("/auth",api.GetAuth)
	apiv1 := engine.Group("/api/v1")
	apiv1.Use(middleware.JWT())
	{
		//标签路由管理
		apiv1.POST("/tags",tag.Create)
		apiv1.DELETE("/tags/:id",tag.Delete)
		apiv1.PUT("/tags/:id",tag.Update)
		apiv1.PATCH("/tags/:id/state",tag.Update)
		apiv1.GET("/tags",tag.List)

		//文章路由管理
		apiv1.POST("/articles",article.Create)
		apiv1.DELETE("/articles/:id",article.Delete)
		apiv1.PUT("/articles/:id",article.Update)
		apiv1.PATCH("/articles/:id/state",article.Update)
		apiv1.GET("/articles/:id",article.Get)
		apiv1.GET("/articles",article.List)

		//列表接口测试
		apiv1.GET("/article",article.Lists)
	}

	ap1v2 := engine.Group("/api/music")
	{
		ap1v2.GET("/students",student.List)
	}

	return engine
}
