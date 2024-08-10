package main

import "github.com/gin-gonic/gin"

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.LoadHTMLGlob("Gin/Study_Response/templates/*")
	//router.LoadHTMLFiles("Gin/Study_Response/template/浏览器下载.html")
	router.GET("/", func(context *gin.Context) {
		context.HTML(200, "浏览器下载.html", map[string]any{
			"title": "Fovik_Blog",
		})
	})

	err := router.Run("127.0.0.1:80")
	if err != nil {
		return
	}
}
