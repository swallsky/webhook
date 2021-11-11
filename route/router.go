package route

import (
	"github.com/gin-gonic/gin"
	"webhook.com/app"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	//记录访问日志
	r.Use(app.LoggerToFile())
	//首页
	r.GET("/", app.Home)
	//日志测试
	r.GET("/logtest", app.LogTest)

	return r
}
