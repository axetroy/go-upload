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

	err, _, _ = uploader.Resolve(Router, config.Config.Upload)

	return Router.Run(config.Http.Host + ":" + config.Http.Port)
}
