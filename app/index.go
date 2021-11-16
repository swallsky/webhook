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

	token := c.Request.Header["X-Gitlab-Token"][0]
	ref := postData["ref"].(string) //将interface {}类型转换为string

	//执行相应的脚本文件
	res := ShellBin(token, ref)

	// Info 级别日志
	Logger().WithFields(logrus.Fields{
		"token": token,
		"ref":   ref,
		"shell": res,
	}).Info(postData["user_name"], " push")

	c.JSON(200, gin.H{
		"status": 200,  //状态
		"msg":    "ok", //返回结果
		"shell":  res,  //返回执行的结果
	})
}
