package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Dynamic(context *gin.Context) {
	userID := context.Param("id")
	fmt.Println(userID)
}

func main() {
	router := gin.Default()

	router.GET("users/:id", Dynamic)

	err := router.Run("127.0.0.1:80")
	if err != nil {
		return
	}
}
