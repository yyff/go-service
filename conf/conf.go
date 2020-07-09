package conf

import (
	log "github.com/sirupsen/logrus"

	"github.com/BurntSushi/toml"
)

const (
	confFile = "go-service.toml"
)

type Config struct {
	Http  *Http
	MySQL *MySQL
	Log   *Log
}

type MySQL struct {
	DSN string
}

type Http struct {
	Port         uint32
	ReadTimeout  uint32 // unit: ms
	WriteTimeout uint32 // unit: ms
}

type Log struct {
	Level      string
	OutputFile string
}

func New() *Config {

	config := new(Config)
	if _, err := toml.DecodeFile(confFile, config); err != nil {
		log.Fatal(err)
	}
	return config
}
