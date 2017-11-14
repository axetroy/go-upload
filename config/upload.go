package config

type UploadConfig struct {
	Path string
	File struct {
		Path      string
		MaxSize   int
		AllowType []string
	}
	Image struct {
		Path    string
		MaxSize int
		Thumbnail struct {
			Path      string
			MaxWidth  int
			MaxHeight int
		}
	}
}

var Upload UploadConfig

func InitUpload() {
	Upload = Config.Upload
}
