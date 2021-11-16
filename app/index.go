package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"

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
	// 获取提交字段
	var postData map[string]interface{}
	data, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(data, &postData)

	fmt.Println(reflect.TypeOf(postData["commits"]))
	fmt.Println(reflect.TypeOf(postData["commits"]))

	c.JSON(200, gin.H{
		"token":         c.Request.Header["X-Gitlab-Token"][0], //获取webhook的token值
		"ref":           postData["ref"],                       //提交的版本信息
		"user_name":     postData["user_name"],                 //gitlab代码提交者昵称
		"user_username": postData["user_username"],             //gitlab代码提交者用户名
		"commits":       postData["commits"],                   //本地提交的信息
	})
}
