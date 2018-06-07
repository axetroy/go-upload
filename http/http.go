package http

import (
	"github.com/gin-gonic/gin"
	"github.com/axetroy/go-upload/config"
	"github.com/axetroy/gin-uploader"
	"github.com/itsjamie/gin-cors"
	"time"
)

var (
	Router *gin.Engine
)

func RunServer() (err error) {
	gin.SetMode(config.Config.Mode)
	Router = gin.Default()
	Router.Use(gin.Logger())
	Router.Use(gin.Recovery())

	Router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, X-Requested-With, X_Requested_With, Content-Type, Accept, Authentication, Authorization, X-Server, X-Password-Pay",
		ExposedHeaders:  "",
		MaxAge:          60 * time.Second,
		Credentials:     true,		// cookies
		ValidateHeaders: false,
	}))

	if loader, err := uploader.New(Router, config.Config.Upload); err != nil {
		return err
	} else {
		loader.Resolve()
	}

	return Router.Run(config.Http.Host + ":" + config.Http.Port)
}
