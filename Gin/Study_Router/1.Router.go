package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func UserView(c *gin.Context) {
	path := c.Request.URL
	fmt.Println(c.Request.Method, path)
}

func main() {
	router := gin.Default()

	apiGroup := router.Group("api")
	apiGroup.Use()

	UserGroup(apiGroup)

	noMiddleWareGroup := router.Group("api")
	LoginGroup(noMiddleWareGroup)

	err := router.Run("127.0.0.1:80")
	if err != nil {
		return
	}
}

func UserGroup(r *gin.RouterGroup) {
	r.GET("users", UserView)
	r.POST("users", UserView)
	r.DELETE("users", UserView)
	r.PUT("users", UserView)
}

func LoginGroup(r *gin.RouterGroup) {
	r.GET("login", UserView)
}
