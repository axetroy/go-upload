package controller

import (
	"crypto/md5"
	"io"
	"encoding/hex"
	"path"
	"os"
	"net/http"
	"strings"
	"image/jpeg"
	"image"
	"image/png"
	"image/gif"
	"errors"
	"strconv"
	"log"
	"github.com/nfnt/resize"
	"github.com/axetroy/go-fs"
	"github.com/axetroy/go-upload/config"
	"github.com/gin-gonic/gin"
)

/**
生成图片的缩略图
 */
func thumbnailify(imagePath string) (outputPath string, err error) {
	var (
		file *os.File
		img  image.Image
	)

	filename := path.Base(imagePath)

	extname := strings.ToLower(path.Ext(imagePath))

	outputPath = path.Join(config.Upload.Path, config.Upload.Thumbnail, filename)

	// 读取文件
	if file, err = os.Open(imagePath); err != nil {
		return
	}

	// decode jpeg into image.Image
	switch extname {
	case ".jpg", ".jpeg":
		img, err = jpeg.Decode(file)
		break
	case ".png":
		img, err = png.Decode(file)
		break
	case ".gif":
		img, err = gif.Decode(file)
		break
	default:
		err = errors.New("Unsupport file type" + extname)
		return
	}

	if img == nil {
		err = errors.New("Generate thumbnail fail...")
		return
	}

	defer file.Close()

	m := resize.Thumbnail(300, 300, img, resize.Lanczos3)

	out, err := os.Create(outputPath)
	if err != nil {
		return
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)
	return
}

/**
update handler
 */
func Uploader(context *gin.Context) {
	var (
		isSupportFile bool
		maxUploadSize = config.Upload.MaxSize
	)
	// Source
	file, _ := context.FormFile("file")

	extname := path.Ext(file.Filename)

	for i := 0; i < len(config.Upload.AllowType); i++ {
		if config.Upload.AllowType[i] == extname {
			isSupportFile = true
		}
	}

	if isSupportFile == false {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Unsupport upload file type " + extname,
		})
		return
	}

	if file.Size > maxUploadSize {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Upload file too large, The max upload limit is " + strconv.Itoa(int(maxUploadSize)),
		})
		return
	}

	src, _ := file.Open()
	defer src.Close()

	hash := md5.New()

	io.Copy(hash, src)

	md5string := hex.EncodeToString(hash.Sum([]byte("")))

	fileName := md5string + extname

	// Destination
	absPath := path.Join(config.Upload.Path, config.Upload.Image, fileName)
	dst, _ := os.Create(absPath)
	defer dst.Close()

	// FIXME: open 2 times
	src, _ = file.Open()

	// Copy
	io.Copy(dst, src)

	// 压缩缩略图
	// 不管成功与否，都会进行下一步的返回
	if _, err := thumbnailify(absPath); err != nil {
		log.Fatal("生成缩略图失败")
	}

	context.JSON(http.StatusOK, gin.H{
		"hash":     md5string,
		"filename": fileName,
		"origin":   file.Filename,
		"size":     file.Size,
	})
}

func UploaderTemplate(context *gin.Context) {
	header := context.Writer.Header()
	header.Set("Content-Type", "text/html; charset=utf-8")
	context.String(200, `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>File Upload</title>
</head>
<body>
<form action="/upload" method="post" enctype="multipart/form-data">
  <h2>Image Upload</h2>
  <input type="file" name="file">
  <input type="submit" value="Submit">
</form>
</body>
</html>
	`)
}

/**
download handler
 */
func Downloader(context *gin.Context) {

	var (
		filepath       string // 根据url，对应的图片路径
		originFilePath string // 原始文件路径
	)

	size := context.Param("size")
	file := context.Param("filename")

	// 两个参数都不传，则返回404
	if size == "" || file == "" {
		http.NotFound(context.Writer, context.Request)
		return
	}

	// 如果路径是以/结尾的话，那么这个图片不存在
	if strings.HasSuffix(context.Request.URL.Path, "/") {
		http.NotFound(context.Writer, context.Request)
		return
	}

	// 尺寸不对
	if size != config.Upload.Image && size != config.Upload.Thumbnail {
		http.NotFound(context.Writer, context.Request)
		return
	}

	fileType := size

	filepath = path.Join(config.Upload.Path, fileType, file)

	originFilePath = path.Join(config.Upload.Path, config.Upload.Image, file)

	if isExistFile := fs.PathExists(filepath); isExistFile == false {
		// 如果是获取缩略图，获取失败的话，尝试获取原图
		if size == config.Upload.Thumbnail {
			// 如果原图存在，则返回原图
			if isExistOriginFIle := fs.PathExists(originFilePath); isExistOriginFIle == true {
				http.ServeFile(context.Writer, context.Request, originFilePath)
				return
			}
		}

		// if the path not found
		http.NotFound(context.Writer, context.Request)
		return
	}

	http.ServeFile(context.Writer, context.Request, filepath)
}
