package upload

import (
	"io"
	"mime/multipart"
	"os"
)

func UploadImage(file *multipart.FileHeader) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	out, err := os.Create("new_file")
	if err != nil {
		return err
	}

	defer out.Close()
	_, err = io.Copy(out, src)
	return err
}
