package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/assets", "static/assets")
	router.StaticFS("/bootstrap-5", http.Dir("static/bootstrap-5"))
	router.StaticFile("/favicon.ico", "static/resources/favicon.ico")

	// 监听并在 0.0.0.0:8080 上启动服务
	router.Run(":8080")
}
