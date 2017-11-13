package config

import (
	"path/filepath"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type ConfigType struct {
	Http   HttpConfig
	Upload UploadConfig
}

var Config ConfigType

func Init() (err error) {

	var (
		configfile string
		yamlFile   []byte
	)

	configfile, err = filepath.Abs("./config.yaml")

	if err != nil {
		return
	}

	yamlFile, err = ioutil.ReadFile(configfile)

	if err != nil {
		return
	}

	err = yaml.Unmarshal(yamlFile, &Config)

	if err != nil {
		return err
	}

	InitPaths()
	InitHttp()
	InitUpload()

	return
}
