package http

import (
	"github.com/gin-gonic/gin"
	"github.com/axetroy/go-upload/config"
	"github.com/axetroy/gin-uploader"
)

var (
	Router *gin.Engine
)

func RunServer() (err error) {
	gin.SetMode(config.Config.Mode)
	Router = gin.Default()
	Router.Use(gin.Logger())
	Router.Use(gin.Recovery())

	if upload, download, err := uploader.New(Router, config.Config.Upload); err != nil {
		return err
	} else {
		if err := uploader.Resolve(upload, download); err != nil {
			return err
		}
	}

	return Router.Run(config.Http.Host + ":" + config.Http.Port)
}
