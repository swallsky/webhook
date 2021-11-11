package app

import (
	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.JSON(200, gin.H{
		"dsn": "dsn",
	})
}

/**
 * post测试
 */
func Post(c *gin.Context) {

}
