package main

import (
	"webhook.com/route"
)

func main() {
	r := route.InitRouter()
	r.Run(":8080")
}
