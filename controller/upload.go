package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"imageWarehouse/model/upload"
)

type Upload struct {

}

func (this *Upload) UploadImage(c *gin.Context) {
	image, err := c.FormFile("image")
	if err != nil {
		c.JSON(200, gin.H{"msg": "upload error"})
		return
	}

	image_name, err := upload.UploadImage(image)

	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{"err": err.Error()})
		return
	}

	c.JSON(200, gin.H{"msg": image_name})
}