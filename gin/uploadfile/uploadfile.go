package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type result struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Result interface{} `json:"result"`
}

func main() {
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20
	router.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		if file == nil {
			c.JSON(http.StatusOK, &result{
				Code: http.StatusBadRequest,
				Msg:  "没有要上传的文件,请先选择要上传的文件!",
			})
			return
		}
		err := c.SaveUploadedFile(file, file.Filename)
		if err != nil {
			c.JSON(http.StatusOK, &result{
				Code: http.StatusInternalServerError,
				Msg:  "上传文件失败,请重新上传!",
			})
			log.Println("上传文件失败,错误为:", err)
			return
		}
		c.JSON(http.StatusOK, &result{
			Code: http.StatusOK,
			Msg:  "上传成功",
		})
	})
	router.Run(":9000")
}
