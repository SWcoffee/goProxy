package utils

import (
	"github.com/sirupsen/logrus"
	"os"
)

// LogConfig 配置结构体
type LogConfig struct {
	Level      string
	OutputFile string
	JSONFormat bool
}

// Logger 是一个封装了Logrus的结构体
type Logger struct {
	*logrus.Logger
}

// NewLogger 创建并初始化一个新的Logger实例
func NewLogger(config LogConfig) *Logger {
	log := logrus.New()

	// 设置日志级别
	level, err := logrus.ParseLevel(config.Level)
	if err != nil {
		log.SetLevel(logrus.InfoLevel) // 默认级别为Info
	} else {
		log.SetLevel(level)
	}

	// 设置日志格式
	if config.JSONFormat {
		log.SetFormatter(&logrus.JSONFormatter{})
	} else {
		log.SetFormatter(&logrus.TextFormatter{})
	}

	// 设置日志输出
	if config.OutputFile != "" {
		file, err := os.OpenFile(config.OutputFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			log.Out = file
		} else {
			log.Out = os.Stdout // 文件打开失败则输出到标准输出
			log.Warn("Failed to log to file, using default stderr")
		}
	} else {
		log.Out = os.Stdout // 默认输出到标准输出
	}

	return &Logger{log}
}

// 示例使用函数
func (l *Logger) Info(args ...interface{}) {
	l.Logger.Infoln(args...)
}

func (l *Logger) Debug(args ...interface{}) {
	l.Logger.Debugln(args...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.Logger.Warnln(args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.Logger.Errorln(args...)
}
