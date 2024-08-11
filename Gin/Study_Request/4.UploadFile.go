package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("file", func(context *gin.Context) {
		fileHeader, err := context.FormFile("file")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(fileHeader.Filename)
		fmt.Println(fileHeader.Size/1000, "KB") //单位是字节

		//file, _ := fileHeader.Open()
		//byteData, _ := io.ReadAll(file)
		//err = os.WriteFile("xxx.jpg", byteData, 0666)
		//if err != nil {
		//	return
		//}

		err = context.SaveUploadedFile(fileHeader, "uploads/xxx/yyy/"+fileHeader.Filename)
		if err != nil {
			return
		}
	})

	err := router.Run("127.0.0.1:80")
	if err != nil {
		return
	}
}
