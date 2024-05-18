package config

import (
	"goProxy/utils"
	"os"
)

var (
	CurrentUrl = "https://baidu.com"
	ProxyUrl   = "https://oaistatic-cdn.closeai.biz"

	Log *utils.Logger
)

func init() {
	currentUrl := os.Getenv("CURRENT_URL")
	if currentUrl != "" {
		CurrentUrl = currentUrl
	}

	proxyUrl := os.Getenv("PROXY_URL")
	if proxyUrl != "" {
		ProxyUrl = proxyUrl
	}

	logConfig := utils.LogConfig{
		Level:      "debug",
		OutputFile: "app.log",
		JSONFormat: true,
	}
	Log = utils.NewLogger(logConfig)

	Log.Info("CurrentUrl:", CurrentUrl)
	Log.Info("ProxyUrl:", ProxyUrl)
}
