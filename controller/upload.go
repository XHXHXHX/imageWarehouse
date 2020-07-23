package controller

import (
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

	err = upload.UploadImage(image)

	if err != nil {
		c.JSON(200, gin.H{"err": err})
		return
	}

	c.JSON(200, gin.H{"msg": "success"})
}