package main

import (
	"go.uber.org/zap"
)

func main() {
	//logger, _ := zap.NewDevelopment() //测试环境
	logger, _ := zap.NewProduction() //生成环境

	defer logger.Sync() // flushes buffer, if any
	url := "https://baidu.com"

	//logger.Info("failed to fetch URL", zap.String("URL", url), zap.Int("num", 3))

	sugar := logger.Sugar()

	//输出的是json格式
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
	)
	sugar.Infof("Failed to fetch URL: %s", url)
	//还支持很多,如Error,Panic...
	//sugar.Error()
	//sugar.Panic()

	/*
		Zap提供了两种类型的日志记录器—Sugared Logger和Logger。
		在性能很好但不是很关键的上下文中，使用SugaredLogger。它比其他结构化日志记录包快4-10倍，并且支持结构化和printf风格的日志记录。
		在每一微秒和每一次内存分配都很重要的上下文中，使用Logger。它甚至比SugaredLogger更快，内存分配次数也更少，但它只支持强类型的结构化日志记录.
	*/
}
