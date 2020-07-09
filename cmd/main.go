package main

import (
	"flag"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/yyff/go-service/conf"
	"github.com/yyff/go-service/dao"
	"github.com/yyff/go-service/server/http"
)

func main() {
	flag.Parse()
	config := conf.New()
	err := initLog(config)
	if err != nil {
		log.Fatal(err)
	}
	dao, err := dao.New(config)
	if err != nil {
		log.Fatal(err)
	}
	http.InitServices(dao)
	http.Run(config)
}

var levelDict = map[string]log.Level{
	"debug": log.DebugLevel,
	"info":  log.InfoLevel,
	"warn":  log.WarnLevel,
	"error": log.ErrorLevel,
	"fatal": log.FatalLevel,
}

func initLog(config *conf.Config) error {
	logConf := config.Log
	// 设置日志格式为text格式
	log.SetFormatter(&log.TextFormatter{})

	l, ok := levelDict[logConf.Level]
	if !ok {
		return fmt.Errorf("invalid level: %v", logConf.Level)
	}
	log.SetLevel(l)

	if logConf.OutputFile != "" {
		f, err := os.OpenFile(logConf.OutputFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			return err
		}
		log.SetOutput(f)
	} else {
		log.SetOutput(os.Stdout)
	}
	return nil
}
