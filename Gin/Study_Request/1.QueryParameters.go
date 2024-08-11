package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Query(context *gin.Context) {
	name := context.Query("name")
	age := context.DefaultQuery("age", "25")
	keyList := context.QueryArray("key")
	fmt.Println(name, age, keyList)
}

func main() {
	Address := "127.0.0.1:80"
	router := gin.Default()

	router.GET("", Query)

	err := router.Run(Address)
	if err != nil {
		return
	}
}
