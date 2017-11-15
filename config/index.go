package config

import (
	"path/filepath"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/axetroy/gin-uploader"
)

type ConfigType struct {
	Http   HttpConfig
	Upload uploader.TConfig
	Env    string
	Mode   string
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

	Config.Env = os.Getenv("GO_ENV")

	if Config.Env == gin.ReleaseMode || Config.Env == "production" || Config.Env == "publish" {
		Config.Mode = gin.ReleaseMode
	} else if Config.Env == gin.TestMode {
		Config.Mode = gin.TestMode
	} else {
		Config.Mode = gin.DebugMode
	}

	InitPaths()
	InitHttp()
	InitUpload()

	return
}
