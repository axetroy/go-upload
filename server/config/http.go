package config

type HttpConfig struct {
	Host string
	Port string
}

var Http HttpConfig

func InitHttp() {
	Http = Config.Http
}
