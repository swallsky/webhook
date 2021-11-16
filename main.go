package main

import (
	"flag"

	"webhook.com/route"
)

func main() {
	var port string //web端口
	flag.StringVar(&port, "p", "8080", "请输入端口号")
	flag.Parse() //解析参数
	r := route.InitRouter()
	r.Run(":" + port)
}
