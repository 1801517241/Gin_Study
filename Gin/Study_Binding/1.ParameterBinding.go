package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("", func(context *gin.Context) {
		type User struct {
			Name string `form:"name"`
			Age  int    `form:"age"`
		}
		var user User
		err := context.ShouldBindQuery(&user)
		if err != nil {
			return
		}
		fmt.Println(user, err)
	})

	router.GET("users/:id/:name", func(context *gin.Context) {
		type User struct {
			Name string `uri:"name"`
			ID   string `uri:"id"`
		}
		var user User
		err := context.ShouldBindUri(&user)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(user)
	})

	router.POST("form", func(context *gin.Context) {
		type User struct {
			Name string `form:"name"`
			Age  string `form:"age"`
		}
		var user User
		err := context.ShouldBind(&user)
		if err != nil {
			return
		}
		fmt.Println(user, err)
	})

	router.POST("json", func(context *gin.Context) {
		type User struct {
			Name string `json:"name"`
			Age  string `json:"age"`
		}
		var user User
		err := context.ShouldBindJSON(&user)
		fmt.Println(user, err)
	})

	router.POST("header", func(context *gin.Context) {
		type User struct {
			Name        string `header:"name"`
			Age         string `header:"Age"`
			UserAgent   string `header:"User-Agent"`
			contentType string `header:"Content-Type"`
		}
		var user User
		err := context.ShouldBindHeader(&user)
		fmt.Println(user, err)
	})

	err := router.Run("127.0.0.1:80")
	if err != nil {
		return
	}
}
