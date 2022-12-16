package main

import (
	"github.com/cczyWyc/crawler-geektime/collect"
	"github.com/cczyWyc/crawler-geektime/log"
	"github.com/cczyWyc/crawler-geektime/proxy"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

func main() {
	plugin, c := log.NewFilePlugin("./crawler.log", zapcore.InfoLevel)
	defer c.Close()
	logger := log.NewLogger(plugin)
	logger.Info("log init succeed")

	proxyURLs := []string{"http://127.0.0.1:7890"}
	p, err := proxy.RoundRobinProxySwitcher(proxyURLs...)
	if err != nil {
		logger.Error("RoundRobinProxySwitcher failed")
	}
	url := "https://www.google.com"

	var f collect.FetCher = collect.BrowserFetch{
		Timeout: 3000 * time.Millisecond,
		Proxy:   p,
	}
	body, err := f.Get(url)
	if err != nil {
		logger.Error("read content failed", zap.Error(err))
		return
	}
	logger.Info("get content", zap.Int("len", len(body)))
}
