package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("json", func(context *gin.Context) {
		type User struct {
			IPList []string `json:"IPList" binding:"ip"`
		}
		var user User
		err := context.ShouldBindJSON(&user)
		if err != nil {
			context.String(200, err.Error())
			return
		}

		context.JSON(200, user)
	})

	err := router.Run("127.0.0.1:80")
	if err != nil {
		return
	}
}
