package http

import (
	"github.com/gin-gonic/gin"
	"github.com/axetroy/go-upload/server/controller"
	"github.com/axetroy/go-upload/server/config"
	"github.com/itsjamie/gin-cors"
	"github.com/axetroy/go-fs"
	"time"
	"path"
)

var (
	Router *gin.Engine
)

func Init() (err error) {
	if err = fs.EnsureDir(config.Upload.Path); err != nil {
		return
	}
	if err = fs.EnsureDir(path.Join(config.Upload.Path, config.Upload.Image)); err != nil {
		return
	}
	if err = fs.EnsureDir(path.Join(config.Upload.Path, config.Upload.Thumbnail)); err != nil {
		return
	}
	return
}

func RunServer() (err error) {

	Router = gin.Default()
	Router.Use(gin.Logger())
	Router.Use(func(context *gin.Context) {
		header := context.Writer.Header()
		// alone dns prefect
		header.Set("X-DNS-Prefetch-Control", "on")
		// IE No Open
		header.Set("X-Download-Options", "noopen")
		// not cache
		header.Set("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate")
		header.Set("Expires", "max-age=0")
		// Content Security Policy
		header.Set("Content-Security-Policy", "default-src 'self'")
		// xss protect
		// it will caught some problems is old IE
		header.Set("X-XSS-Protection", "1; mode=block")
		// Referrer Policy
		header.Set("Referrer-Header", "no-referrer")
		// cros frame, allow same origin
		header.Set("X-Frame-Options", "SAMEORIGIN")
		// HSTS
		header.Set("Strict-Transport-Security", "max-age=5184000;includeSubDomains")
		// no sniff
		header.Set("X-Content-Type-Options", "nosniff")
	})
	Router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))
	Router.Use(gin.Recovery())

	Router.POST("/upload", controller.Uploader)
	Router.GET("/upload/example", controller.UploaderTemplate)
	Router.GET("/download/:size/:filename", controller.Downloader)

	Router.Use(controller.NotFound)

	return Router.Run(config.Http.Host + ":" + config.Http.Port)
}
