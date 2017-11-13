package config

type UploadConfig struct {
	Path               string
	Image              string
	Thumbnail          string
	ThumbnailMaxWidth  int      // 缩略图最大宽度, 暂无用处，默认300
	ThumbnailMaxHeight int      // 缩略图最大高度，暂无用处，默认300
	AllowType          []string // if this is a empty array, then allow all file type
	MaxSize            int64    // 上传图片的最大文件大小，单位byte
}

var Upload UploadConfig

func InitUpload() {
	Upload = Config.Upload
}
