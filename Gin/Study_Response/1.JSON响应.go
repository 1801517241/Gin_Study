package main

import (
	"Gin_Study_New/Gin/Study_Response/res"
	"github.com/gin-gonic/gin"
)

func json(context *gin.Context) {
	//context.JSON(200, gin.H{"code": "0","msg":  "ok","data": gin.H{},})
	res.OkWithMsg(context, "登录成功")
}

func users(context *gin.Context) {
	res.OkWithData(context, map[string]any{
		"Name": "Fovik",
	})
}

func usersFail(context *gin.Context) {
	res.FailWithMsg(context, "参数错误")
}

func main() {
	router := gin.Default()

	router.GET("/json", json)
	router.GET("/users", users)
	router.POST("/users", usersFail)

	err := router.Run("127.0.0.1:80")
	if err != nil {
		return
	}
}
