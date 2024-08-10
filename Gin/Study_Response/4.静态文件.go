package main

import "github.com/gin-gonic/gin"

func main() {
	//gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	//第一个参数是别名，第二个参数是路劲
	router.Static("st", "Gin/Study_Response/static")
	//指定一个别名到文件
	router.StaticFile("abc", "Gin/Study_Response/static/abc.txt")

	//静态文件的路径不能再被路由使用
	//router.GET("abc", func(context *gin.Context) {
	//	context.String(200, "Hello World")
	//})

	err := router.Run("127.0.0.1:80")
	if err != nil {
		return
	}
}
