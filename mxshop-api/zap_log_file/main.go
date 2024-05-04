package main

import (
	"go.uber.org/zap"
	"time"
)

func NewLogger() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		//输出到当前项目目录的日志文件
		"./myproject.log",
		//标准错误输出
		"stderr",
		//标准输出
		"stdout",
	}
	return cfg.Build()
}

func main() {
	//logger, _ := zap.NewProduction()
	logger, err := NewLogger()
	if err != nil {
		panic(err)
		//panic("初始化logger失败")
	}
	su := logger.Sugar()
	defer su.Sync()
	url := "https://baidu.com"
	su.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}
