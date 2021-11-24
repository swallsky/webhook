package route

import (
	"github.com/gin-gonic/gin"
	"webhook.com/app"
)

func InitRouter() *gin.Engine {
	//设置模式
	// gin.SetMode(gin.DebugMode) //默认模式
	// gin.SetMode(gin.ReleaseMode)
	// gin.SetMode(gin.TestMode)

	r := gin.Default()
	//记录访问日志 中间件
	// r.Use(app.LoggerToFile())
	//首页
	r.GET("/", app.Home)
	//webhooks
	r.POST("webhooks", app.WebHooks)

	return r
}
