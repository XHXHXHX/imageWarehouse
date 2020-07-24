package main

import (
	"fmt"
	"imageWarehouse/config"
	"imageWarehouse/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	defer MainPanicHandler()

	err := config.InitConfig("config.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	router := gin.New()

	router.LoadHTMLGlob("views/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.html", gin.H{})
	})

	uploadCtr := &controller.Upload{}
	router.POST("/upload/image", uploadCtr.UploadImage)

	_ = router.Run(":7111")
}

func MainPanicHandler() {
	if err := recover(); err != nil {

	}
}