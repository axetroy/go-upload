package http

import (
	"github.com/gin-gonic/gin"
	"github.com/axetroy/go-upload/controller"
	"github.com/axetroy/go-upload/config"
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
	if err = fs.EnsureDir(path.Join(config.Upload.Path, config.Upload.File.Path)); err != nil {
		return
	}
	if err = fs.EnsureDir(path.Join(config.Upload.Path, config.Upload.Image.Path)); err != nil {
		return
	}
	if err = fs.EnsureDir(path.Join(config.Upload.Path, config.Upload.Image.Thumbnail.Path)); err != nil {
		return
	}
	return
}

func RunServer() (err error) {
	gin.SetMode(config.Config.Mode)
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
		Methods:         "GET, POST",
		RequestHeaders:  "Origin, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          60 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))
	Router.Use(gin.Recovery())

	// upload all
	uploader := Router.Group("/upload")

	uploader.POST("/image", controller.UploaderImage)
	uploader.POST("/file", controller.UploadFile)

	if config.Config.Mode != gin.ReleaseMode {
		uploader.GET("/example", controller.UploaderTemplate("image"))
	}

	// download all
	downloader := Router.Group("/download")

	// download image
	downloadImage := downloader.Group("/image")
	downloadImage.GET("/thumbnail/:filename", controller.GetThumbnailImage)
	downloadImage.GET("/origin/:filename", controller.GetOriginImage)

	// download file
	uploadFile := downloader.Group("/file")
	uploadFile.GET("/raw/:filename", controller.GetFileRaw)
	uploadFile.GET("/download/:filename", controller.DownloadFile)

	Router.Use(controller.NotFound)

	return Router.Run(config.Http.Host + ":" + config.Http.Port)
}
