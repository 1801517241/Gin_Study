package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	fmt.Println("Home")
	fmt.Println(c.Get("GM1"))
	fmt.Println(c.Get("GM2"))
	c.String(200, "Home")
}

func M1(c *gin.Context) {
	fmt.Println("M1 请求部分")
	c.Abort()
	c.String(200, "M1返回的响应")
	fmt.Println("M1 响应部分")
}

func M2(c *gin.Context) {
	fmt.Println("M2 请求部分")
	c.Next()
	fmt.Println("M2 响应部分")
}

func GM1(c *gin.Context) {
	fmt.Println("GM1 请求部分")
	c.Set("GM1", "GM1")
	c.Next()
	fmt.Println(c.Get("GM1"))
	fmt.Println("GM1 响应部分")
}

func GM2(c *gin.Context) {
	fmt.Println("GM2 请求部分")
	c.Set("GM2", "GM2")
	c.Next()
	fmt.Println(c.Get("GM2"))
	fmt.Println("GM2 响应部分")
}

func AuthMiddleware(c *gin.Context) {

}

func main() {
	r := gin.Default()
	g := r.Group("api")
	g.Use(GM1)
	g.Use(GM2)
	g.GET("users", Home)
	err := r.Run("127.0.0.1:80")
	if err != nil {
		return
	}
}
