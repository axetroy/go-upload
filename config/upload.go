package config

import "github.com/axetroy/gin-uploader"

var Upload uploader.TConfig

func InitUpload() {
	Upload = Config.Upload
}
