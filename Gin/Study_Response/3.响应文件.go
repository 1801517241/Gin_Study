package main

import "github.com/gin-gonic/gin"

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.GET("", func(context *gin.Context) {
		context.Header("Content-Type", "application/octet-stream")
		context.Header("Content-Disposition", "attachment; filename=当前页面源代码.go")
		context.File("Gin/Study_Response/3.响应文件.go")
	})

	err := router.Run("127.0.0.1:80")
	if err != nil {
		return
	}
}
