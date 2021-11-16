package route

import (
	"github.com/gin-gonic/gin"
	"webhook.com/app"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	//记录访问日志 中间件
	// r.Use(app.LoggerToFile())
	//首页
	r.GET("/", app.Home)
	//日志测试
	r.GET("/logtest", app.LogTest)
	//post测试
	r.POST("indexpost", app.IndexPost)
	// shell测试
	r.GET("shelltest", app.ShellTest)

	return r
}
