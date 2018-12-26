package routers

import (
	"fmt"
	"gin-blog/middleware/jwt"
	"gin-blog/pkg/logging"
	"gin-blog/pkg/setting"
	"gin-blog/pkg/upload"
	"gin-blog/pkg/util"
	"gin-blog/routers/api"
	"gin-blog/routers/api/v1"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path"
	"time"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 修改gin的日志格式, 可以配合日志分析软件来进行通用的日志分析
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	gin.SetMode(setting.ServerSetting.RunMode)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		// 新建标签
		apiv1.POST("/tags", v1.AddTag)
		// 更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		// 删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTags)

		// 获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		// 获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//	新建文章
		apiv1.POST("/articles", v1.AddArticle)
		// 修改文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		// 删除文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	r.POST("/upload", api.UploadImage)

	// 官方提供的上传方法
	r.POST("/upload1", func(c *gin.Context) {
		file, _ := c.FormFile("image")

		log.Println(file.Filename)

		err := c.SaveUploadedFile(file, "runtime/upload/images/"+util.EncodeMD5(file.Filename)+path.Ext(file.Filename))
		if err != nil {
			logging.Error(err)
		}

		c.String(http.StatusOK, "upload success")
	})

	// 查看上传的图片
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))

	r.GET("/auth", api.GetAuth)

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})
	return r
}
