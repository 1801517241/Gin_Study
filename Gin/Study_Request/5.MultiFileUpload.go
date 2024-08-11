package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("users", func(context *gin.Context) {
		form, err := context.MultipartForm()
		if err != nil {
			return
		}
		for _, headers := range form.File {
			for _, header := range headers {
				err := context.SaveUploadedFile(header, "uploads/xxx/zzz/"+header.Filename)
				if err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	})

	err := router.Run("127.0.0.1:80")
	if err != nil {
		return
	}
}
