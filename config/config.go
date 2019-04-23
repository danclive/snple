package config

import (
	"github.com/danclive/mqtt-console/log"

	"github.com/BurntSushi/toml"
)

var Config struct {
	Debug bool
	Api   Api
	Mongo Mongo
}

type Api struct {
	Addr string
}

type Mongo struct {
	Uri      string
	Database string
}

func InitConfig(file string) {
	if _, err := toml.DecodeFile(file, &Config); err != nil {
		log.Fatal("无法读取配置文件: ", err)
	}
}
