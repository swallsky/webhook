package app

import (
	"encoding/json"
	"io/ioutil"

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
func IndexPost(c *gin.Context) {
	var postData map[string]interface{}
	data, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(data, &postData)

	c.JSON(200, gin.H{
		"ref": postData["ref"],
	})
}
