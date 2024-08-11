package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Dynamic(context *gin.Context) {
	/**
	 * 获取 URL 参数中的用户 ID。
	 * 使用 context.Param 函数来获取 URL 中的参数值。这里的 "id" 是在路由中定义的参数名，用于标识用户的唯一 ID。
	 * 获取到用户 ID 后，将其打印到控制台。
	 * 打印用户 ID 到控制台，以便查看和调试。
	 * 请注意，这里没有处理获取用户 ID 失败的情况，在实际应用中应加入适当的错误处理逻辑。
	 */
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
