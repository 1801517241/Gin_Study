package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/url"
)

func main() {
	router := gin.Default()

	router.POST("", func(context *gin.Context) {
		byteData, err := io.ReadAll(context.Request.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		//读取之后，body自动销毁
		fmt.Println(string(byteData))

		context.Request.Body = io.NopCloser(bytes.NewReader(byteData))
		fmt.Println(context.Request.Header)

		name := context.PostForm("name")
		fmt.Println(name)

		_, err = url.Parse("?")
		if err != nil {
			return
		}
	})

	err := router.Run("127.0.0.1:80")
	if err != nil {
		return
	}
}
