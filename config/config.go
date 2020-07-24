package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type ConfigInfo struct {
	ImagePath string 	`json:"image_path"`
	TmpImagePath string 	`json:"tmp_image_path"`
	StaticUrl string 	`json:"static_url"`
}

var ImagePath string
var TmpImagePath string
var StaticUrl string

func InitConfig(configPath string) error {
	configFile, err := os.Open(configPath)
	if err != nil {
		return err
	}
	bytes, _ := ioutil.ReadAll(configFile)
	configFile.Close()

	var configInfo *ConfigInfo
	err = json.Unmarshal(bytes, &configInfo)
	if err != nil {
		return err
	}
	ImagePath = configInfo.ImagePath
	TmpImagePath = configInfo.TmpImagePath
	StaticUrl = configInfo.StaticUrl

	return nil
}
