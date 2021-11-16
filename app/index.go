package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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

	// Info 级别日志
	Logger().WithFields(logrus.Fields{
		"token": c.Request.Header["X-Gitlab-Token"][0],
		"ref":   postData["ref"],
	}).Info(postData["user_name"], " push")

	c.JSON(200, gin.H{
		"status": 200,  //状态
		"msg":    "ok", //返回结果
	})
}
