package main

import "github.com/gin-gonic/gin"

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Index(context *gin.Context) {
	context.JSON(200, Response{
		Code: 0,
		Msg:  "成功",
		Data: map[string]any{},
	})
}

func main() {
	//关闭调试日志
	//gin.SetMode("release")

	//初始化引擎
	router := gin.Default()

	router.GET("/index", Index)

	err := router.Run("127.0.0.1:80")
	if err != nil {
		return
	}
}
