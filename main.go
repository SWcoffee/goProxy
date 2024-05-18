package main

import (
	"github.com/gin-gonic/gin"
	_ "goProxy/config"
	"goProxy/controller"
)

func main() {

	r := gin.Default()
	r.Any("/*path", controller.ProxyAll)
	r.Run(":8080")
}
