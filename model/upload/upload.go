package upload

import (
	"errors"
	"io"
	"math/rand"
	"os"
	"mime/multipart"
	"regexp"
	"strconv"
	"strings"
	"time"

	"imageWarehouse/config"
)

const (
	RAND_CHAR = "0123456789abcedfghigklmnopqrstuvwxvzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	NAME_LENGTH = 15
)

var allowExt []string = []string{"jpg", "JPG", "JPEG", "png", "PNG", "gif", "GIF"}

type ImageInfo struct {
	Name string		// 文件名
	Ext string		// 后缀
	OriginName string 	// 原始文件名
	FilePath string 	// 文件位置
}

func InitImageInfo(file *multipart.FileHeader) (*ImageInfo, error) {
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	imageInfo := &ImageInfo{
		OriginName: file.Filename,
		FilePath: config.ImagePath,
	}
	err = imageInfo.FileExt()
	if err != nil {
		return nil, err
	}
	err = imageInfo.CopyTmpFile(src)

	if err != nil {
		return nil, err
	}

	return imageInfo, nil
}

func (image *ImageInfo) CopyTmpFile(handle io.Reader) error {
	image.Name = RandImageName()
	image.FilePath = image.FilePath + "/" + image.Name + "." + image.Ext
	out, err := os.Create(image.FilePath)
	if err != nil {
		return err
	}

	defer out.Close()
	_, err = io.Copy(out, handle)

	return err
}

func (image *ImageInfo) FileExt() error {
	reg := regexp.MustCompile("\\.(.*?)$")
	ext := reg.FindString(image.OriginName)

	ext = strings.TrimPrefix(ext, ".")
	if len(ext) == 0 {
		return errors.New("File extension error")
	}

	for item := range allowExt {
		if allowExt[item] == ext {
			image.Ext = strings.ToLower(ext)
			return nil
		}
	}

	return errors.New("File extension error")
}

func UploadImage(file *multipart.FileHeader) (string, error) {
	imageInfo, err := InitImageInfo(file)
	if(err != nil) {
		return "", err
	}

	return config.StaticUrl + imageInfo.Name + "." + imageInfo.Ext, nil
}

func RandImageName() string {
	var s []uint8
	rand.Seed(time.Now().Unix())
	max := len(RAND_CHAR) - 1
	for i := 0; i < NAME_LENGTH; i ++ {
		s = append(s, RAND_CHAR[rand.Intn(max)])
	}
	t := strconv.FormatInt(time.Now().UnixNano(), 10)
	return string(s) + t
}