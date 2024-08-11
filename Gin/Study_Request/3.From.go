package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Users(context *gin.Context) {
	name := context.PostForm("name")
	age, ok := context.GetPostForm("age")
	if !ok {
		age = context.DefaultPostForm("age", "18")
	}
	fmt.Println(name)
	fmt.Println(age, ok)
}

func main() {
	router := gin.Default()

	router.POST("users", Users)

	err := router.Run("127.0.0.1:80")
	if err != nil {
		return
	}
}
