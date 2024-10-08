package logs

import (
	"chat-bot/src/config"
	"fmt"
	"go.uber.org/fx"
	"time"
)

// ConsoleLogger 是 Logger 接口的实现，输出日志到控制台
type ConsoleLogger struct {
	logLevel int // 用于控制输出的日志级别
}

// 日志等级常量
const (
	DebugLevel = iota // 0
	InfoLevel         // 1
	WarnLevel         // 2
	ErrorLevel        // 3
)

// Params 注入参数
type Params struct {
	fx.In
	Config *config.ChatBotConfig
}

// NewConsoleLogger 创建并返回一个 ConsoleLogger 实例
func NewConsoleLogger(p Params) *ConsoleLogger {
	return &ConsoleLogger{
		logLevel: p.Config.LogLevel,
	}
}

// log 是一个通用的日志输出方法，根据不同的等级输出不同的内容
func (l *ConsoleLogger) log(level int, levelStr, message string) {
	if level >= l.logLevel {
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		fmt.Printf("[%s] [%s]: %s\n", timestamp, levelStr, message)
	}
}

// Debug 输出调试信息
func (l *ConsoleLogger) Debug(message string) {
	l.log(DebugLevel, "DEBUG", message)
}

// Info 输出普通信息
func (l *ConsoleLogger) Info(message string) {
	l.log(InfoLevel, "INFO", message)
}

// Warn 输出警告信息
func (l *ConsoleLogger) Warn(message string) {
	l.log(WarnLevel, "WARN", message)
}

// Error 输出错误信息
func (l *ConsoleLogger) Error(message string) {
	l.log(ErrorLevel, "ERROR", message)
}
